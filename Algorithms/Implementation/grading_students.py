#!/bin/python3

n = int(input().strip())

# Constraints N
if n >= 1 and n <= 60:
    for a0 in range(n):
        grade = int(input().strip())
        # Constraints GRADE
        if grade >=0 and grade <= 100:
            next_multiple = [x for x in filter(lambda w: w%5==0,range(grade,101))][0]
            if(grade >= 38):
                if(next_multiple-grade < 3):
                    grade = next_multiple
        print(grade)
