#!/bin/python3

n,k = input().strip().split(' ')
n,k = [int(n),int(k)]
height = sorted(list(map(int, input().strip().split(' '))))

max_height = height[-1]
max_jump = k

if max_jump >= max_height:
    print(0)
else: print(max_height - max_jump)
