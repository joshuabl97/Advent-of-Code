#!/usr/bin/env python3
counter = 0

def parser(): 
    with open("/Users/jblau407/Code/adventofcode/day6/input.txt", 'r') as f:
        file = f.readlines()
        fileList = list(file[0].split(','))
        lanternFish = []
        for i in range(len(fileList)):
            lanternFish.append(int(fileList[i]))
        return lanternFish

def part1(x, counter):
    if counter == 80:
        print(len(x))
    else: 
        for i in range(len(x)):
            if x[i] == 0:
                x[i] = 6
                x.append(8)
            else:
                x[i] = x[i]-1
        counter += 1
        part1(x, counter)   

def parser2(lanternFish):
    fishDict = {0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0}
    for i in range(8):
        for j in range(len(lanternFish)):
            if lanternFish[j] == i:
                fishDict[i] += 1
    counter == 0
    return fishDict

def part2(x, counter):
    if counter < 256: 
        counter += 1
        renew = 0
        for i in range(len(x)):
            if i == 0:
                renew += x[i]
                x[i] = 0
            else:
                x[i-1] = x[i]
        x[6] += renew
        x[8] = renew
        part2(x, counter)
    else:
        fish_sum = 0
        for i in range(len(x)):
            fish_sum += x[i]
        print(fish_sum)
    
def main():
    part1(parser(), counter)
    part2(parser2(parser()), counter)

if __name__ == "__main__":
    main()