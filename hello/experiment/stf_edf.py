import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

stf = [4961, 12934, 21433, 30081]
edf = [5459, 15072, 24703, 32603]
total = [5521, 17319, 32554, 59174]

dic = {
    'Max Number': [100, 300, 500, 1000, 100, 300, 500, 1000],
    'Processing Number': [4961, 12934, 21433, 30081, 5459, 15072, 24703, 32603],
    'Algorithm': ['SJF'] * 4 + ['EDF'] * 4,
    'Total Number': [5521, 17319, 32554, 59174] * 2
}

df = pd.DataFrame(dic)
f, ax1 = plt.subplots()
sns.barplot(x='Max Number', y='Processing Number', hue='Algorithm', data=df)
ax2 = ax1.twinx()
sns.lineplot(x=df['Max Number'][:4], y=df['Total Number'][:4])
plt.show()