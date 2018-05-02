import rsa
import sys,os
import ipfsapi
from Crypto.Cipher import AES
import rsa.randnum


class PrpCrypt(object):

    def __init__(self, key):
        self.key = key
        self.mode = AES.MODE_CBC

    def decrypt(self, text):
        cryptor = AES.new(self.key, self.mode,  b'0000000000000000')
        plain_text = cryptor.decrypt(text)
        return plain_text


def DownLoad(HashFileName):
    f = open(HashFileName, 'r')
    f.readline()
    ShardNum = f.readline()
    for i in range(0,7):
        f.readline()

    ShardNum = int(ShardNum.strip('\n'))

    try:
        api = ipfsapi.connect('127.0.0.1', 5001)
        for i in range(0, ShardNum):
            filename = "part" + str(i)
            Hash = f.readline()
            Hash = Hash.strip('\n')
            api.get(Hash)
            os.rename(Hash, filename)
        f.close()
    except:
        print(sys.exc_info())

def mergefile(filename,ShardNum, add):
    if os.path.exists(filename):
        os.remove(filename)
    outfile = open(filename,'wb')
    for partNum in range(0, ShardNum):
        partFile = 'part' + str(partNum)
        infile = open(partFile , 'rb')
        data = infile.read()
        if partNum == ShardNum - 1:
            data = data[0:-add * len(b'\0')]
        outfile.write(data)
        infile.close()
        os.remove(partFile)
    outfile.close()


def get(Hashfilename,TargetDir):

    #DownLoad all parts
    DownLoad(Hashfilename)

    #get the information
    f = open(Hashfilename, 'r')
    filename = f.readline().strip('\n')
    ShardNum = int(f.readline().strip('\n'))
    add = int(f.readline().strip('\n'))
    privateKey_e = int(f.readline().strip('\n'))
    privateKey_n = int(f.readline().strip('\n'))
    privateKey_d = int(f.readline().strip('\n'))
    privateKey_p = int(f.readline().strip('\n'))
    privateKey_q = int(f.readline().strip('\n'))
    privateKey = rsa.PrivateKey(privateKey_n, privateKey_e, privateKey_d, privateKey_p, privateKey_q)

    #get the encrapyed aes key
    Crypted_Hash = f.readline().strip('\n')
    api = ipfsapi.connect('127.0.0.1', 5001)
    api.get(Crypted_Hash)
    f.close()
    with open(Crypted_Hash, 'rb') as f:
        crypted_aes_key = f.read()
    f.close()
    os.remove(Crypted_Hash)

    # decrypt the aes key
    aes_key = rsa.decrypt(crypted_aes_key, privateKey)

    # decrypt all files
    pc = PrpCrypt(aes_key)
    for i in range(0, ShardNum):
        partFile = 'part' + str(i)
        f = open(partFile, 'rb')
        value = f.read()
        f.close()
        decry_value = pc.decrypt(value)
        f = open(partFile, 'wb')
        f.write(decry_value)
        f.close()

    # merge files
    filepath = os.path.join(TargetDir, filename)
    mergefile(filepath, ShardNum, add)
