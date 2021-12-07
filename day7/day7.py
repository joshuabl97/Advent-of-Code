#!/usr/bin/env python3
import statistics

def parser():
    with open("/Users/jblau407/Code/adventofcode/day7/test.txt", 'r') as f:
        lines = f.readlines()
        numbers = [int(line.strip()) for line in lines[0].split(",")]
    return numbers

def problem1(nums):
    nums_med = statistics.median(nums)
    fuel_consumption = 0
    for i in range(len(nums)):
        fuel_consumption += abs(nums[i]-nums_med)
    print(fuel_consumption)

def problem2(nums):
    nums_avg = statistics.mean(nums)
    fuel_consumption = 0
    for i in range(len(nums)):
        fuel_consumption += abs(nums[i]-nums_avg)
    return fuel_consumption

def main():
    print(problem1(parser()))
    print(problem2(parser()))

if __name__ == "__main__":
    main()