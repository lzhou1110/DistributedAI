# _*_ coding:utf-8 _*_
import numpy as np
from functools import reduce
import math


class Conv2D(object):
    # [batch_size, 28, 28, 1], 12, 5, 1, "VALID", wt
    def __init__(self, shape, output_channels, ksize=3, stride=1, method='VALID', wt=0, isFirstConv = 2):
        '''
        输入数据的shape, filter的个数，filter的尺寸， 步长， 是否通过padding输出原尺寸大小, wt 模型参数， isFirstConv是否第二层参数,默认取二
        '''
        self.input_shape = shape                    # [batch_size, 28, 28, 1]
        self.output_channels = output_channels      # define  kernel mounts
        self.input_channels = shape[-1]             # 输入的图像是几维
        self.batchsize = shape[0]                   # 每一批多少个数量
        self.stride = stride                        # 步长
        self.ksize = ksize                          # filter的size
        self.method = method                        # padding 是否补0
        self.wt = wt                                # 模型的参数
        self.isFirstConv = isFirstConv
        # 设置参数初始化，进行normalization, 加速收敛
        weights_scale = math.sqrt(reduce(lambda x, y: x * y, shape) / self.output_channels)
        self.weights = np.random.standard_normal(
            (ksize, ksize, self.input_channels, self.output_channels)) / weights_scale
        self.bias = np.random.standard_normal(self.output_channels) / weights_scale
        # TODO(fallenkliu@gmail.com): if global train times > 1 change weights and bias
        # 判断传入参数 非0则表示 修改第一层的参数
        if wt != 0 and self.isFirstConv == 1:
            self.weights = wt[0]
            self.bias = wt[1]

        if method == 'VALID':
            self.eta = np.zeros((shape[0], int((shape[1] - ksize + 1) / self.stride), int((shape[1] - ksize + 1) / self.stride),
             self.output_channels))

        if method == 'SAME':
            self.eta = np.zeros((shape[0], shape[1]//self.stride, shape[2]//self.stride,self.output_channels))

        self.w_gradient = np.zeros(self.weights.shape)
        self.b_gradient = np.zeros(self.bias.shape)
        self.output_shape = self.eta.shape

        if (shape[1] - ksize) % stride != 0:
            print('input tensor width can\'t fit stride')
        if (shape[2] - ksize) % stride != 0:
            print('input tensor height can\'t fit stride')

    def forward(self, x):
        # im2col 优化方法
        # kernel weights reshape from self.
        # 1. 首先我们将卷积层的参数weights通过ndarray自带的reshape方法reshape到上图中Kernal Matrix的形状。
        col_weights = self.weights.reshape([-1, self.output_channels])
        # 2. 根据self.method，选择是否对输入的数据进行padding,这里我们调用 \[np.pad()\]  方法，
        #    对我们的输入数据四维ndarray的第二维和第三维分别padding上与卷积核大小相匹配的0元素
        if self.method == 'SAME':
            # 张量 x 周围补0
            x = np.pad(x, (
                (0, 0), (self.ksize / 2, self.ksize / 2), (self.ksize / 2, self.ksize / 2), (0, 0)),
                             'constant', constant_values=0)
        # 3. 声明一个list用于存储转换为column的image,在backward中我们还会用到
        self.col_image = []
        conv_out = np.zeros(self.eta.shape)
        # 4. 对于batch中的每一个数据，分别调用im2col方法，将该数据转化为的Input features(Matrix),
        # 然后调用 \[np.dot()\]  完成矩阵乘法得到Output features(Matrix), reshape输出的shape,填充到输出数据中。
        for i in range(self.batchsize):
            img_i = x[i][np.newaxis, :]
            self.col_image_i = im2col(img_i, self.ksize, self.stride)
            conv_out[i] = np.reshape(np.dot(self.col_image_i, col_weights) + self.bias, self.eta[0].shape)
            self.col_image.append(self.col_image_i)

        self.col_image = np.array(self.col_image)
        return conv_out

    # 反向传播BP: 1. 得到最后的结果loss函数 2. 目标是对loss优化(loss = func(x,y, w))=>求极值=>求导(loss对参数w)寻找最小值的w的区间
    # 因此: 通过loss对w求导，获取loss最小值，也是导数=0处；由此获得w的值
    #
    def gradient(self, eta):
        self.eta = eta
        col_eta = np.reshape(eta, [self.batchsize, -1, self.output_channels])

        for i in range(self.batchsize):
            self.w_gradient += np.dot(self.col_image[i].T, col_eta[i]).reshape(self.weights.shape)
        self.b_gradient += np.sum(col_eta, axis=(0, 1))

        # deconv of padded eta with flippd kernel to get next_eta
        if self.method == 'VALID':
            pad_eta = np.pad(self.eta, (
                (0, 0), (self.ksize - 1, self.ksize - 1), (self.ksize - 1, self.ksize - 1), (0, 0)),
                             'constant', constant_values=0)

        if self.method == 'SAME':
            pad_eta = np.pad(self.eta, (
                (0, 0), (self.ksize / 2, self.ksize / 2), (self.ksize / 2, self.ksize / 2), (0, 0)),
                             'constant', constant_values=0)

        col_pad_eta = np.array([im2col(pad_eta[i][np.newaxis, :], self.ksize, self.stride) for i in range(self.batchsize)])
        flip_weights = np.flipud(np.fliplr(self.weights))
        col_flip_weights = flip_weights.reshape([-1, self.input_channels])
        next_eta = np.dot(col_pad_eta, col_flip_weights)
        next_eta = np.reshape(next_eta, self.input_shape)
        return next_eta

    def backward(self, alpha=0.00001, weight_decay=0.0004):
        # weight_decay = L2 regularization
        self.weights *= (1 - weight_decay)
        self.bias *= (1 - weight_decay)
        self.weights -= alpha * self.w_gradient
        self.bias -= alpha * self.bias

        self.w_gradient = np.zeros(self.weights.shape)
        self.b_gradient = np.zeros(self.bias.shape)


# def im2col(image, ksize, stride):
#     # image is a 4d tensor([batchsize, width ,height, channel])
#     img = []
#     for b in range(image.shape[0]):
#         per_image_col = []
#         for i in range(0, image.shape[1] - ksize + 1, stride):
#             for j in range(0, image.shape[2] - ksize + 1, stride):
#                 col = image[0, i:i + ksize, j:j + ksize, :].reshape([-1])
#                 per_image_col.append(col)
#         per_image_col = np.array(per_image_col)
#         img.append(per_image_col)
#     return np.array(img)

# im2col优化方法：通过将图像展开，使得卷积运算可以变成两个矩阵乘法
def im2col(image, ksize, stride):
    # image is a 4d tensor([batchsize, width ,height, channel])
    image_col = []
    for i in range(0, image.shape[1] - ksize + 1, stride):
        for j in range(0, image.shape[2] - ksize + 1, stride):
            col = image[:, i:i + ksize, j:j + ksize, :].reshape([-1])
            image_col.append(col)
    image_col = np.array(image_col)

    return image_col


if __name__ == "__main__":
    # img = np.random.standard_normal((2, 32, 32, 3))
    img = np.ones((1, 32, 32, 3))
    img *= 2
    conv = Conv2D(img.shape, 12, 3, 1)
    next = conv.forward(img)
    next1 = next.copy() + 1
    conv.gradient(next1-next)
    print(conv.w_gradient)
    print(conv.b_gradient)
    conv.backward()
