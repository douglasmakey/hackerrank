#!/bin/python3
import sys

s,t = input().strip().split(' ')
s,t = [int(s),int(t)]
a,b = input().strip().split(' ')
a,b = [int(a),int(b)]
m,n = input().strip().split(' ')
m,n = [int(m),int(n)]
apple = [int(apple_temp) for apple_temp in input().strip().split(' ')]
orange = [int(orange_temp) for orange_temp in input().strip().split(' ')]
apple_pos =  [a + app for app in apple]
orange_pos = [b + org for org in orange]

def is_on_house(pos):
    return s <= pos <= t

print(len(list(filter(is_on_house, apple_pos))))
print(len(list(filter(is_on_house, orange_pos))))
