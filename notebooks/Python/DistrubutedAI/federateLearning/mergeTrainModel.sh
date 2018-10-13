# 1. 从协议层获取hash，模型方下载模型  2. 合并模型 3. 还原数据
# go run trans.go downloadModel.go "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"

# go run upAndDownModel.go -isUp=false -path="/mergeModel/trainModel" -localHash="QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs"
# go run upAndDownModel.go -isUp=false -path="/mergeModel/trainModelB" -localHash="QmUSy3Mzsb2LD9jAQ24xWDmpMpYqXKdWi48vuyEWd1jzVF

echo "合并数据"
python modelProvider.py "merge"