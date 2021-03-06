import os
import tensorflow as tf
from tensorflow.examples.tutorials.mnist import input_data
from tensorflow.python.framework import graph_util

# 初始化CNN的权重
def weight_variable(shape):
    initial = tf.truncated_normal(shape=shape, stddev=0.1)
    return tf.Variable(initial)


# 初始化CNN的偏置
def bias_variable(shape):
    initial = tf.constant(0.1, shape=shape)
    return tf.Variable(initial)


# 卷积层定义
def conv2d(x, W):
    return tf.nn.conv2d(x, W, strides=[1, 1, 1, 1], padding='SAME')


# 池化层定义
def max_pool_2x2(x):
    return tf.nn.max_pool(x, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding='SAME')

# pb模型恢复
def load_model(path_to_model):
    if not os.path.exists(path_to_model):
        raise ValueError("'path_to_model.pb' is not exist.")
    model_graph = tf.Graph()
    with model_graph.as_default():
        od_graph_def = tf.GraphDef()
        with tf.gfile.GFile(path_to_model.pb, 'rb') as fid:
            serialized_graph = fid.read()
            od_graph_def.ParseFromString(serialized_graph)
            tf.import_graph_def(od_graph_def, name='')
    return model_graph

# 训练模型方法：模型种类，结构(list), 权重参数
def train(neuralNetworkKind, structure, weights):
    '''
    :param neuralNetworkKind: 模型种类 simpleFullConnect-0,CNN-1
    :param structure: 结构list数组，如[3，4，4]
    :param weights: 参数权重，list数组
    :return: null
    '''
    mnist = input_data.read_data_sets('./MNIST_data', one_hot=True)
    # 建立会话
    sess = tf.InteractiveSession()
    # 定义输入的placeholder
    # x_image 使用reshape将1*784的形式转换成原始形式28*28  -1代表样本数量不固定，1代表颜色通道数量
    x = tf.placeholder(tf.float32, [None, 784])
    y_ = tf.placeholder(tf.float32, [None, 10])
    x_image = tf.reshape(x, [-1, 28, 28, 1])

    # 定义第一层卷积层
    W_conv1 = weight_variable([5, 5, 1, 32])
    b_conv1 = bias_variable([32])
    h_conv1 = tf.nn.relu(conv2d(x_image, W_conv1) + b_conv1)
    h_pool1 = max_pool_2x2(h_conv1)

    # 第二层卷积层
    W_conv2 = weight_variable([5, 5, 32, 64])
    b_conv2 = bias_variable([64])
    h_conv2 = tf.nn.relu(conv2d(h_pool1, W_conv2) + b_conv2)
    h_pool2 = max_pool_2x2(h_conv2)

    # 定义第三层全连接层
    W_fc1 = weight_variable([7 * 7 * 64, 1024])
    b_fc1 = bias_variable([1024])
    h_pool2_flat = tf.reshape(h_pool2, [-1, 7 * 7 * 64])
    h_fc1 = tf.nn.relu(tf.matmul(h_pool2_flat, W_fc1) + b_fc1)

    # dropout层 用来减少过拟合
    keep_prob = tf.placeholder(tf.float32)
    h_fc1_drop = tf.nn.dropout(h_fc1, keep_prob)

    # 将dropout层输出连接一个softmax层 得到概率输出
    W_fc2 = weight_variable([1024, 10])
    b_fc2 = bias_variable([10])
    y_conv = tf.nn.softmax(tf.matmul(h_fc1_drop, W_fc2) + b_fc2, name="prediction")

    # 定义损失函数
    cross_entropy = tf.reduce_mean(-tf.reduce_sum(y_ * tf.log(y_conv), reduction_indices=[1]))
    train_step = tf.train.AdamOptimizer(1e-4).minimize(cross_entropy)

    correct_prediction = tf.equal(tf.argmax(y_conv, 1), tf.argmax(y_, 1))
    accuracy = tf.reduce_mean(tf.cast(correct_prediction, tf.float32))

    init = tf.global_variables_initializer()
    sess.run(init)

    # 保存变量
    graph = tf.Graph()
    # 固化需要用到的变量
    # 这里的输出需要加上name属性 需要指定output_node_names，可以多个
    constant_graph = graph_util.convert_variables_to_constants(sess, sess.graph_def, ['prediction'])
    with tf.gfile.FastGFile("./CNNData.pb", mode='wb') as f:
        f.write(constant_graph.SerializeToString())

    for i in range(20000):
        batch = mnist.train.next_batch(50)
        if i % 100 == 0:
            train_accuracy = sess.run(accuracy, feed_dict={x: batch[0], y_: batch[1], keep_prob: 0.1})
            print('step %d, training accuracy %g' % (i, train_accuracy))

        sess.run(train_step, feed_dict={x: batch[0], y_: batch[1], keep_prob: 0.5})
    print('test accuracy %g' % sess.run(accuracy, feed_dict={x: mnist.test.images, y_: mnist.test.labels, keep_prob: 1.0}))


if __name__ == '__main__':
    # 调用方法: parameters and structure
    train(1, 3, 4)
    pass
