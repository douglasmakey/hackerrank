#!/bin/python3

n = int(input().strip())
arr = [int(arr_temp) for arr_temp in input().strip().split(' ')]

positive = [ x for x in arr if x >0]
negative = [x for x in arr if x < 0]
zero = [ x for x in arr if x ==0]
print('%.6f' %(len(positive)/n))
print('%.6f' %(len(negative)/n))
print('%.6f' %(len(zero)/n))