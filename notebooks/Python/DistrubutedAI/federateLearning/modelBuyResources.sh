#step 5 模型方购买数据方，运算方资源 hash
# 1. 模型方浏览相应UI页面

# 2. 选择相应数据，运算资源购买 得到hash值

# go run trans.go uploadDataComputing.go "QmVjebptMcJpuEppB3djYxEWQ92wQY3nYh3yKreQ6o4KSs" "checkpoint" "123" "" "" ""
# go run trans.go uploadDataComputing.go "QmUSy3Mzsb2LD9jAQ24xWDmpMpYqXKdWi48vuyEWd1jzVF" "cpk.data-00000-of-00001" "123" "" "" ""
# go run trans.go uploadDataComputing.go "QmYt7NHqo7Z2zyXMJPFJD8nALRYETM5LQDGn8Dj8dE3RR5" "cpk.index" "123" "" "" ""
# go run trans.go uploadDataComputing.go "QmVSyrkRMdp7FGJLDxhDWtPREFiBbkgcNaQomrh8ouJ5Qj" "cpk.meta" "123" "" "" ""

# 购买 数据 返回hash
go run trans.go buyDataComputing.go "checkpoint:123" "0x60ea115ff78700c6d698c703a1f5a4f1115fd1a0" "2 cpu and 8g :120" "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"
go run trans.go buyDataComputing.go "cpk.data-00000-of-00001:123" "0x60ea115ff78700c6d698c703a1f5a4f1115fd1a0" "2 cpu and 8g :120" "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"


# 购买 运算资源  返回hash
go run trans.go buyDataComputing.go "cpk.index:123" "0x60ea115ff78700c6d698c703a1f5a4f1115fd1a0" "2 cpu and 8g :120" "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"
go run trans.go buyDataComputing.go "checkpoint:123" "0x60ea115ff78700c6d698c703a1f5a4f1115fd1a0" "2 cpu and 8g :120" "0x9a5123ba78f6645c1dcbb05074d00dfcaf4b4373"