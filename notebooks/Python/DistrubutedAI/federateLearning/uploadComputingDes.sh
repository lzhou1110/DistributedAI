# upload data price
# isUp = true为上传文件， -path 为本地文件路径
# 流程：1.本地生成json文件  2. 上传ipfs层返回得到hash值  3. 将hash上传至协议层
echo "训练方本地生成运算资源文件"
python computing/computingProviderB.py


echo "训练方上传ipfs层返回得到hash值"
go run upAndDownModel.go -isUp=true -path="./computingRep"