from tensorflow.examples.tutorials.mnist import input_data
mnist = input_data.read_data_sets("./MNIST_data",one_hot = True)

import tensorflow as tf
import os

batch_size = 100
display_step = 1
#Network Parameters
n_input = 784
n_classes = 10

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
#create class Saver
model_saver = tf.train.Saver()

#Launch the gtrph
with tf.Session() as sess:
    #create dir for model saver
    model_dir = "mnist"
    model_name = "cpk"
    model_path=os.path.join(model_dir,model_name)
    model_saver.restore(sess,model_path)

    img=mnist.test.images[100].reshape(-1,784)
    img_label=sess.run(tf.argmax(mnist.test.labels[100]))

    ret=sess.run(pred,feed_dict={x:img})
    num_pred=sess.run(tf.argmax(ret,1))

    print("预测值:%d\n" % num_pred)
    print("真实值:",img_label)
    print("模型恢复成功")