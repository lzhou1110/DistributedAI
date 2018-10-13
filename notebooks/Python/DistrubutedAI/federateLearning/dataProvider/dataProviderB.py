import json

# 设置价格

DataSchema = "json"
MetaSchema = "mnist data set test_image"
Payment = 77.0

bidData = {
    "MetaSchema" : MetaSchema,
    "DataSchema": DataSchema,
    "Payment": Payment,
}

bidDataJson = json.dumps(bidData)

fileObject = open('./dataProviderRep/dataProviderB.json', 'w')
fileObject.write(bidDataJson)
fileObject.close()