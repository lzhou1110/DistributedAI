import rsa
import os
import math
import ipfsapi
import rsa.randnum
from Crypto.Cipher import AES


class PrpCrypt(object):

    def __init__(self, key):
        self.key = key
        self.mode = AES.MODE_CBC

    def encrypt(self, text):
        cryptor = AES.new(self.key, self.mode, b'0000000000000000')
        ciphertext = cryptor.encrypt(text)

        return ciphertext

def UpLoad(filename):

    api = ipfsapi.connect('127.0.0.1', 5001)
    f = open('HashRecord.txt', 'a')
    res = api.add(filename)
    f.write(res['Hash'] + '\n')
    os.remove(filename)
    f.close()


def splitfile(filename):
    #calculate and add the bytes to file
    totalSize = os.path.getsize(filename)
    add = 16 - totalSize % 16

    # initial the split
    totalSize = totalSize + add * len(b'\0')
    chunksize = int(math.pow(16, int(math.log(totalSize, 16))))
    if chunksize > 1024 * 1024 * 32:
        chunksize = 1024 * 1024 * 32
    ShardNum = int(totalSize / chunksize)

    # initial encrpyted
    key = rsa.newkeys(256)
    publicKey = key[0]
    privateKey = key[1]

    # random key for aes
    aes_key = rsa.randnum.read_random_bits(128)
    pc = PrpCrypt(aes_key)

    # write hashRecord
    with open('HashRecord.txt', 'w') as f:
        f.write(filename + '\n')
        f.write(str(ShardNum) + '\n')
        f.write(str(add) + '\n')
        f.write(str(privateKey.e) + '\n')
        f.write(str(privateKey.n) + '\n')
        f.write(str(privateKey.d) + '\n')
        f.write(str(privateKey.p) + '\n')
        f.write(str(privateKey.q) + '\n')
        f.close()

    # encrypt the aes_key and save upload it
    crypted_aes_key = rsa.encrypt(aes_key, publicKey)
    with open('Crypted', 'wb') as f:
        f.write(crypted_aes_key)
        f.close()
        UpLoad('Crypted')

    # open the fromfile
    inputfile = open(filename, 'rb')
    for partNum in range(0, ShardNum - 1):
        chunk = inputfile.read(chunksize)
        if not chunk:  # check chunk is empty
            print("error in spliting to parts")
            break
        filename = 'part%d' % partNum
        chunk = pc.encrypt(chunk)
        fileobj = open(filename, 'wb')  # make partfile
        fileobj.write(chunk)  # writedata into partfile
        fileobj.close()
        UpLoad(filename)
    if partNum + 1 == ShardNum - 1:
        chunk = inputfile.read(totalSize - (ShardNum - 1) * chunksize - add * len(b'\0'))
        filename = 'part%d' % (partNum + 1)
        chunk = pc.encrypt(chunk+add * b'\0')
        fileobj = open(filename, 'wb')  # make partfile
        fileobj.write(chunk)  # writedata into partfile
        fileobj.close()
        UpLoad(filename)
    else:
        print("can not split up to ShardNum parts")
    inputfile.close()