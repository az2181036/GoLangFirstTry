# -*- coding:utf-8 -*-
import matplotlib
import numpy as np
import matplotlib.pyplot as plt
matplotlib.rcParams['font.sans-serif'] = ['SimHei']
x = [100, 300, 500, 1000]
node_stf = [4961, 12934, 21433, 30081]
node_edf = [5459, 15072, 24703, 32603]
node_total = [5521, 17319, 32554, 59174]

violation_stf = [150, 1480, 3545, 17722]
violation_edf = [21, 248, 391, 3996]

def node_cmp(y1, y2, y3):
        bar_width = 0.3
        plt.bar(x=np.arange(len(x))-bar_width /2, height=y1, label='SJF',
                color='steelblue', alpha=0.8, width=bar_width)
        plt.bar(x=np.arange(len(x))+bar_width/2, height=y2, label='EDF',
                color='darkorange', alpha=0.8, width=bar_width)
        plt.title("SJF与EDF节点性能对比结果")
        plt.xlabel("节点最大任务数")
        plt.ylabel("任务数量")
        plt.legend()
        plt.plot(np.arange(len(x)), y3, marker='*', label='总任务数', color='orangered')
        plt.xticks(range(len(x)), labels = ['100', '300', '500', '1000'])
        plt.legend(loc="upper left")
        plt.savefig("stf_edf1.png")
        plt.show()

def violation_cmp(y1, y2):
        plt.plot(np.arange(len(x)), y1, marker='*', label='SJF', color='red')
        plt.plot(np.arange(len(x)), y2, marker='+', label='EDF', color='blue')
        plt.title("SJF与EDF全局性能对比结果")
        plt.xlabel("节点最大任务数")
        plt.ylabel("违反deadline约束的任务数量")
        plt.xticks(range(len(x)), labels=['100', '300', '500', '1000'])
        plt.legend(loc="upper left")
        plt.savefig("stf_edf2.png")
        plt.show()

if __name__ == '__main__':
        node_cmp(node_stf, node_edf, node_total)
        violation_cmp(violation_stf, violation_edf)