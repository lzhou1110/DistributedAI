import tensorflow as tf
from tensorflow.python.framework import graph_util

with tf.Session(graph=tf.Graph()) as sess:
    x = tf.placeholder(tf.int32, name='input_x')
    y = tf.placeholder(tf.int32, name='input_y')
    b = tf.Variable(1, name='b')
    xy = tf.multiply(x, y)
    op = tf.add(xy, b, name='sum')
    # 初始化变量
    sess.run(tf.global_variables_initializer())

    # 测试 OP
    print(sess.run(op, feed_dict = {x: 10, y: 3}))

    # 固化需要用到的变量
    # 这里的输出需要加上name属性 需要指定output_node_names，可以多个
    constant_graph = graph_util.convert_variables_to_constants(sess, sess.graph_def, ['sum'])

    # 写入序列化的PB文件
    with tf.gfile.FastGFile("./test.pb", mode='wb') as f:
        f.write(constant_graph.SerializeToString())


with tf.Session() as sess:
    with tf.gfile.FastGFile("./test.pb", 'rb') as f:
        graph_def = tf.GraphDef()
        graph_def.ParseFromString(f.read()) #加载图
        # 从图上读取张量（第一种方法），同时把图设为默认图
        input_x,input_y,op=tf.import_graph_def(graph_def, return_elements=["input_x:0","input_y:0","sum:0"])

    # 需要有一个初始化的过程
    sess.run(tf.global_variables_initializer())

    print(sess.run(op,  feed_dict={input_x: 5, input_y: 5}))