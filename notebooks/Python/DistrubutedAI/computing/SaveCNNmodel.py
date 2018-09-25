from tensorflow.examples.tutorials.mnist import input_data
# mnist = input_data.read_data_sets("./MNIST_data",one_hot = True)

import tensorflow as tf
import os
import pickle
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
    'conv1':tf.Variable(tf.random_normal([5,5,1,32])),
    'conv2':tf.Variable(tf.random_normal([5,5,32,64])),
    'fc1':tf.Variable(tf.random_normal([7*7*64,256])),
    'out':tf.Variable(tf.random_normal([256,n_classes]))
}
biases={
    'conv_b1':tf.Variable(tf.random_normal([32])),
    'conv_b2':tf.Variable(tf.random_normal([64])),
    'fc1_b':tf.Variable(tf.random_normal([256])),
    'out_b':tf.Variable(tf.random_normal([n_classes]))
}
# 1. Construct model
pred = multilayer_preceptron(x,weights,biases)

#Define loss and optimizer
cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred,labels=y))
optimizer=tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)
#Initializing the variables
init = tf.global_variables_initializer()
# 2. Saver model
model_saver = tf.train.Saver(var_list=tf.global_variables())

#Launch the gtrph
with tf.Session() as sess:
    sess.run(init)
    #Training cycle
    # for epoch in range(training_epochs):
    #     avg_cost=0.
    #     total_batch=int(mnist.train.num_examples/batch_size)
    #     #Loop over all batches
    #     for i in range(total_batch):
    #         batch_x,batch_y=mnist.train.next_batch(batch_size)
    #         #run optimization op (backprop)and cost op (to get loss value)
    #         _,c=sess.run([optimizer,cost],feed_dict={x:batch_x,y:batch_y})
    #         #Compute average loss
    #         avg_cost+=c/total_batch
    #         #Display logs per epoch step
    #     if epoch % display_step==0:
    #         print("Epoch:",'%04d' % (epoch+1),"cost=","{:.9f}".format(avg_cost))
    # print("Optimization Finished!")
    # correct_prediction=tf.equal(tf.argmax(pred,1),tf.argmax(y,1))
    # #Calcuate accuracy
    # accuracy = tf.reduce_mean(tf.cast(correct_prediction,"float"))
    # print("Accuracy:",accuracy.`({x:mnist.test.images,y:mnist.test.labels}))

    #create dir for model saver
    model_dir = "uploadModel"
    model_name = "cpk"
    if not os.path.exists(model_dir):
        os.makedirs(model_dir)
    model_saver.save(sess,os.path.join(model_dir,model_name))
    print("model saved sucessfully!! the path is in flie uploadModel")

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
def readAllEpoch():
    with open("allEpoch.pickle", 'rb') as f:
        data = pickle.load(f)
    print("---------->len:",len(data["conv1"]))

def averageWeight(nums, weights, biases):
    average_weights = sum(weights) / len(weights)
    average_biases = sum(weights) / len(biases)
    return average_weights, average_biases

def federatingLearning(nums, weights, biases):
    # average
    average_weights, average_biases = averageWeight(nums, weights, biases)
    return average_weights, average_biases
    pass