from tensorflow.examples.tutorials.mnist import input_data
# mnist = input_data.read_data_sets("./MNIST_data",one_hot = True)

import tensorflow as tf
import os
import pickle
import numpy as np
import requests
import hashlib
import time
#Parameters
learning_rate = 0.1
training_epochs = 5

batch_size = 100
display_step = 1
#Network Parameters
n_input = 784
n_classes = 10

#tf Graph input
x = tf.placeholder("float",[None,n_input])
y = tf.placeholder("float",[None,n_classes])


#pre-define
#convolution layer
def conv2d(x,W):
    return tf.nn.conv2d(x,W,
                        strides=[1,1,1,1],
                        padding='SAME')
# pooling layer
def max_pool_2x2(x):
    return tf.nn.max_pool(x,ksize=[1,2,2,1],
                          strides=[1,2,2,1],
                          padding='SAME')
#Create model
def multilayer_preceptron(x,weights,biases):
    #now,we want to change this to a CNN network
    #first,reshape the data to 4_D ,
    x_image=tf.reshape(x,[-1,28,28,1])
    #then apply cnn layers ,cnn layer and activation function --relu
    h_conv1=tf.nn.relu(conv2d(x_image,weights['conv1'])+biases['conv_b1'])
    #first pool layer
    h_pool1=max_pool_2x2(h_conv1)
    #second cnn layer
    h_conv2=tf.nn.relu(conv2d(h_pool1,weights['conv2'])+biases['conv_b2'])
    #second pool layer
    h_pool2=max_pool_2x2(h_conv2)

    h_pool2_flat=tf.reshape(h_pool2,[-1,7*7*64])
    h_fc1=tf.nn.relu(tf.matmul(h_pool2_flat,weights['fc1'])+biases['fc1_b'])
    out_layer=tf.matmul(h_fc1,weights['out'])+biases['out_b']
    return out_layer

weights={
    'conv1':tf.Variable(tf.random_normal([5,5,1,32]), name = 'conv1'),
    'conv2':tf.Variable(tf.random_normal([5,5,32,64]), name = 'conv2'),
    'fc1':tf.Variable(tf.random_normal([7*7*64,256]), name = 'fc1'),
    'out':tf.Variable(tf.random_normal([256,n_classes]), name = 'out')
}
biases={
    'conv_b1':tf.Variable(tf.random_normal([32]), name = 'conv_b1'),
    'conv_b2':tf.Variable(tf.random_normal([64]), name = 'conv_b2'),
    'fc1_b':tf.Variable(tf.random_normal([256]), name = 'fc1_b'),
    'out_b':tf.Variable(tf.random_normal([n_classes]), name = 'out_b')
}
# 1. Construct model
pred = multilayer_preceptron(x,weights,biases)

#Define loss and optimizer
cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred,labels=y))
optimizer=tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)
#Initializing the variables
init = tf.global_variables_initializer()
# 2. Saver model
model_saver = tf.train.Saver()

def readAllEpoch():
    with open("allEpoch.pickle", 'rb') as f:
        data = pickle.load(f)
    print("---------->len:",data["conv1"])

#Launch the gtrph
with tf.Session() as sess:
    sess.run(init)
    model_dir = "uploadModel"
    model_name = "cpk"
    if not os.path.exists(model_dir):
        os.makedirs(model_dir)
    model_saver.save(sess,os.path.join(model_dir,model_name))
    print("model saved sucessfully!! the path is in flie uploadModel")
    # readAllEpoch()
# 3. upload model and data 多轮上传模型，数据？ round1, 2, 3?
# url = "http://text"
# clientId = "1111"
# clientKey = "2222"
# timestap = (str)(int(round(time.time()*1000)))
# clientSecret = hashlib.sha256(clientId.encode("utf-8") + clientKey.encode("utf-8") + timestap.encode("utf-8"))
#
# header = {'clientId': clientId, 'timestap': timestap, 'clientSecret': clientSecret}
# files = {'zip':open('~/Workspaces/DistributedAI/notebooks/Python/DistrubutedAI/computing/mnist.zip','rb')}
# data = {'enctype': 'multipart/form-data', 'name': 'liu'}
# reponse = requests.post(url, data=data, header=header, files=files)
# text = reponse.text
# print(text)
#
# 模型方：获取10个经过训练的模型参数，进行加权平均
# 加权平均
# read trained CNN model
# def readAllEpoch():
#     with open("allEpoch.pickle", 'rb') as f:
#         data = pickle.load(f)
#     print("---------->len:",len(data["conv1"]))

def averageWeight(nums, weights, biases):
    average_weights = sum(weights) / len(weights)
    average_biases = sum(weights) / len(biases)
    return average_weights, average_biases

def federatingLearning(nums, weights, biases):
    # average
    average_weights, average_biases = averageWeight(nums, weights, biases)
    return average_weights, average_biases
    pass

# 读取本地所有模型
def readModel(path):
    init_op = tf.global_variables_initializer()

    model_saver = tf.train.Saver()
    # Launch the gtrph
    with tf.Session() as sess:
        # create dir for model saver
        init = sess.run(init_op)
        model_dir = "./"+path
        model_name = "cpk"
        model_path = os.path.join(model_dir, model_name)
        model_saver.restore(sess, model_path)
        conv1 = sess.run(tf.get_default_graph().get_tensor_by_name("conv1:0"))
        conv2 = sess.run(tf.get_default_graph().get_tensor_by_name("conv2:0"))
        fc1 = sess.run(tf.get_default_graph().get_tensor_by_name("fc1:0"))
        out = sess.run(tf.get_default_graph().get_tensor_by_name("out:0"))
        conv_b1 = sess.run(tf.get_default_graph().get_tensor_by_name("conv_b1:0"))
        conv_b2 = sess.run(tf.get_default_graph().get_tensor_by_name("conv_b2:0"))
        fc1_b = sess.run(tf.get_default_graph().get_tensor_by_name("fc1_b:0"))
        out_b = sess.run(tf.get_default_graph().get_tensor_by_name("out_b:0"))
        weightsAndBiases = {
            "conv1": conv1,
            "conv2": conv2,
            "fc1": fc1,
            "out": out,
            "conv_b1": conv_b1,
            "conv_b2": conv_b2,
            "fc1_b": fc1_b,
            "out_b": out_b,
        }
        # print(weightsAndBiases)
        return weightsAndBiases
if __name__ == "__main__":
    modelList = np.array(["trainedModel", "uploadModel", "downloadModel"])
    for index in range(len(modelList)):
        print("============================>")
        weightsAndBiases = readModel(modelList[index])
        print(weightsAndBiases)
    pass