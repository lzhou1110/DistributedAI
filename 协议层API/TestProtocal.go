package main

import (
	//"os"
	//"io"
	//"bytes"
	//"mime/multipart"
	"fmt"
	"net/http"
	//"io/ioutil"
	"strings"
	//"bufio"
)

/*func HttpUpload(filename string) string{
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
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
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
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if resp.Status != "200 OK" {
		fmt.Println(resp.Status)
		return ""
	}
	return string(resp_body)
}*/

func putParams(fileName string,ipfsHash string,description string)string{
	//targetUrl := "http://36.26.80.184:8899/storage?filename=" + fileName+"&ipfshash="+ipfsHash+"&description="+description
	targetUrl := "http://localhost:8899/invoke?filename=" + fileName+"&ipfshash="+ipfsHash+"&description="+description
	//发起网络请求
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()

	//读取文件
	/*resp_body, err := ioutil.ReadAll(response.Body)
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

	f, err := os.OpenFile("./" + filename, os.O_WRONLY | os.O_CREATE, 0666)
	file := bufio.NewWriter(f)
	file.Write(resp_body)
	file.Flush()
	f.Close()*/
  ipfsHash_des:=string(response.Header.Get("Content-Disposition"))
  index := strings.Index(ipfsHash_des, "\"")
	ipfsHash_des = ipfsHash_des[index+1: len(ipfsHash_des)-1]
	return ipfsHash_des
}


func main(){
	/*hash := HttpUpload("KeyFile")
	fmt.Println("hash: " + hash)
	if hash != ""{
		fmt.Println("Upload successfully")
		fmt.Println("Hash = " + hash)

	}*/
  ipfsHash_des:=putParams("KeyFile","QmPkQySUCQuwFvKjQ128prBd5k5DmEUa1spSH3LNQ9yrMF","abc")
  if ipfsHash_des!=""{
    fmt.Println("Test successfully!! the ipfsHash_des is: ",ipfsHash_des)
  }else{
    fmt.Println("Test fail!!")
  }

	//HttpDownload("QmPkQySUCQuwFvKjQ128prBd5k5DmEUa1spSH3LNQ9yrMF")
}
