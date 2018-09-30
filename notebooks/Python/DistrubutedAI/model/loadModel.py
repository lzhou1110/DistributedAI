import tensorflow as tf
import os
# 读取本地所有模型
def readModel(path):
    # downModelLoadPath = path
    saver = tf.train.Saver()
    init_op = tf.global_variables_initializer()
    with tf.Session() as sess:
        sess.run(init_op)
        model_dir = "./"+path
        model_name = "model_test"
        model_path = os.path.join(model_dir, model_name)
        saver.restore(sess, model_path)


if __name__ == "__main__":
    readModel("model_saved")