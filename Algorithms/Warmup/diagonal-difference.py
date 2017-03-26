#!/bin/python3

n = int(input().strip())
a = []
for a_i in range(n):
    a_t = [int(a_temp) for a_temp in input().strip().split(' ')]
    a.append(a_t)

first = sum([row[idx] for idx, row in enumerate(a)])
second = sum([row[n-idx-1] for idx, row in enumerate(a)])

print(str(abs(first - second)))