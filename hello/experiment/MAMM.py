# -*- coding:utf-8 -*-
import matplotlib
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
matplotlib.rcParams['font.sans-serif'] = ['SimHei']

e1 = pd.read_csv('../../experiment_edf.csv')
e2 = pd.read_csv('../../experiment_edf_v2.csv')
e3 = pd.read_csv('../../experiment_edf_sort.csv')
e4 = pd.read_csv('../../experiment_edf_climb.csv')
x = e1['Tasks Number Limit']
y1 = e2['Migration Number']
y2 = e3['Migration Number']
y3 = e1['Migration Number']
y4 = e4['Migration Number']

plt.plot(np.arange(len(x)), y1, marker='s', label='OAMM')
plt.plot(np.arange(len(x)), y3, marker='x', label='BM')
plt.plot(np.arange(len(x)), y2, marker='o', label='BM_Sort')
plt.plot(np.arange(len(x)), y3, marker='+', label='BM_Climb')
plt.title("OAMM的改进效果")
plt.xlabel("节点最大任务数")
plt.ylabel("节点任务迁移")
plt.xticks(range(len(x)), labels=x)
plt.legend(loc="upper left")
plt.savefig("MAMM.png")
plt.show()