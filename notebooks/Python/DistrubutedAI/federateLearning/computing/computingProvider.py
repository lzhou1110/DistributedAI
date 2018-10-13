import json

# 设置价格
OperationSchemaAddress = "运算类型地址"
ComputingAddress = "运算资源地址"
PaymentAddress = "运算价格地址"
ComputerAttributesAddress = "运算资源信息"
price = 112.0

bidData = {
    "OperationSchemaAddress": OperationSchemaAddress,
    "ComputingAddress":ComputingAddress,
    "PaymentAddress": PaymentAddress,
    "ComputerAttributesAddress": ComputerAttributesAddress,
    "price": price
};

bidDataJson = json.dumps(bidData)

fileObject = open('./computingRep/computingA.json', 'w')
fileObject.write(bidDataJson)
fileObject.close()