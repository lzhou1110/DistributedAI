# download InitModel
# 1. 得到model的hash值 2. 从ipfs上下载数据

go run upAndDownModel.go -isUp=false -path="./model/downloadModel" -localHash="QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs"
go run upAndDownModel.go -isUp=false -path="./model/downloadModel" -localHash="QmUSy3Mzsb2LD9jAQ24xWDmpMpYqXKdWi48vuyEWd1jzVF
go run upAndDownModel.go -isUp=false -path="./model/downloadModel" -localHash="QmYt7NHqo7Z2zyXMJPFJD8nALRYETM5LQDGn8Dj8dE3RR5"
go run upAndDownModel.go -isUp=false -path="./model/downloadModel" -localHash="QmVSyrkRMdp7FGJLDxhDWtPREFiBbkgcNaQomrh8ouJ5Qj"


# 3. 上传区块链
go run trans.go uploadDataComputing.go "QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs" "checkpoint" "123" "" "" ""
go run trans.go uploadDataComputing.go "QmUSy3Mzsb2LD9jAQ24xWDmpMpYqXKdWi48vuyEWd1jzVF" "cpk.data-00000-of-00001" "123" "" "" ""
go run trans.go uploadDataComputing.go "QmYt7NHqo7Z2zyXMJPFJD8nALRYETM5LQDGn8Dj8dE3RR5" "cpk.index" "123" "" "" ""
go run trans.go uploadDataComputing.go "QmVSyrkRMdp7FGJLDxhDWtPREFiBbkgcNaQomrh8ouJ5Qj" "cpk.meta" "123" "" "" ""