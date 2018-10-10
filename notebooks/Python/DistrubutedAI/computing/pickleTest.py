import tensorflow as tf
import os
import pickle
import numpy as np
# 读取本地所有模型
def readModel(path):
    init_op = tf.global_variables_initializer()
    # downModelLoadPath = path
    downModelLoadPath = "./trainedModel/"
    saver = tf.train.import_meta_graph(downModelLoadPath + "cpk.meta")

    # Launch the gtrph
    with tf.Session() as sess:
        # create dir for model saver
        # model_dir = "mnist"
        # model_name = "cpk"
        # model_path=os.path.join(model_dir,model_name)
        # saver.restore(sess,model_path)
        init = sess.run(init_op)
        model_dir = "./trainedModel"
        model_name = "cpk"
        model_path = os.path.join(model_dir, model_name)
        saver.restore(sess, model_path)
        conv1 = sess.run(tf.get_default_graph().get_tensor_by_name("conv1:0"))
        print(conv1)

def readAllEpoch():
    with open("allEpoch.pickle", 'rb') as f:
        data = pickle.load(f)
    data["out"] = np.transpose(data["out"])
    print("out shape:",data["out"].shape)
    print("out dataProvider[out][0]", data["out"])
    print("out_b shape:", data["out_b"].shape)

if __name__ == "__main__":
    # readModel("trainedModel")
    readAllEpoch()