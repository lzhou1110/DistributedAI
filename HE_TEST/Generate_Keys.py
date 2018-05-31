import numpy as np

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

# c' = Mc* = M(bitvecotr(c)) = M(bitvecotr((wI)*x))
def encrypt(T, x, w, l):
    return keySwitch(keySwitchMatrix(np.eye(len(x)), T, l), w * x, l)

# 产生秘钥

T_keys = list()
for i in range(max_dim):
    if i > low_dim and i != equal:
        T_keys.append(list)
        continue
    T_keys.append(np.random.rand(i + 1, 1))

# S' = [I, T]
def getSecretKey(T):
    assert (T.ndim == 2)
    I = np.eye(len(T))  # num rows
    return hCat(I, T)

def one_way_encrypt_vector(vector, scaling_factor=1000):
    padded_vector = np.random.rand(len(vector) + 1)
    padded_vector[0:len(vector)] = vector

    vec_len = len(padded_vector)
    print(vec_len - 2)
    print("M_keys:",len(M_keys))
    M_temp = (M_keys[vec_len - 2].T * padded_vector * scaling_factor / (vec_len - 1)).T
    e_vector = innerProd(c_ones[vec_len - 2], c_ones[vec_len - 2], M_temp, l)
    return e_vector.astype('int')

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

def one_way_encrypt_vector(vector, scaling_factor=1000):
    padded_vector = np.random.rand(len(vector) + 1)
    padded_vector[0:len(vector)] = vector

    vec_len = len(padded_vector)
    print(vec_len - 2)
    print("M_keys:",len(M_keys))
    M_temp = (M_keys[vec_len - 2].T * padded_vector * scaling_factor / (vec_len - 1)).T
    e_vector = innerProd(c_ones[vec_len - 2], c_ones[vec_len - 2], M_temp, l)
    return e_vector.astype('int')


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

def int2bin(x):
    s = list()
    mod = 2
    while (x > 0):
        s.append(int(x % 2))
        x = int(x / 2)
    return np.array(list(reversed(s))).astype('int64')

# c' = Mc* = M(bitvecotr(c)) = M(bitvecotr((wI)*x))
def encrypt(T, x, w, l):
    return keySwitch(keySwitchMatrix(np.eye(len(x)), T, l), w * x, l)

# M, notice that it has aBound and eBound here
def keySwitchMatrix(S, T, l):
    S_star = getBitMatrix(S, l)
    A = getRandomMatrix(T.shape[1], S_star.shape[1], aBound)
    E = getRandomMatrix(S_star.shape[0], S_star.shape[1], eBound)
    return vCat(S_star + E - T.dot(A), A)

# Random matrix, |matrix| <= 1
def getRandomMatrix(row, col, bound):
    A = np.zeros((row, col))
    for i in range(row):
        for j in range(col):
            A[i][j] = np.random.randint(bound)
    return A

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


