# upload data description
# isUp = true为上传文件， -path 为本地文件路径
# 流程: 1. 数据方本地生成json文件 2. 上传至ipfs层得到hash值 3. 将hash值通过页面UI传到协议层，返回hash值
echo "数据方本地生成数据描述文件"
python dataProvider/dataProviderB.py


# 2. 上传至ipfs层得到hash值 3. 将hash值通过页面UI传到协议层，返回hash值
echo "上传至ipfs层得到hash值"
go run upAndDownModel.go -isUp=true -path="./dataProviderRep"