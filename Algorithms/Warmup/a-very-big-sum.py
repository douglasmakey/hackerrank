#!/bin/python3

n = int(input().strip())
arr = [int(arr_temp) for arr_temp in input().strip().split(' ')]

sum_total = 0
if 0 <= n <= 10:
    for num in arr:
        if 0 <= num <= 10**10:
            sum_total += num

print(sum_total)
