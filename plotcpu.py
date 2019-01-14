#!/usr/bin/env python3

import pandas as pd
from datetime import datetime
import csv
import matplotlib.pyplot as plt
import matplotlib.dates as mdates

#headers = ['Sensor Value','Date','Time']
#df = pd.read_csv('sre.cpu.data',names=headers)
df = pd.read_csv('sre.cpu.data')
print (df)

#df['Date'] = df['Date'].map(lambda x: datetime.strptime(str(x), '%Y/%m/%d %H:%M:%S.%f'))
x = df['ts']
y = df['cpuavg']
z = df['cpumax']

# plot
plt.figure(figsize=(20, 10), dpi=144)
plt.plot(x,y, alpha=0.1, color='blue',   label="avg", linewidth=0.4, linestyle="-")
plt.plot(x,z, alpha=1,   color='orange', label="max", linewidth=0.4, linestyle="-")
#plt.bar(x,y)
# beautify the x-labels
#plt.gcf().autofmt_xdate()
plt.xlabel('epoch')
plt.ylabel('cpu %')
plt.show()
