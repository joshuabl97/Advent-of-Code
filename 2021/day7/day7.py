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
    return round(fuel_consumption)

def brute_force(nums):
    nums.sort() 
    num_fuel_list = []

    for i in range(len(nums)):
        local_fuel_list = []
        for j in range(len(nums)):
            local_fuel_list.append(abs(nums[i]-nums[j])) 
        num_fuel_list.append(local_fuel_list) 

    avg_distances = [] 
    for i in range(len(num_fuel_list)):
        fuel_avg = 0
        for j in range(len(num_fuel_list[i])):
            fuel_avg += num_fuel_list[j][i]
        avg_distances.append(fuel_avg)

    avg_distances.sort()
    # this method is more accurate than the method above (brute force checker)
    return avg_distances[0]

def p2_estimate(nums):
    nums_avg = statistics.mean(nums)

    fuel_consumption = 0
    for i in range(len(nums)):
        diff = round(abs(nums[i]-nums_avg))
        cost = 0
        for i in range(diff):
            cost += i+1
        fuel_consumption += cost

    return fuel_consumption

def problem2(nums):
    nums.sort()
    

if __name__ == "__main__":
    print(f"estimated shortest path: {problem1(parser())}")
    shortest_distance = brute_force(parser())
    print(f"shortest path: {shortest_distance}")
    print(f"part two estimate: {p2_estimate(parser())}")
    print(f"part two: {problem2(parser())}")