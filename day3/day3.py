#!/usr/bin/env python3
import sys

sys.path.insert(0, '/Users/jblau407/Code/')
from adventofcode import problemInput

day = 3
oxygenCounter = 0
c02counter = 0

def formatter(input):
    x = input.split()
    for i in range(len(x)): 
        x[i] = int(x[i])
    return x

def problem1(x):
    a = [0] * (len(str(x[0]))+1) 
    for i in range(len(x)): 
        num_str = f'{x[i]:012d}' ### TIL int to string conversion removes leading zeros... ###
        for i in range(len(num_str)):
            num_int = int(num_str[i])
            if num_int == 1: 
                a[i] += 1
            else: 
                a[i] -= 1
    return a

def gamma_epsilon(a):
    gamma = []
    epsilon = []
    for i in range(len(a)): 
        if a[i] > 0:
            epsilon.append(0) 
            gamma.append(1)
        else: 
            epsilon.append(1)
            gamma.append(0)

    return gamma, epsilon

def join_list(list): 
    s = [str(i) for i in list]
    res = int("".join(s))
    return res

def binary2int(binary):
    int_val, i= 0, 0
    while(binary != 0): 
        a = binary % 10
        int_val = int_val + a * pow(2, i) 
        binary = binary//10
        i += 1
    return int_val

def problem2_format(x):
    a = []
    for i in range(len(x)): 
        num_str = f'{x[i]:012d}'
        a.append(num_str)

    return a

def oxygenGen(a, oxygenCounter): 
    if len(a) == 1: 
        print('test')
        return a[0]

    zeros, ones = 0, 0
    zero_list, one_list = [], [] 

    for i in range(len(a)):
        correct_num = a[i]
        num_int = int(correct_num[oxygenCounter])
        if num_int == 1:
            ones += 1
            one_list.append(a[i]) 
        else:
            zeros += 1
            zero_list.append(a[i])

    if zeros > ones:
        oxygenCounter += 1
        print(zero_list)
        oxygenGen(zero_list, oxygenCounter)
    else: 
        oxygenCounter += 1
        print(one_list)
        oxygenGen(one_list, oxygenCounter)

def c02Gen(a, c02counter): #### this func is very incomplete/a copy/paste of the above func
    if (len(a)) == 1:
        print('test')
        return a 

    zeros, ones = 0, 0
    zero_list, one_list = [], [] 

    for i in range(len(a)):
        correct_num = a[i]
        num_int = int(correct_num[c02counter])
        if num_int == 1:
            ones += 1
            one_list.append(a[i]) 
        else:
            zeros += 1
            zero_list.append(a[i])

   
    if zeros < ones:
        c02counter += 1
        print(zero_list)
        oxygenGen(zero_list, c02counter)
    else: 
        c02counter += 1
        print(one_list)
        oxygenGen(one_list, c02counter)


def main(): 
    input = (problemInput.get_input(day))
    x = (formatter(input))
    y,z = gamma_epsilon(problem1(x))
    a,b = binary2int(join_list(y)), binary2int(join_list(z))
    print(a*b) 
    a = (problem2_format(x))
    c02 = oxygenGen(a, oxygenCounter), c02Gen(a, c02counter)
    print(c02)

if __name__ == "__main__":
    main()