from tensorflow.examples.tutorials.mnist import input_data
import tensorflow as tf
import os
import pickle
import numpy as np
import sys

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
    #first,reshape the dataProvider to 4_D ,
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

# 建立模型，训练， 保存
def constructAndTrain():
    weights = {
        'conv1': tf.Variable(tf.random_normal([5, 5, 1, 32]), name='conv1'),
        'conv2': tf.Variable(tf.random_normal([5, 5, 32, 64]), name='conv2'),
        'fc1': tf.Variable(tf.random_normal([7 * 7 * 64, 256]), name='fc1'),
        'out': tf.Variable(tf.random_normal([256, n_classes]), name='out')
    }
    biases = {
        'conv_b1': tf.Variable(tf.random_normal([32]), name='conv_b1'),
        'conv_b2': tf.Variable(tf.random_normal([64]), name='conv_b2'),
        'fc1_b': tf.Variable(tf.random_normal([256]), name='fc1_b'),
        'out_b': tf.Variable(tf.random_normal([n_classes]), name='out_b')
    }
    # 1. Construct model
    pred = multilayer_preceptron(x, weights, biases)

    # Define loss and optimizer
    cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred, labels=y))
    optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)
    # Initializing the variables
    init = tf.global_variables_initializer()
    # 2. Saver model
    model_saver = tf.train.Saver()

    # Launch the gtrph
    with tf.Session() as sess:
        sess.run(init)
        model_dir = "model/uploadModel"
        model_name = "cpk"
        if not os.path.exists(model_dir):
            os.makedirs(model_dir)
        model_saver.save(sess, os.path.join(model_dir, model_name))
        # print("model saved sucessfully!! the path is in flie trainedModel")
        print("model merge sucessfully!! the path is in flie trainedModel")

# 将合并后的参数重新保存
def weightsRestore(allWeightsAndBiases):
    weights = {
        'conv1': allWeightsAndBiases["conv1"],
        'conv2': allWeightsAndBiases["conv2"],
        'fc1': allWeightsAndBiases["fc1"],
        'out': allWeightsAndBiases["out"]
    }
    biases = {
        'conv_b1': allWeightsAndBiases["conv_b1"],
        'conv_b2': allWeightsAndBiases["conv_b2"],
        'fc1_b': allWeightsAndBiases["fc1_b"],
        'out_b': allWeightsAndBiases["out_b"]
    }
    # 1. Construct model
    pred = multilayer_preceptron(x, weights, biases)

    # Define loss and optimizer
    cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred, labels=y))
    optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)
    # Initializing the variables
    init = tf.global_variables_initializer()
    # 2. Saver model
    model_saver = tf.train.Saver()

    # Launch the gtrph
    with tf.Session() as sess:
        sess.run(init)
        model_dir = "model/mergeModel"
        model_name = "cpk"
        if not os.path.exists(model_dir):
            os.makedirs(model_dir)
        model_saver.save(sess, os.path.join(model_dir, model_name))
        print("merge model saved sucessfully!! the path is in flie mergeModel")

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

def mergeModelWeights():
    modelList = np.array(["./model/mergeModel/trainedModel", "./model/mergeModel/trainedModelB"])
    allWeightsAndBiases = 0
    for index in range(len(modelList)):
        print("============================>")
        weightsAndBiases = readModel(modelList[index])
        allWeightsAndBiases += weightsAndBiases
        print(weightsAndBiases)
    allWeightsAndBiases /= len(modelList)
    return allWeightsAndBiases

if __name__ == "__main__":
    if len(sys.argv) > 1 and sys.argv[0] == "merge":
        allWeightsAndBiases = mergeModelWeights()
        weightsRestore(allWeightsAndBiases)
    else:
        constructAndTrain()
