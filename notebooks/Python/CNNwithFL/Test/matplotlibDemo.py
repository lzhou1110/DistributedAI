#_*_ coding:utf-8 _*_
import matplotlib.pyplot as plt
import numpy as np
# 简单的绘图
x = np.linspace(0, 2 * np.pi, 50)
plt.plot(x, np.sin(x)) # 如果没有第一个参数 x，图形的 x 坐标默认为数组的索引
plt.show() # 显示图形