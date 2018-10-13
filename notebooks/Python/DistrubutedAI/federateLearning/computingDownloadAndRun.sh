# 训练方 下载模型
go run upAndDownModel.go -isUp=false -path="./model/downloadModel" -localHash="jasjdsisnflafaf"
# 训练方 下载数据
go run upAndDownModel.go -isUp=false -path="./model/downloadData" -localHash="jasjdsisnflafaf"
# run 训练方训练数据
python computing.py