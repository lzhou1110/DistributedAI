******  
_Name_: File Manage System  
_Version_: 1.0  
_Date_: 2018-5-2  
_Introduction_: This project is to upload and download files from **IPFS** with **AES(AES128-CBC)**, **RSA** encryption and **Sharding**. Now we have successfully tested the project for file format **txt, zip, bmp, MP4, nc, mkv** and the maxsize of a file is about **2GB**. We believe other file formats with larger size are feasible.
******  

## Example ##
a video: 2.mkv (2GB)     
add time : 554s  
get time : 329s  
  
![avatar](http://chuantu.biz/t6/301/1525258288x-1404758293.png)
![avatar](http://chuantu.biz/t6/301/1525258345x-1404758293.png)
![avatar](http://chuantu.biz/t6/301/1525258357x-1404758293.png)
  
## Before Install ##
you need to prepare:
### ① ipfs-go ###
download ipfs-go [here](https://dist.ipfs.io/#go-ipfs) and then install ipfs in command with：  
  
`cd/d D:\go-ipfs`    _-----the path you install go-ipfs_

`ipfs  init`  

`ipfs  daemon`  
### ② python ###
install following packages:  

`ras`

`pycryptodome`
  
`ipfsapi`

    
## Usage and Step ##

### ① start ipfs ###
Before you execute any instruction, please start with the following instructions in command: 
  
`cd/d D:\go-ipfs` _--the path you install go-ipfs_

`ipfs  daemon`  _--start the ipfs_

### ② Add files to ipfs ###
Don't shut down the command window for ipfs and restart a command window to excute:
#### Usage: ####
  
`cd/d  D:/pyproject/fms`  _--the path of the fms project_
   
`python fms.py add <filename>`      	_--upload the file to ipfs_
#### E.g.: ####
`python fms.py add dataset.txt`
#### Step: ####
1. Sharding the original file into several shards.  
2. Generate a random AES key.   
3. Encrypt every shard separately with the AES key.  
4. Encrypt the AES key and generate a pair of public&private key for it using RSA. 
5. Upload all encrypted shards and encrypted AES key to IPFS.  
6. Save hashes of shards and the encrypted AES key, the private key(for decrypting the AES key), shard number, and other neccesary information into _HashRecord.txt_.   
![avatar](http://chuantu.biz/t6/301/1525258093x-1404781216.png)
***************
                
### ③ Download files from ipfs: ###
    
#### Usage: ####
`python fms.py get <HashFileName>`	               _--download the file with record information to current directory_  
`python fms.py get <HashFileName> <TargetDir>`	   _--download the file with record information to target directory_
#### E.g.: ####
`python fms.py get HashRecord.txt`   

`python fms.py get HashRecord.txt D:/Recover`
   
#### Step: ####
1. Download the shards from IPFS.  
2. Decrypt the AES key with the private key.  
3. Decrypt every shard with the AES key.   
4. Merge into a complete file.   
5. Save the file to target directory.    
***************
   
### ④ Help: ###
`python fms.py help `	           _--check all commands_ 
**************
    
### ⑤ The format of HashRecord.txt: ###
Line 1: **FileName**  
Line 2: **ShardNumber**  
Line 3: **The bytes add to file's tail _(AES encryption requires the whole file to be divisible by 16)_**  
Line 4-8: **private key**  
Line 9: **the Hash of the encrypted AES**    
Line 10: **the Hash of the First SubFile**  
Line 11: **the Hash of the Second SubFile**  
....  
Line ShaedNumber+9: **the Hash of the Last SubFile**  
 **************
    
### ⑥ Frequently Problem ###
`ConnectionRefusedError: [WinError 10061] 由于目标计算机积极拒绝，无法连接。`	

Solve it with:
`ipfs daemon`
