#!/usr/bin/env python3
import sys

sys.path.insert(0, '/Users/jblau407/Code/')
from adventofcode import problemInput

day = 1

def convert(input): 
    x = input.split()
    for i in range(0, len(x)):
        x[i] = int(x[i])
    
    return x

def solve1(x): 
    count = 0
    for i in range(len(x)-1): 
        if x[i] < x[i+1]: 
            count += 1
    return count

def solve2(x): 
    count = 0
    for i in range(len(x)-3):  
        a = x[i] + x[i+1] + x[i+2]
        b = x[i+1] + x[i+2] + x[i+3] 
        if a < b: 
            count += 1
    return count

def main():
    input = problemInput.get_input(day)
    x = convert(input)
    count1 = solve1(x)    
    print(count1)
    count2 = solve2(x)
    print(count2)

if __name__ == "__main__":
    main()