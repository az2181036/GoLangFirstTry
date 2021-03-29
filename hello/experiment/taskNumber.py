# -*- coding:utf-8 -*-
import matplotlib
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
matplotlib.rcParams['font.sans-serif'] = ['SimHei']

rst_v2 = pd.read_csv('../../experiment_edf_v2.csv')
rst_no_migration = pd.read_csv('../../experiment_EdgeOnlyNoMigration.csv')
rst_migration = pd.read_csv('../../experiment_EdgeOnlyWithMigration.csv')
rst_cluster = pd.read_csv('../../experiment_EdgeOnlyWithClusterings.csv')
rst_cloud = pd.read_csv('../../experiment_EdgeCloudWithMigration.csv')

x = [100, 300, 500, 800, 1000]
NoMigration = rst_no_migration['Amount of Tasks'] - rst_no_migration['Violation Tasks']
WithMigration = rst_migration['Amount of Tasks'] - rst_migration['Violation Tasks']
OnlyClustering = rst_cluster['Amount of Tasks'] - rst_cluster['Violation Tasks']
ECWM = rst_cloud['Amount of Tasks'] - rst_cloud['Violation Tasks']
ENCA_MAMM = rst_v2['Amount of Tasks'] - rst_v2['Violation Tasks']


bar_width = 0.15
plt.bar(x=np.arange(len(x))-bar_width*2, height=NoMigration, label='EWNO', alpha=0.8, width=bar_width)
plt.bar(x=np.arange(len(x))-bar_width, height=WithMigration, label='EWO', alpha=0.8, width=bar_width)
plt.bar(x=np.arange(len(x)), height=OnlyClustering, label='EWC', alpha=0.8, width=bar_width)
plt.bar(x=np.arange(len(x))+bar_width, height=ECWM, label='ECWO', alpha=0.8, width=bar_width)
plt.bar(x=np.arange(len(x))+bar_width*2, height=ENCA_MAMM, label='ENCA+OAMM', alpha=0.8, width=bar_width)
plt.title("各算法可完成任务数结果图")
plt.xlabel("节点最大任务数")
plt.ylabel("可完成任务数")
plt.xticks(range(len(x)), labels=['100', '300', '500', '800', '1000'])
plt.legend(loc="upper left")
plt.savefig("taskNumber.png")
plt.show()