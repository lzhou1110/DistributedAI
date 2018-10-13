import json

# 设置价格
DataSchema = "json"
MetaSchema = "mnist data set a description"
Payment = 112.0

bidData = {
    "MetaSchema" : MetaSchema,
    "DataSchema": DataSchema,
    "Payment": Payment,
}

bidDataJson = json.dumps(bidData)

fileObject = open('./dataProviderRep/dataProviderA.json', 'w')
fileObject.write(bidDataJson)
fileObject.close()