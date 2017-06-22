#!/bin/python3

import sys


n = int(input().strip())

count = 0

for i in range(n):
    count = count + 1
    hash_count = "#" * count
    difference = n - count
    space = " " * difference
    output = space + hash_count
    print(output)
