# upload model
# 1. 保存初始化模型
# 初始化model 数据
python modelProvider.py

# 2. 上传至ipfs
# isUp = true为上传文件， -path 为本地文件路径
go run upAndDownModel.go -isUp=true -path="./model/uploadModel"

# 3. 上传协议层
# go run trans.go uploadModel.go "./model/uploadModel/checkpoint" "QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.data-00000-of-00001" "QmPzisU28SKL7ng2VdtZ8fGukEme4iJeMhv4sd9bSgJkUL"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.index" "QmNb2aVTAhhkV2ZBTpLPo3hdsRS36TMfeEQCvdMNo2cj2e"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.meta" "QmNhqtmYyb8d5sMyoQLvaU59JkkEwAsHSchXAAPGWXE946"

# 4. 训练方下载模型，数据
go run trans.go downloadModel.go "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"
go run trans.go downloadModel.go "0x60ea115ff78700c6d698c703a1f5a4f1115fd1a0"