#!/usr/bin/env python3
# _*_ coding:utf-8 _*_
import numpy as np
from layers.base_convolution import Conv2D
from layers.fullyconnect import FullyConnect
from layers.pooling import MaxPooling, AvgPooling
from layers.loss import Softmax
from layers.activator import Relu

import time
import struct
from glob import glob

def load_mnist(path, kind='train'):
    """Load MNIST data from path`"""
    images_path = glob('./%s/%s*3-ubyte' % (path, kind))[0]
    labels_path = glob('./%s/%s*1-ubyte' % (path, kind))[0]
    with open(labels_path, 'rb') as lbpath:
        magic, n = struct.unpack('>II',
                                 lbpath.read(8))
        labels = np.fromfile(lbpath,
                             dtype=np.uint8)

    with open(images_path, 'rb') as imgpath:
        magic, num, rows, cols = struct.unpack('>IIII',
                                               imgpath.read(16))
        images = np.fromfile(imgpath,
                             dtype=np.uint8).reshape(len(labels), 784)
    # 将images和labels分类 1, 2, .. 0
    # 循环 labels,新建10个images,labels,判断是都等于某个数， 筛选建立shape为(10, images),(10, labels)
    # listdata = []
    # for j in range(len(labels)):
    #     listdata.append(images[j])
    #TODO(fallenkliu@gmail.com): change the way of getting data
    imagesOne = []
    labelsOne = []

    imagesTwo = []
    labelsTwo = []

    imagesThree = []
    labelsThree = []

    imagesFour = []
    labelsFour = []

    imagesFive = []
    labelsFive = []

    imagesSix = []
    labelsSix = []

    imagesSeven = []
    labelsSeven = []

    imagesEight = []
    labelsEight = []

    imagesNine = []
    labelsNine = []

    imagesZero = []
    labelsZero = []
    categoryImages = [imagesZero, imagesOne, imagesTwo, imagesThree, imagesFour, imagesFive, imagesSix, imagesSeven,
                      imagesEight, imagesNine]
    categoryLabels = [labelsZero, labelsOne, labelsTwo, labelsThree, labelsFour, labelsFive, labelsSix, labelsSeven,
                      labelsEight, labelsNine]
    for i in range(len(labels)):
        num = labels[i]
        categoryImages[num].append(images[i])
        categoryLabels[num].append(labels[i])

    return categoryImages, categoryLabels

# load mnist data
categoryImages, categoryLabels = load_mnist('./data/mnist')
test_categoryImages, test_categoryLabels = load_mnist('./data/mnist', 't10k')
batch_size = 64


def trainAndTest(k, wt, batch_size, epoch):
    '''
    :param k: client k 编号
    :param wt: 全局参数
    :param batch_size: 每一个batch大小
    :param epoch: 整体训练次数
    :return: 训练好的参数wt， 所用训练的数据集个数nk(根据k的大小)
    '''
    # 构建 两个 卷积层和pool层
    # 输入数据的shape, 卷积核的个数，卷积核的尺寸， 步长， 是否输出原尺寸大小
    conv1 = Conv2D([batch_size, 28, 28, 1], 12, 5, 1, "VALID", wt, 1)
    relu1 = Relu(conv1.output_shape)
    pool1 = MaxPooling(relu1.output_shape)
    conv2 = Conv2D(pool1.output_shape, 24, 3, 1, "VALID", wt)
    relu2 = Relu(conv2.output_shape)
    pool2 = MaxPooling(relu2.output_shape)
    fc = FullyConnect(pool2.output_shape, 10)
    sf = Softmax(fc.output_shape)

    # print out format
    print("INPUT:28*28*1", "memory:%d*%d*%d*%d"%(batch_size, 28, 28, 1), "weights:0")
    print("CONV1:",conv1.output_shape)
    print("POOL1:", pool1.output_shape)
    print("CONV2:", conv2.output_shape)
    print("POOL2:", pool2.output_shape)
    print("fc:", fc.output_shape)


    # 根据k 数据分类
    images = np.asarray(categoryImages[k])
    labels = np.asarray(categoryLabels[k])
    test_images = np.asarray(test_categoryImages[k])
    test_labels = np.asarray(test_categoryLabels[k])
    #
    nk = len(labels)

    # 定义epoch
    for epoch in range(epoch):
        learning_rate = 1e-5  # 学习率
        batch_loss = 0  # 损失函数
        batch_acc = 0  # 正确率
        val_acc = 0  # 总的正确率
        val_loss = 0  # 损失率


        # train
        train_acc = 0
        train_loss = 0

        for i in range(images.shape[0] // batch_size):
            # one piece of image and label 28*28; batch_size is 64
            img = images[i * batch_size:(i + 1) * batch_size].reshape([batch_size, 28, 28, 1])
            label = labels[i * batch_size:(i + 1) * batch_size]
            conv1_out = relu1.forward(conv1.forward(img))
            pool1_out = pool1.forward(conv1_out)
            conv2_out = relu2.forward(conv2.forward(pool1_out))
            pool2_out = pool2.forward(conv2_out)
            fc_out = fc.forward(pool2_out)

            # print(i, 'fc_out', fc_out)

            batch_loss += sf.cal_loss(fc_out, np.array(label))
            train_loss += sf.cal_loss(fc_out, np.array(label)) # 总的loss

            # image和label匹配并且
            for j in range(batch_size):
                if np.argmax(sf.softmax[j]) == label[j]:
                    # print(label[j])
                    batch_acc += 1
                    train_acc += 1

            sf.gradient()
            conv1.gradient(relu1.gradient(pool1.gradient(
                conv2.gradient(relu2.gradient(pool2.gradient(
                    fc.gradient(sf.eta)))))))

            if i % 1 == 0:
                # TO-DO 更新参数
                fc.backward(alpha=learning_rate, weight_decay=0.0004)
                conv2.backward(alpha=learning_rate, weight_decay=0.0004)
                conv1.backward(alpha=learning_rate, weight_decay=0.0004)
                conv1.backward()
                if i % 50 == 0:
                    print(time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) + \
                          "  epoch: %d ,  batch: %5d , avg_batch_acc: %.4f learning_rate %f" % (epoch,
                                                                                                     i, batch_acc / float(
                              batch_size), learning_rate))
                    # print(time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) + \
                    #       "  epoch: %d ,  batch: %5d " % (epoch, i))

                batch_loss = 0
                batch_acc = 0

        print(time.strftime("%Y-%m-%d %H:%M:%S",
                            time.localtime()) + "  epoch: %5d ,client train label:%s, train_acc: %.4f " % (
                  epoch, labels[0], train_acc / float(images.shape[0])))

        # validation
        for i in range(test_images.shape[0] // batch_size):
            img = test_images[i * batch_size:(i + 1) * batch_size].reshape([batch_size, 28, 28, 1])
            label = test_labels[i * batch_size:(i + 1) * batch_size]
            conv1_out = relu1.forward(conv1.forward(img))
            pool1_out = pool1.forward(conv1_out)
            conv2_out = relu2.forward(conv2.forward(pool1_out))
            pool2_out = pool2.forward(conv2_out)
            fc_out = fc.forward(pool2_out)
            val_loss += sf.cal_loss(fc_out, np.array(label))

            for j in range(batch_size):
                if np.argmax(sf.softmax[j]) == label[j]:
                    val_acc += 1

        print(time.strftime("%Y-%m-%d %H:%M:%S",
                            time.localtime()) + "  epoch: %5d ,client test label:%s, val_acc: %.4f  avg_val_loss: %.4f" % (
                  epoch, test_labels[0], val_acc / float(test_images.shape[0]), val_loss / test_images.shape[0]))

    wt = (conv1.weights, conv1.bias)
    return wt, nk



# client k
def clientUpdate(k, wt):
    '''
    :param k: client k-编号
    :param wt: 传入的全局参数
    :return: newWkt, nk: 更新后的wt, 本地训练数据集的个数
    '''
    # 1. 数据集分块 根据Pk按batch_size为B进行分块
    B = 64
    batch_size = B
    epoch = 1
    print("client %d will begin training" % k)
    # 2. 传入参数， 训练数据(根据k进行识别编号)，测试数据, 得到返回参数
    newWkt, nk = trainAndTest(k, wt, batch_size, epoch)
    #TODO(fallenkliu@gmail.com): setting parameters
    print("当前参数newWkt:\n", newWkt)
    print("当前client训练集个数:", nk)
    print("=================client %d train finished=================" % k)
    return newWkt, nk

# server execute
def server(w0 = 0):
    '''
    :param w0: 训练的参数
    :return:
    '''
    # 1. initialize w0
    # w0 = np.zeros()
    # TODO(fallekliu@gmail.com): global train times
    # 2. 总的全局训练次数
    t = 2
    wt = w0 # 全局参数学习
    sumN = 0 # 全局总的训练集个数
    for i in range(t):
        # 3. 选择一批client 集合
        # m←max(C ·K, 1): m 为c个clients和
        # St ←(random set of m clients) : 每次从m个中选取St个clients
        # st = np.array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
        st = np.array([0, 1])
        listWAndnk = [] # 当前一次总的训练得到的参数

        # 所有clients 训练的训练数据集 当前循环t
        n = 0
        for k in range(len(st)):
            # newWt为当前client k的训练所得的参数；nk为client k的本地的训练集个数；
            newWkt, nk = clientUpdate(k, wt)
            n += nk
            listWAndnk.append((newWkt, nk))
        # 4. 更新w值(加权平均)
        # Wt+1 <-- sum(Nk/n * Wt+1): 即n个clients的参数加权求和得出全局的Wt+1
        # list 转为np.array
        listWAndnk = np.asarray(listWAndnk)
        weights = np.zeros_like(listWAndnk[k][0][0])
        bias = np.zeros_like(listWAndnk[k][0][1])

        # accumulate weights and bias
        if wt == 0:
            wt = np.zeros_like(np.array([weights, bias]))
        # 当前更新每次迭代的参数
        for k in range(len(listWAndnk)):
            weights += np.asarray(listWAndnk[k][0][0])*listWAndnk[k][1]/n
            bias += np.asarray(listWAndnk[k][0][1])*listWAndnk[k][1]/n
        # 全局总的训练集
        sumN +=n

        # 更新全局参数 累计; 之前的参数*((sumN-n) /sum)+ 现在的参数*(n/(sumN))
        lastWt = np.array([wt[0]*((sumN-n)/sumN), wt[1]*((sumN-n)/sumN)])
        currentWt = np.array([weights*(n*sumN), bias*(n/sumN)])
        wt =  lastWt + currentWt
        print("当前更新的全局参数:\n", wt)
        print("=================总的循环次数t:%d=================" % (i+1))

    print("所有迭代更新完成的全局参数:\n", wt)
if __name__ == "__main__":

    server(0)