#!/usr/bin/env python3
count = 0

def parser():
    with open("/Users/jblau407/Code/adventofcode/day3/input.txt", 'r') as f:
        file = f.readlines()
    int_list = [int(i) for i in file]
    binary_list = []
    for i in range(len(int_list)):
        binary_list.append(f'{int_list[i]:012d}') # <-- f string conversion needs to be changed for binary num length
    print(f'test {int_list[0:5]}')
    return binary_list

def gamma(binary_list, count):
    zeros, ones = [], []
    for i in range(len(binary_list)):   
        if binary_list[i][count] == '0':
            zeros.append(binary_list[i])
        else:
            ones.append(binary_list[i])
    count += 1
    if count < len(binary_list[0]):
        if len(zeros) > len(ones):
            print(zeros)
            return gamma(zeros, count)
        else:
            print(ones)
            return gamma(ones, count)
    
    elif len(ones)> len(zeros):
        return ones[0]
    else: 
        return zeros[0]
    
def gamma_epsilon(binary):
    eps = ''
    for i in range(len(binary)):
        if int(binary[i]) == 0:
            eps = eps + '1'
        else:
            eps = eps + '0'
    return binary, eps

def binary2int(binary):
    int_val, i= 0, 0
    while(binary != 0): 
        a = binary % 10
        int_val = int_val + a * pow(2, i) 
        binary = binary//10
        i += 1
    print(int_val)
    return int_val

def main():
    x = gamma(parser(),count)
    print(x)
    g, e = gamma_epsilon(x)
    print(g,e)
    print(binary2int(g)*binary2int(e)) 

if __name__ == "__main__":
    main()