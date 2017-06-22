#!/bin/python

import sys


a,b,c,d,e = input().strip().split(' ')
a,b,c,d,e = [int(a),int(b),int(c),int(d),int(e)]
arr = [a,b,c,d,e]
result = []
for num in arr:
    if num >= 1 and num <= 10**9:
        result.append(sum(arr)-num)


print(min(result), max(result))