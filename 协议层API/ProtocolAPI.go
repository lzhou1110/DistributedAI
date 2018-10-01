package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	//"strconv"
	"strings"
	//"time"
)

//var hash2Filename map[string]string = map[string]string{}

const contractAddress string = "0xf1b8919da687c85e5aeb7ff65e4e9f8f0fff6c8f"

func index(response http.ResponseWriter, request *http.Request) {
	html := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
    			<meta charset="UTF-8">
    			<title>Title</title>
			</head>
			<body>
    			<form id="formID" action="http://localhost:8899/invoke" method="get" enctype="text/plain">
        			<!--<input type="file" name="uploadfile" />-->
              <input type="text" name="ipfshash"/>
              <input type="text" name="filename"/>
        			<input type="hidden" name="token" value="{{.}}"/>
              <input type="text" name="description"/>
        			<input type="submit" value="upload" />
				</form>
			</body>
			</html>
		`
	io.WriteString(response, html)

}

func Invoke_interact(ipfsHash_des string,fileName string){
  myIpfsHash_des := strings.Split(ipfsHash_des, "\000")
	finalhash := myIpfsHash_des[0]
	cmd := exec.Command("go", "run", "ballot2.go","main2.go", fileName,finalhash, contractAddress)
  buf, err2 := cmd.Output()
  if err2!=nil{
    fmt.Printf("get the command line output fail: %v",err2)
  }

	err := cmd.Run()
	if err != nil {
		fmt.Print("run error: %v\n",err)
	}
	fmt.Println("interact with smart contract successfully!")
  fmt.Printf("output: %s\n" , buf)
}

func Download_file(hash string, filename string) {
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command("ipfs", "get", finalhash, "-o="+filename)
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("download from ipfs successfully!")
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
	//fmt.Println("upload to ipfs successfully!")
	//keep the relationship between IPFS hash and filename,needed to data persistence
	//hash2Filename[hash]=filename
	Download_file(hash, "downloadFile/"+filename)
	return hash
}



func storage(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	if request.Method == "POST" {
		fmt.Println("Get the post request")
		// save the file
		request.ParseMultipartForm(32 << 20)

		// get the handler of file
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println("form file err: ", err)
		}
		defer file.Close()

		//create(copy) the file upload
		fmt.Println("Get the file : " + handler.Filename)
		//f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open file err: ", err)
			return
		}
		defer f.Close()

		fmt.Println("Copy file ...")
		//copy file
		io.Copy(f, file)
		fmt.Println("File saved successfully")
		f.Close()

		fmt.Println("Upload file ...")
		//hash := Upload_file("./files/" + handler.Filename)
		hash := Upload_file(handler.Filename)
		fmt.Println("Upload file successfully! Hash :" + hash)

		//os.Remove("./files/" + handler.Filename)
		os.Remove(handler.Filename)
		fmt.Println("Delete temp file successfully!")

		fmt.Fprintf(response, "%v", hash)

	}
	if request.Method == "GET" {


	}

}


func invoke(response http.ResponseWriter, request *http.Request){
  request.ParseForm()
  if request.Method == "GET" {
    if len(request.Form["ipfshash"]) > 0 {
			ipfsHash := request.Form["ipfshash"][0]
      fileName:=request.Form["filename"][0]
      description:=request.Form["description"][0]

      ipfsHash_des:=ipfsHash+":"+description

			fmt.Println("Get the ipfshash : " + ipfsHash)
      fmt.Println("Get the filename: "+fileName)
      fmt.Println("Get the ipfsHash_des: "+ipfsHash_des)
			//paramFileName:=request.Form["fileName"][0]
			//fmt.Println("=======get fileName from cliet=====:"+paramFileName)
			//FName:=hash2Filename[hash]
			//FName := "f_" + strconv.FormatInt(time.Now().Unix(), 10)
			//fileName := "downloadFile/" + FName
			//fmt.Println("Get the file from ipfs, file will save to " + fileName)
			//Download_file(hash, "./"+fileName)
      Invoke_interact(ipfsHash_des,fileName)
			//fmt.Println("invoke successfully!!")

			// Turn to download page
			//fmt.Println("Begin downloading")
			//http.Redirect(response, request, "http://localhost:1010/files/" + FName, http.StatusFound)

			// Set Response Type
			response.Header().Set("Content-Type", "text/plain")
			response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; ipfsHash_des=\"%s\"", ipfsHash_des))
      fmt.Fprint(response, "succeed to invoke protocal API")
			//Open file
			/*f, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
			if err != nil {
				fmt.Println("open file err: ", err)
				return
			}
			defer f.Close()

			//File writer
			bufferWriter := bufio.NewWriter(response)
			bufferReader := bufio.NewReader(f)

			buffer := make([]byte, 4096) // 4KB

			//Begin Copy file
			for {
				n, err := bufferReader.Read(buffer)
				bufferWriter.Write(buffer[0:n])

				if err != nil {
					if err != io.EOF {
						fmt.Println("Read file error")
					}
					if err == io.EOF {
						break
					}
				}
			}
			bufferWriter.Flush()
			fmt.Println("Download successfully")
			f.Close()
			/*os.Remove("./" + fileName)
			fmt.Println("Remove file successfully")*/

		} else {
			fmt.Fprint(response, "Failed to invoke protocal API")
		}

  }
}



func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/invoke", invoke)
	//http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./downloadFile"))))
	http.ListenAndServe(":8899", nil)
}
