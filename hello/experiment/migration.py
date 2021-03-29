# -*- coding:utf-8 -*-
import matplotlib
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
matplotlib.rcParams['font.sans-serif'] = ['SimHei']

rst_v2 = pd.read_csv('../../experiment_edf_v2_800.csv')
rst_no_migration = pd.read_csv('../../experiment_EdgeOnlyNoMigration.csv')
rst_migration = pd.read_csv('../../experiment_EdgeOnlyWithMigration.csv')
rst_cluster = pd.read_csv('../../experiment_EdgeOnlyWithClusterings.csv')
rst_cloud = pd.read_csv('../../experiment_EdgeCloudWithMigration.csv')

x = [100, 300, 500, 800, 1000]
NoMigration = rst_no_migration['Migration Number']
WithMigration = rst_migration['Migration Number']
OnlyClustering = rst_cluster['Migration Number']
ECWM = rst_cloud['Migration Number']
ENCA_MAMM = rst_v2['Migration Number']

plt.plot(np.arange(len(x)), NoMigration, label='EWNO', marker="s")
plt.plot(np.arange(len(x)), WithMigration, label='EWO', marker="+")
plt.plot(np.arange(len(x)), OnlyClustering, label='EWC', marker='o')
plt.plot(np.arange(len(x)), ECWM, label='ECWO', alpha=0.8, marker='x')
plt.plot(np.arange(len(x)), ENCA_MAMM, label='ENCA+OAMM', marker='h')
plt.title("各算法迁移任务总数结果图")
plt.xlabel("节点最大任务数")
plt.ylabel("迁移任务总数")
plt.xticks(range(len(x)), labels=['100', '300', '500','800', '1000'])
plt.legend(loc="upper left")
plt.savefig("migration.png")
plt.show()