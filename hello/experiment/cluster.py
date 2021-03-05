import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

clusterings = pd.read_csv('../../clusterings.csv')
centers = pd.read_csv('../../centers.csv')
clusterings['X'] = pd.to_numeric(clusterings['X'], errors='coerce')
clusterings['Clusterings'] = clusterings['Clusterings'].astype(int)

fig, ax = plt.subplots()
sns.scatterplot(data=clusterings, x="X", y="Y", hue="Clusterings")
sns.scatterplot(x=centers['X'], y=centers['Y'])

r = 25 / 2
for row in centers.itertuples():
    a = row.X
    b = row.Y
    theta = np.arange(0, 2*np.pi, 0.01)
    x = a + r * np.cos(theta)
    y = b + r * np.sin(theta)
    plt.plot(x, y)
ax.set_xlim(0, 110)
ax.set_ylim(0, 110)
plt.show()