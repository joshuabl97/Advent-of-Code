#!/usr/bin/env python3
import sys

sys.path.insert(0, '/Users/jblau407/Code/')
from adventofcode import problemInput

day = 2

def formatter(input):
    x = input.split()
    for i in range(len(x)): 
        if i % 2 != 0:  
            x[i] = int(x[i])
    return x

def solve1(x): 
    horizontal = 0 
    depth = 0
    for i in range(len(x)-1): 
        if i % 2 == 0: 
            if x[i] == 'forward':
                horizontal += x[i+1]
            elif x[i] == 'down': 
                depth += x[i+1]
            elif x[i] == 'up':
                depth -= x[i+1]
    return horizontal*depth

def solve2(x):
    horizontal = 0 
    depth = 0
    aim = 0
    for i in range(len(x)-1): 
        if i % 2 == 0: 
            if x[i] == 'forward':
                horizontal += x[i+1]
                depth += (aim*x[i+1])
            elif x[i] == 'down': 
                aim += x[i+1]
            elif x[i] == 'up':
                aim -= x[i+1]
    return horizontal*depth    

def main(): 
    input = problemInput.get_input(day)
    x = formatter(input)
    print(solve1(x))
    print(solve2(x))

if __name__ == "__main__":
    main()