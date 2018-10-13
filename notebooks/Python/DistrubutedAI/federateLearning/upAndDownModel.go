package main

import (
	"os"
	"io"
	"bytes"
	"mime/multipart"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"bufio"
	"flag"
)
// fileName2hash for save keep relationship between upload filename and hash
var fileName2hash map[string]string = make(map[string]string)

func HttpUpload(filename string, localfileName string) string{
	targetUrl := "http://36.26.80.184:8899/storage"
	//targetUrl := "http://localhost:8899/storage"

	//打开文件操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return ""
	}
	defer fh.Close()
	fmt.Println("open file success")
	//创建表单文件
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	// 处理路径名
	// fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", localfileName)
	// fmt.Println("fileWriter====>",fileWriter)
	if err != nil {
		fmt.Println("error writing to buffer")
		return ""
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("error copy file")
		return ""
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	//发起网络请求
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("发起网络请求", err)
		return ""
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp_body, err:",err)
		return ""
	}
	if resp.Status != "200 OK" {
		fmt.Println("resp.Status:",resp.Status)
		return ""
	}
	return string(resp_body)
}


func CheckFileExist(fileName string) bool {
    _, err := os.Stat(fileName)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func HttpDownload(hash string, downloadPath string)string{
	targetUrl := "http://36.26.80.184:8899/storage?hash=" + hash
	//targetUrl := "http://localhost:8899/storage?hash="+hash
	//发起网络请求
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()

	//读取文件
	resp_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		return ""
	}

	filename := string(response.Header.Get("Content-Disposition"))
	index := strings.Index(filename, "\"")
	filename = filename[index+1: len(filename)-1]
	fmt.Println(filename)

	// f, err := os.OpenFile("./downloadMnist/" + filename, os.O_WRONLY | os.O_CREATE, 0666)
	f, err := os.OpenFile(downloadPath + "/" + filename, os.O_WRONLY | os.O_CREATE, 0666)
	// ============
	// var filePathName = "./downloadMnist2/" + filename
	// var f *os.File
    // if CheckFileExist(filePathName) {  //文件存在
    //     f, err = os.OpenFile(filePathName, os.O_WRONLY | os.O_CREATE, 0666) //打开文件
    //     if err != nil{
    //         fmt.Println("file open fail", err)
    //         return ""
    //     }
    // }else {  //文件不存在
    //     f, err = os.Create("./downloadMnist2/") //创建文件
    //     if err != nil {
    //         fmt.Println("file create fail")
    //         return ""
    //     }
    // }
	// ========
	file := bufio.NewWriter(f)
	file.Write(resp_body)
	file.Flush()
	f.Close()
	return filename
}

// upload model
func uploadModel(Path string) {
	// folderPath := "./"+Path
	folderPath := Path
	filesName, _ := ioutil.ReadDir(folderPath)
	for _, localfile := range filesName {
		localfileName := folderPath + "/" + localfile.Name()
		fmt.Println("localfileName=====>:",localfileName)
		hash := HttpUpload(localfileName, localfile.Name())
		// save hash anf filename
		fileName2hash[localfileName] = hash
		// hash := HttpUpload("./test_images0")
		// fmt.Println("hash: " + hash)
		if hash != ""{
			fmt.Println("Upload successfully")
			fmt.Println("Hash = " + hash)
			// 下载模型层模型与数据
			HttpDownload(hash, "./model/downloadModel")
		}
	}
}

// download model
func downloadModel(hash string, path string) {
	if hash !="" {
		HttpDownload(hash, path)
	}
}

// upload data
func uploadData(folderPath string) {
	uploadModel(folderPath)
}
// download data
func downloadData(hash string) {
	if hash !="" {
		HttpDownload(hash, "downloadData")
	}
}

func main() {
	// 获取命令行命令判断
	isUp := flag.Bool("isUp", true, "isUploadModel");
	var path string
	flag.StringVar(&path, "path", "", "文件路径");
	var localHash string
	flag.StringVar(&localHash, "localHash", "", "本地哈希值");
	flag.Parse();

	fmt.Println("isUp: ", *isUp);
	fmt.Println("path: ", path);
	// 获取上传的模型
	if *isUp == true && path != "" && localHash ==""{
		fmt.Println("upload======>:");
		uploadModel(path);
	} else if *isUp == false && path != "" && localHash !=""{
		downloadModel(localHash, path)
	} else {
		fmt.Println("路径无效");
	}
	// 上传训练完成的模型
	// uploadModel("trainedModel")
	// 下载协议层下载数据
	// downloadModel("QmV1UsLkznQzzyMrFwHrTBkxos9WaDZKBX4mUBL4EZe2p8")
}
