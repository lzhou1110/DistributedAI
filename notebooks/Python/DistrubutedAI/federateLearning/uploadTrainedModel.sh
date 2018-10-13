# 1. 训练方A, B分别训练数据C，D

python computing.py
python computingB.py

# 2. 分别保存模型上传ipfs层
go run upAndDownModel.go -isUp=true -path="./model/trainedModel"
go run upAndDownModel.go -isUp=true -path="./model/trainedModelB"

# 3. 上传协议层
# go run trans.go uploadModel.go "./model/uploadModel/checkpoint" "QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.data-00000-of-00001" "QmPzisU28SKL7ng2VdtZ8fGukEme4iJeMhv4sd9bSgJkUL"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.index" "QmNb2aVTAhhkV2ZBTpLPo3hdsRS36TMfeEQCvdMNo2cj2e"
# go run trans.go uploadModel.go "./model/uploadModel/cpk.meta" "QmNhqtmYyb8d5sMyoQLvaU59JkkEwAsHSchXAAPGWXE946"


