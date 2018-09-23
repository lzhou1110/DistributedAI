from tensorflow.examples.tutorials.mnist import input_data
# mnist = input_data.read_data_sets("./MNIST_data",one_hot = True)

import tensorflow as tf
import os
import numpy as np
# 获取数据， 模型
index = 2
trainPath = '/Users/liulifeng/Desktop/Work/mnist_data/train_images/images'+str(index)
trainData = np.loadtxt(trainPath)
testPath = '/Users/liulifeng/Desktop/Work/mnist_data/test_images/test_images'+str(index)
testData = np.loadtxt(trainPath)
# print("====>testData",len(testData))
batch_size = 25
display_step = 1
#Network Parameters
n_input = 784
n_classes = 10

#Parameters
learning_rate = 0.0001
training_epochs = 2

#tf Graph input
x = tf.placeholder("float",[None,n_input])
y = tf.placeholder("float",[None,n_classes])

#pre-define
def conv2d(x,W):
    return tf.nn.conv2d(x,W,
                        strides=[1,1,1,1],
                        padding='SAME')
def max_pool_2x2(x):
    return tf.nn.max_pool(x,ksize=[1,2,2,1],
                          strides=[1,2,2,1],
                          padding='SAME')
#Create model
def multilayer_preceptron(x,weights,biases):
    #now,we want to change this to a CNN network
    #first,reshape the data to 4_D
    x_image=tf.reshape(x,[-1,28,28,1])
    #then apply cnn layers
    h_conv1=tf.nn.relu(conv2d(x_image,weights['conv1'])+biases['conv_b1'])
    h_pool1=max_pool_2x2(h_conv1)

    h_conv2=tf.nn.relu(conv2d(h_pool1,weights['conv2'])+biases['conv_b2'])
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
#Construct model
pred = multilayer_preceptron(x,weights,biases)

#Define loss and optimizer
cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred,labels=y))
optimizer=tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)

# Calcuate accuracy
correct_prediction = tf.equal(tf.argmax(pred, 1), tf.argmax(y, 1))
accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))

#create class Saver
saver = tf.train.Saver()

# 在下面的代码中，默认加载了TensorFlow计算图上定义的全部变量
# 直接加载持久化的图
# saver = tf.train.import_meta_graph("mnist/cpk.meta")
# download mdoel data path
# downModelLoadPath = "/Users/liulifeng/Desktop/Work/uploadIpfs/model/"
# saver = tf.train.import_meta_graph(downModelLoadPath+"model_save.meta")

#Launch the gtrph
with tf.Session() as sess:
    # create dir for model saver
    # model_dir = "mnist"
    # model_name = "cpk"
    # model_path=os.path.join(model_dir,model_name)
    # saver.restore(sess,model_path)

    model_dir = "/Users/liulifeng/Desktop/Work/uploadIpfs/downloadMnist/"
    model_name = "cpk"
    model_path = os.path.join(model_dir, model_name)
    saver.restore(sess, model_path)


    # saver.restore(sess, "mnist/cpk.index")
    # saver.restore(sess, downModelLoadPath+"model_save.index")

    # ----------------------------------------------------------
    # Training cycle
    for epoch in range(training_epochs):
        avg_cost = 0.
        # total_batch = int(mnist.train.num_examples / batch_size)
        total_batch = int(len(trainData) / batch_size)
        # Loop over all batches
        for i in range(total_batch):
            # batch_x, batch_y = mnist.train.next_batch(batch_size)
            start = i*batch_size
            end = (i+1)*batch_size
            if (i + 1) * batch_size > len(trainData):
                end = len(trainData)
            batchX = trainData[start: end]
            labelArray = np.array([0, 0, 0., 0, 0, 0, 0, 0, 0, 0])
            labelData = np.array([0, 0, 0., 0, 0, 0, 0, 0, 0, 0])
            labelData[index] = labelArray[index] = 1.0
            for _ in range(len(trainData) - 1):
                labelData = np.vstack((labelData, labelArray))
            batchY = labelData[start: end]
            # run optimization op (backprop)and cost op (to get loss value)
            # _, c = sess.run([optimizer, cost], feed_dict={x: batch_x, y: batch_y})
            _, c, correct = sess.run([optimizer, cost, accuracy], feed_dict={x: batchX, y: batchY})
            # Compute average loss
            avg_cost += c / total_batch
            if i % 50 == 0:
                print("=====>test:", '%04d' % i, "cost=", "{:.9f}".format(c))
                print("<===correct:", correct)
            # Display logs per epoch step
        if epoch % display_step == 0:
            print("Epoch:", '%04d' % (epoch + 1), "cost=", "{:.9f}".format(avg_cost))
    print("Optimization Finished!")

    labelArray = np.array([0, 0, 0., 0, 0, 0, 0, 0, 0, 0])
    labelData = np.array([0, 0, 0., 0, 0, 0, 0, 0, 0, 0])
    labelData[index] = labelArray[index] = 1.0
    for _ in range(len(testData) - 1):
        labelData = np.vstack((labelData, labelArray))
    # print("Accuracy:", accuracy.eval({x: mnist.test.images, y: mnist.test.labels}))
    print("labelData len:",len(labelData))
    # print("====>testData",testData,"===>labelData",labelData)
    print("Test Accuracy:", accuracy.eval({x: testData, y: labelData}))
    # ----------------------------------------------------------

    # upload model and data

    # predict model
    # img=mnist.test.images[100].reshape(-1,784)
    # img_label=sess.run(tf.argmax(mnist.test.labels[100]))
    #
    # ret=sess.run(pred,feed_dict={x:img})
    # num_pred=sess.run(tf.argmax(ret,1))
    #
    # print("预测值:%d\n" % num_pred)
    # print("真实值:",img_label)
    # print("模型恢复成功")
