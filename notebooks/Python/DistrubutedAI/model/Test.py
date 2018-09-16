from tensorflow.examples.tutorials.mnist import input_data
import numpy as np

def readData():
    imageData2 = np.loadtxt('/Users/liulifeng/Desktop/Work/mnist_data/image/images2')

    print(imageData2.shape)
    pass

if __name__ == '__main__':
    readData()
    pass