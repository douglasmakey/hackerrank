#!/bin/python3

n = int(input().strip())

# Constraints N
if 1 <= n <= 10**5:
    height = sorted([int(height_temp) for height_temp in input().strip().split(' ')])

major = height[-1]
minnor = 0
arr = []

for candle in height:
    # Constraints Height
    if 1 <= candle <= 10**7:
        if candle >= major and candle >= minnor:
            major = candle
            arr.append(major)
        else:
            minnor = candle

print(len(arr))