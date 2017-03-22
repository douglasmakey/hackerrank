#!/bin/python3

import sys


s = input().strip()
count = 1

for letter in s:
    if ord(letter) >= 65 and ord(letter) <= 90:
        count += 1

print(count)