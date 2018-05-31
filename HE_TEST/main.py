import numpy as np
import copy
import sys
import os
import struct

def get_T(n):
    n_prime = n + 1
    T = (10 * np.random.rand(n,n_prime - n)).astype('int')
    return T

x = np.array([0,1,2,5])

m = len(x)
n = m
T = get_T(n)

#-----------------------------------

# horizontal concat
def hCat(A, B):
    return np.concatenate((A, B), 1)


# vertical concat two matrix [A, B]
def vCat(A, B):
    return np.concatenate((A, B), 0)


# vec(M)
def vectorize(M):
    ans = np.zeros((len(M) * len(M[0]), 1))
    for i in range(len(M)):
        for j in range(len(M[0])):
            ans[i * len(M[0]) + j][0] = M[i][j]
    return ans


# Random matrix, |matrix| <= 1
def getRandomMatrix(row, col, bound):
    A = np.zeros((row, col))
    for i in range(row):
        for j in range(col):
            A[i][j] = np.random.randint(bound)
    return A


# c*
def getBitVector(c, l):
    m = len(c)
    c_star = np.zeros(l * m, dtype='int64')
    for i in range(m):
        local_c = int(c[i])
        if (local_c < 0):
            local_c = -local_c
        b = int2bin(local_c)
        if (c[i] < 0):
            b *= -1
        if (c[i] == 0):
            b *= 0
        c_star[(i * l) + (l - len(b)): (i + 1) * l] += b
    return c_star


# S*
def getBitMatrix(S, l):
    S_star = list()
    for i in range(l):
        S_star.append(S * 2 ** (l - i - 1))
    S_star = np.array(S_star).transpose(1, 2, 0).reshape(len(S), len(S[0]) * l)
    return S_star


# c' = Mc* = M(bitvecotr(c))
def keySwitch(M, c, l):
    c_star = getBitVector(c, l)
    return M.dot(c_star)


# S' = [I, T]
def getSecretKey(T):
    assert (T.ndim == 2)
    I = np.eye(len(T))  # num rows
    return hCat(I, T)


# M, notice that it has aBound and eBound here
def keySwitchMatrix(S, T, l):
    S_star = getBitMatrix(S, l)
    A = getRandomMatrix(T.shape[1], S_star.shape[1], aBound)
    E = getRandomMatrix(S_star.shape[0], S_star.shape[1], eBound)
    return vCat(S_star + E - T.dot(A), A)


# c' = Mc* = M(bitvecotr(c)) = M(bitvecotr((wI)*x))
def encrypt(T, x, w, l):
    return keySwitch(keySwitchMatrix(np.eye(len(x)), T, l), w * x, l)


def decrypt(S, c, w):
    Sc = S.dot(c)
    return (Sc / w).astype('float').round().astype('int')

# Addition
def addVectors(c1, c2):
    return c1 + c2

# Linear transformation
def linearTransformClient(G, S, T, l):
    return keySwitchMatrix(G.dot(S), T, l)

def linearTransform(M, c, l):
    return M.dot(getBitVector(c, l)).astype('int64')

# Inner Product
def copyRows(row, numrows):
    ans = np.zeros((numrows, len(row[0])))
    for i in range(len(ans)):
        for j in range(len(ans[0])):
            ans[i][j] = row[0][j]
    return ans

def innerProdClient(T,l):
    S = getSecretKey(T)
    tvsts = vectorize(S.T.dot(S)).T
    mvsts = copyRows(tvsts, len(T))
    return keySwitchMatrix(mvsts,T,l)
#c1‘*c2‘ 算的是两个矢量的内积c1 = [1,2,3] c2 = [2,3,4] 结果：20
def innerProd(c1, c2, M,l):
    cc1 = np.zeros((len(c1),1))
    for i in range(len(c1)):
        cc1[i][0] = c1[i]
    cc2 = np.zeros((1, len(c2)))
    for i in range(len(c2)):
        cc2[0][i] = c2[i]
    cc = vectorize(cc1.dot(cc2))
    bv = getBitVector((cc / w).round().astype('int64'),l)
    return M.dot(bv)


def one_way_encrypt_vector(vector, scaling_factor=1000):
    padded_vector = np.random.rand(len(vector) + 1)
    padded_vector[0:len(vector)] = vector

    vec_len = len(padded_vector)
    print(vec_len - 2)
    print("M_keys:",len(M_keys))
    M_temp = (M_keys[vec_len - 2].T * padded_vector * scaling_factor / (vec_len - 1)).T
    e_vector = innerProd(c_ones[vec_len - 2], c_ones[vec_len - 2], M_temp, l)
    return e_vector.astype('int')


def load_linear_transformation(syn0_text, scaling_factor=1000):
    syn0_text *= scaling_factor
    return linearTransformClient(syn0_text.T, getSecretKey(T), T, l)


def s_decrypt(vec):
    return decrypt(getSecretKey(T_keys[len(vec) - 2]), vec, w)


def add_vectors(x, y, scaling_factor=10000):
    return x + y


def transpose(syn1):
    rows = len(syn1)
    cols = len(syn1[0]) - 1

    max_rc = max(rows, cols)

    syn1_c = list()
    for i in range(len(syn1)):
        tmp = np.zeros(max_rc + 1)
        tmp[:len(syn1[i])] = syn1[i]
        syn1_c.append(tmp)

    syn1_c_transposed = list()

    for row_i in range(cols):
        syn1t_column = innerProd(syn1_c[0], v_onehot[max_rc - 1][row_i], M_onehot[max_rc - 1][0], l) / scaling_factor
        for col_i in range(rows - 1):
            syn1t_column += innerProd(syn1_c[col_i + 1], v_onehot[max_rc - 1][row_i], M_onehot[max_rc - 1][col_i + 1],
                                      l) / scaling_factor

        syn1_c_transposed.append(syn1t_column[0:rows + 1])

    return syn1_c_transposed


def int2bin(x):
    s = list()
    mod = 2
    while (x > 0):
        s.append(int(x % 2))
        x = int(x / 2)
    return np.array(list(reversed(s))).astype('int64')

#列表不能运算
def sigmoid_func(x):
    return 1 / (1 + np.exp(-x))

def ReLU(vec):
    res = list()
    for num in vec:
        if num <= 0:
            num = 0
        res.append(num)
    return res

# def sigmoid(layer_2_c):
#     out_rows = list()
#     for position in range(len(layer_2_c) - 1):
#         M_position = M_onehot[len(layer_2_c) - 2][0]
#         #layer_2_index_c是矢量中第position个位置的值
#         #len(layer_2_c) - 2从v_onehot中取出相应维度的加密后的单位矩阵
#         layer_2_index_c = innerProd(layer_2_c, v_onehot[len(layer_2_c) - 2][position], M_position, l) / scaling_factor
#         x = layer_2_index_c
#         x2 = innerProd(x, x, M_position, l) / scaling_factor#x^2
#         x3 = innerProd(x, x2, M_position, l) / scaling_factor#x^3
#         x5 = innerProd(x3, x2, M_position, l) / scaling_factor#x^5
#         x7 = innerProd(x5, x2, M_position, l) / scaling_factor#x^7
#         xs = copy.deepcopy(v_onehot[4][0])
#         xs[1] = x[0]
#         xs[2] = x3[0]
#         xs[3] = x5[0]
#         xs[4] = x7[0]
#         #out = mat_mul_forward(xs, H_sigmoid[0:2], scaling_factor)
#         out = innerProd(xs,H_sigmoid[0],M_onehot[len(xs) - 2][0],l) / float(scaling_factor)
#         out_rows.append(out)
#     return transpose(out_rows)[0]


def mat_mul_forward(layer_1, syn1, scaling_factor):
    input_dim = len(layer_1)
    output_dim = len(syn1)
    buff = np.zeros(max(output_dim + 1, input_dim + 1))
    buff[0:len(layer_1)] = layer_1
    layer_1_c = buff
    syn1_c = list()
    for i in range(len(syn1)):
        buff = np.zeros(max(output_dim + 1, input_dim + 1))
        buff[0:len(syn1[i])] = syn1[i]
        syn1_c.append(buff)
    layer_2 = innerProd(syn1_c[0], layer_1_c, M_onehot[len(layer_1_c) - 2][0], l) / float(scaling_factor)
    for i in range(len(syn1) - 1):
        layer_2 += innerProd(syn1_c[i + 1], layer_1_c, M_onehot[len(layer_1_c) - 2][i + 1], l) / float(scaling_factor)
    return layer_2[0:output_dim + 1]

#加密后的矢量乘以矩阵
def vector_mat_mul(layer_1, syn1, scaling_factor):
    t = transpose(syn1)
    layer_2 = list()
    for i in range(len(t)):
        layer_2.append(innerProd(t[i], layer_1, M_onehot[len(layer_1) - 2][0], l) / float(scaling_factor))
    return layer_2

#解密加密后的矢量
def vector_decrypt(vector,scaling_factor = 1000):
    res = list()
    for i in range(len(vector)):
        res.append((s_decrypt(vector[i]))[0] / scaling_factor)
    return res

def elementwise_vector_mult(x, y, scaling_factor):
    y = [y]
    one_minus_layer_1 = transpose(y)
    outer_result = list()
    for i in range(len(x) - 1):
        outer_result.append(mat_mul_forward( x * onehot[len(x) - 1][i], y, scaling_factor))
    return transpose(outer_result)[0]



def outer_product(x, y):
    flip = False
    if (len(x) < len(y)):
        flip = True
        tmp = x
        x = y
        y = tmp
    y_matrix = list()
    for i in range(len(x) - 1):
        y_matrix.append(y)
    y_matrix_transpose = transpose(y_matrix)
    outer_result = list()
    for i in range(len(x) - 1):
        outer_result.append(mat_mul_forward(x * onehot[len(x) - 1][i], y_matrix_transpose, scaling_factor))
    if (flip):
        return transpose(outer_result)
    return outer_result

# 在服务端运行

l = 100

w = 2 ** 25


aBound = 10

tBound = 10

eBound = 10

max_dim = 785

low_dim = 10

high_dim = 780

equal = 783

scaling_factor = 1000

# 产生秘钥

T_keys = list()

for i in range(max_dim):
    if i > low_dim and i != equal:
        T_keys.append(list)
        continue
    T_keys.append(np.random.rand(i + 1, 1))


# 加密转换

M_keys = list()

for i in range(max_dim):
    if i > low_dim and i != equal:
        M_keys.append(list())
        continue
    M_keys.append(innerProdClient(T_keys[i], l))


M_onehot = list()

for h in range(max_dim):
    if h > low_dim and h != equal:
        M_onehot.append(list())
        continue
    i = h + 1

    buffered_eyes = list()

    for row in np.eye(i + 1):
        buffer = np.ones(i + 1)

        buffer[0:i + 1] = row

        buffered_eyes.append((M_keys[i - 1].T * buffer).T)

    M_onehot.append(buffered_eyes)

c_ones = list()

for i in range(max_dim):
    if i > low_dim and i != equal:
        c_ones.append(list())
        continue
    c_ones.append(encrypt(T_keys[i], np.ones(i + 1), w, l).astype('int'))

#帮助向量取值的一个加密过的单位矩阵
v_onehot = list()

onehot = list()

for i in range(max_dim):
    if i > low_dim and i != equal:
        onehot.append(list())
        continue

    eyes = list()

    eyes_txt = list()

    for eye in np.eye(i + 1):
        eyes_txt.append(eye)
        eyes.append(one_way_encrypt_vector(eye, scaling_factor))

    v_onehot.append(eyes)

    onehot.append(eyes_txt)


#
# #H_sigmoid矩阵是用于演算sigmoid泰勒展开的矩阵
#
# H_sigmoid_txt = np.zeros((5, 5))
#
# H_sigmoid_txt[0][0] = 0.5
#
# H_sigmoid_txt[0][1] = 0.25
#
# H_sigmoid_txt[0][2] = -1 / 48.0
#
# H_sigmoid_txt[0][3] = 1 / 480.0
#
# H_sigmoid_txt[0][4] = -17 / 80640.0
#
# H_sigmoid = list()
#
# for row in H_sigmoid_txt:
#     H_sigmoid.append(one_way_encrypt_vector(row))

#载入训练数据
#载入数据
train_images_idx3_ubyte_file = './Data/train-images.idx3-ubyte'
train_labels_idx1_ubyte_file = './Data/train-labels.idx1-ubyte'


test_images_idx3_ubyte_file = './Data/t10k-images.idx3-ubyte'
test_labels_idx1_ubyte_file = './Data/t10k-labels.idx1-ubyte'

#解码图片
def decode_idx3_ubyte(idx3_ubyte_file):
    bin_data = open(idx3_ubyte_file, 'rb').read()

    offset = 0
    fmt_header = '>IIII'
    magic_number, num_images, num_rows, num_cols = struct.unpack_from(fmt_header, bin_data, offset)
    #print ("magic:%d, count: %d, size: %d*%d" % (magic_number, num_images, num_rows, num_cols))

    image_size = num_rows * num_cols
    offset += struct.calcsize(fmt_header)
    fmt_image = '>' + str(image_size) + 'B'
    images = np.empty((num_images, num_rows, num_cols))
    for i in range(num_images):
        #if (i + 1) % 10000 == 0:
            #print("done %d" % (i + 1) + "pictures")
        images[i] = np.array(struct.unpack_from(fmt_image, bin_data, offset)).reshape((num_rows, num_cols))
        offset += struct.calcsize(fmt_image)
    return images

#将标签解码
def decode_idx1_ubyte(idx1_ubyte_file):
    bin_data = open(idx1_ubyte_file, 'rb').read()

    offset = 0
    fmt_header = '>ii'
    magic_number, num_images = struct.unpack_from(fmt_header, bin_data, offset)
    #print("magic:%d, num_images: %d zhang" % (magic_number, num_images))

    offset += struct.calcsize(fmt_header)
    fmt_image = '>B'
    labels = np.empty(num_images)
    for i in range(num_images):
        # if (i + 1) % 10000 == 0:
        #     print("done %d" % (i + 1) + "zhang")
        labels[i] = struct.unpack_from(fmt_image, bin_data, offset)[0]
        offset += struct.calcsize(fmt_image)
    return labels

#载入数据
def load_train_images(idx_ubyte_file=train_images_idx3_ubyte_file):
    return decode_idx3_ubyte(idx_ubyte_file)


def load_train_labels(idx_ubyte_file=train_labels_idx1_ubyte_file):
    return decode_idx1_ubyte(idx_ubyte_file)


def load_test_images(idx_ubyte_file=test_images_idx3_ubyte_file):
    return decode_idx3_ubyte(idx_ubyte_file)


def load_test_labels(idx_ubyte_file=test_labels_idx1_ubyte_file):
    return decode_idx1_ubyte(idx_ubyte_file)

#标准化
def normalize_data(ima):
    a_max = np.max(ima)
    a_min = np.min(ima)
    for j in range(ima.shape[0]):
        ima[j] = (ima[j] - a_min) / (a_max - a_min)
    return ima

def sigmoid(x):
    s=1/(1+np.exp(-x))
    return s
def imageToVector(image):
    v=np.reshape(image,[784,1])
    return v
def softmax(x):
    v=np.argmax(x)
    return v

def initialize_with_zeros(n_x, n_h, n_y):
    np.random.seed(2)

    W1 = np.random.uniform(-np.sqrt(6) / np.sqrt(n_x + n_h), np.sqrt(6) / np.sqrt(n_h + n_x), size=(n_h, n_x))
    W1_e = list()
    for row in W1:
        #print(row)
        W1_e.append(one_way_encrypt_vector(row,scaling_factor))

    b1 = np.zeros((n_h, 1))
    b1_e = list()
    for row in b1:
        b1_e.append(one_way_encrypt_vector(row,scaling_factor))

    W2 = np.random.uniform(-np.sqrt(6) / np.sqrt(n_y + n_h), np.sqrt(6) / np.sqrt(n_y + n_h), size=(n_y, n_h))
    W2_e = list()
    for row in W2:
        W2_e.append(one_way_encrypt_vector(row,scaling_factor))

    b2 = np.zeros((n_y, 1))
    b2_e = list()
    for row in b2:
        b2_e.append(one_way_encrypt_vector(row,scaling_factor))
    assert (W1.shape == (n_h, n_x))
    assert (b1.shape == (n_h, 1))
    assert (W2.shape == (n_y, n_h))
    assert (b2.shape == (n_y, 1))

    parameters = {"W1": W1,
                  "b1": b1,
                  "W2": W2,
                  "b2": b2}
    parameters_e = {"W1_e": W1_e,
                  "b1_e": b1_e,
                  "W2_e": W2_e,
                  "b2_e": b2_e}
    return parameters, parameters_e


def forward_propagation(X, parameters):
    W1 = parameters["W1"]
    b1 = parameters["b1"]
    W2 = parameters["W2"]
    b2 = parameters["b2"]
    W1_e = parameters["W1_e"]
    b1_e = parameters["b1_e"]
    W2_e = parameters["W2_e"]
    b2_e = parameters["b2_e"]
    # print W1,X,b1
    Z1 = np.dot(W1, X) + b1

    Z1_e = list()
    for i in range(len(W1_e)):
        Z1_e.append(innerProd(W1_e[i],transpose(X),M_onehot[len(row) - 2][0],l) + b1_e[i])
    print("b2_e:",b2_e)
    print("Z1_e",Z1_e)
    # A1=sigmoid(Z1)
    A1 = np.tanh(Z1)
    Z2 = np.dot(W2, A1) + b2
    A2 = sigmoid(Z2)
    # assert(A2.shape == (1, X.shape[1]))
    cache = {"Z1": Z1,
             "A1": A1,
             "Z2": Z2,
             "A2": A2}
    return A2, cache

#计算损失函数
def costloss(A2, Y, parameters):
    t = 0.00000000001
    #如果y = 0 值是-log(h(x))，如果y = 1值是-log(1-h(x))
    logprobs = np.multiply(np.log(A2 + t), Y) + np.multiply(np.log(1 - A2 + t), (1 - Y))
    cost = np.sum(logprobs, axis=0, keepdims=True) / A2.shape[0]
    return cost

#反向传播
def back_propagation(parameters,cache,X,Y):
    W1 = parameters["W1"]
    W2 = parameters["W2"]
    A1 = cache["A1"]
    A2 = cache["A2"]
    Z1 = cache["Z1"]

    dZ2 = A2-Y
    dW2 = np.dot(dZ2,A1.T)
    db2 = np.sum(dZ2,axis=1,keepdims=True)
    dZ1 = np.dot(W2.T,dZ2)*(1-np.power(A1,2))
    # dZ1=np.dot(W2.T,dZ2)*sigmoid(Z1)*(1-sigmoid(Z1))
    dW1 = np.dot(dZ1,X.T)
    db1 = np.sum(dZ1,axis=1,keepdims=True)
    grads = {"dW1": dW1,
             "db1": db1,
             "dW2": dW2,
             "db2": db2}
    return grads

#更新参数
def update_para(parameters, grads, learning_rate ):
    W1 = parameters["W1"]
    b1 = parameters["b1"]
    W2 = parameters["W2"]
    b2 = parameters["b2"]
    dW1 = grads["dW1"]
    db1 = grads["db1"]
    dW2 = grads["dW2"]
    db2 = grads["db2"]
    W1=W1-learning_rate*dW1
    b1=b1-learning_rate*db1
    W2=W2-learning_rate*dW2
    b2=b2-learning_rate*db2

    parameters = {"W1": W1,
                  "b1": b1,
                  "W2": W2,
                  "b2": b2}
    return parameters

#开始训练模型
if __name__ == '__main__':
    train_images = load_train_images()
    train_labels = load_train_labels()
    test_images = load_test_images()
    test_labels = load_test_labels()
    print(train_images[0])
    ii = 0
    n_x = 28 * 28
    n_h = 32#中间层神经元个数
    n_y = 10
    parameters, parameters_e = initialize_with_zeros(n_x, n_h, n_y)
    for i in range(20001):
        img_train = train_images[i]
        label_train1 = train_labels[i]
        label_train = np.zeros((10, 1))
        learning_rate = 0.001
        if i > 1000:
            learning_rate = learning_rate * 0.999
        label_train[int(train_labels[i])] = 1
        imgvector1 = imageToVector(img_train)
        # print("imgvector1: before transform: ",imgvector1[0])
        imgvector = normalize_data(imgvector1)
        # print("after transform: ",imgvector[0])
        #A2是为经过激活函数的输出
        A2, cache = forward_propagation(imgvector, parameters)
        pre_label = softmax(A2)

        costl = costloss(A2, label_train, parameters)
        grads = back_propagation(parameters, cache, imgvector, label_train)
        parameters = update_para(parameters, grads, learning_rate=learning_rate)
        grads["dW1"] = 0
        grads["dW2"] = 0
        grads["db1"] = 0
        grads["db2"] = 0
        if i % 1000 == 0:
            print("迭代%i次后的代价函数值:" % (i))
            print(costl[0])
        # print("ii de value: ",ii/50000.)
    # print('parameters',parameters["W1"],parameters["W2"],parameters["b1"],parameters["b2"])      # plt.imshow(train_images[i], cmap='gray')
    # print("cost : ",costl)
    # plt.show()
    for i in range(10000):
        img_train = test_images[i]
        vector_image = normalize_data(imageToVector(img_train))
        label_trainx = test_labels[i]
        aa2, xxx = forward_propagation(vector_image, parameters)
        predict_value = softmax(aa2)
        if predict_value == int(label_trainx):
            ii = ii + 1
    #正确识别数
    print("最终正确率:",float(ii) / 10000)


