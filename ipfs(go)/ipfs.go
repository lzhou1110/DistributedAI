package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"encoding/csv"
	"os"
	"strings"
	"crypto/aes"
	"crypto/cipher"

	"math/rand"
	"time"
	"io/ioutil"
	"io"
	"log"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	//origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}


func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


//random16 bytes AES key
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func AesAdd(originalFilename string) {
	// AES-256。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	rand.Seed(time.Now().Unix())
	key := []byte(RandStringBytes(32))
	//fmt.Println("key:",key)
	file, err := os.Open("origData/"+originalFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	record, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Println("original data readed:",record)
	result, err := AesEncrypt([]byte(record), key)
	if err != nil {
		panic(err)
	}
	//fmt.Println("encrpted data:",base64.StdEncoding.EncodeToString(result))
	file.Close()

	ioutil.WriteFile("encrypted_"+originalFilename, result, os.ModeAppend)
	hash := Upload_file("encrypted_" + originalFilename) //upload to ipfs , hash is string

	//delete file
	del := os.Remove("encrypted_" + originalFilename)
	if del != nil {
		fmt.Println(del)
	}
	// remove the expanded-name
	fileNames := strings.Split(originalFilename, ".")

	hashfileName := "Hash_"+fileNames[0]+".txt"
	hashfile, err := os.OpenFile("hash/"+hashfileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//write date and hash
	hashfile.Write([]byte(originalFilename))
	hashfile.Write([]byte("\n"))
	hashfile.Write([]byte(hash))
	hashfile.Close()

	keyfileName :="Key_"+fileNames[0]+".txt"
	keyfile, err := os.OpenFile("key/"+keyfileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	keyfile.Write([]byte(originalFilename))
	keyfile.Write([]byte("\n"))
	keyfile.Write(key)
	keyfile.Close()
}

func AesGet(hashfileName string, hashstr string , keyfileName string, keystr string) {
	hashfile, err := os.OpenFile(hashstr+hashfileName,os.O_RDONLY,0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	hashcontent := make([]byte, 1024)
	hashfile.Read(hashcontent)
	hash_out_str := strings.Split(string(hashcontent), "\n")
	hash_originalFilename := hash_out_str[0]
	hash := hash_out_str[1]

	keyfile, err := os.OpenFile(keystr+keyfileName,os.O_RDONLY,0666)
	keycontent := make([]byte, 1024)
	keyfile.Read(keycontent)
	key_out_str := strings.Split(string(keycontent), "\n")
	key_originalFilename := key_out_str[0]
	key_out_str = strings.Split(key_out_str[1], "\000")
	key:= []byte(key_out_str[0])
	if hash_originalFilename ==key_originalFilename{
		originalFilename := key_originalFilename
		encrytedfileName :="get_encrypted_" + originalFilename
		Download_file(hash,encrytedfileName)
		encryptedfile, err := os.Open(encrytedfileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		encryptedrecord ,err := ioutil.ReadAll(encryptedfile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		origData, err := AesDecrypt([]byte(encryptedrecord), key)
		if err != nil {
			panic(err)
		}
		// 如果文件不存在，则以 perm 权限创建该文件
		// 如果文件存在，则先清空文件，然后再写入
		ioutil.WriteFile("getData/decrypted_"+originalFilename, origData, os.ModeAppend)
		encryptedfile.Close()
		//delete file
		del := os.Remove("get_encrypted_" + originalFilename)
		if del != nil {
			fmt.Println(del)
		}
	}

}


func Upload_file(filename string) string {
	cmd := exec.Command("ipfs", "add", "-r", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
	out_str := strings.Split(out.String(), " ")
	hash := out_str[1]
	return hash
}



func Download_file(hash string, filename string) {
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command("ipfs", "get", finalhash, "-o="+filename)
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
}
func ShardCsvByDay(fileName string){

	file, _ := os.Open(fileName)
	reader := csv.NewReader(file)
	title, _ := reader.Read()

	lastDate := ""
	newFileName := ""
	data, err := reader.Read()
	for err != io.EOF{
		site := data[0]
		dateTemp := data[2]
		date1 := strings.Split(dateTemp," ")
		date := date1[0]
		times := strings.Split(date, "/")
		day := times[0]
		month := times[1]
		year := times[2]

		if date != lastDate{
			if newFileName!=""{
				AesAdd(newFileName)
				tempFileNames:=strings.Split(newFileName, ".")
				AesGet("Hash_"+tempFileNames[0]+".txt","hash/","Key_"+tempFileNames[0]+".txt","key/")
			}

			lastDate = date
			newFileName = site+"_"+year+month+day+".csv"
			writeFile, _:= os.OpenFile("origData/"+newFileName, os.O_CREATE, os.ModePerm)
			writer := csv.NewWriter(writeFile)
			writer.Write(title)
			writer.Flush()
			writeFile.Close()
		}
		writeFile, _ := os.OpenFile("origData/"+newFileName, os.O_WRONLY|os.O_APPEND, 0666)
		writer := csv.NewWriter(writeFile)
		writer.Write(data)
		writer.Flush()
		writeFile.Close()

		data, err = reader.Read()
	}
	AesAdd(newFileName)
	tempFileNames:=strings.Split(newFileName, ".")
	AesGet("Hash_"+tempFileNames[0]+".txt","hash/","Key_"+tempFileNames[0]+".txt","key/")

	file.Close()
}

func CreatePath(path string)  {
	_, err := os.Stat(path)
	if err == nil {
		err2 := os.RemoveAll(path)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	err3 := os.Mkdir(path, os.ModePerm)
	if err3 != nil {
		log.Fatal(err3)
	}
}

func main() {
	originalFilename:= "PM10_KnC_CromwellRoad.csv"
	CreatePath("./hash")
	CreatePath("./key")
	CreatePath("./origData")
	CreatePath("./getData")
	ShardCsvByDay(originalFilename)

}
