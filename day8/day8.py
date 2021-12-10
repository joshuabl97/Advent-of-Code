#!/usr/bin/env python3
def parser():
    with open('/Users/jblau407/Code/adventofcode/day8/smallTest.txt', 'r') as f:
        lines = f.readlines()
        before,after = [],[]
        for line in lines:
            before_parse, after_parse = line.split('|')
            before.append(before_parse)
            after.append(after_parse.strip())
    return before, after

def problem1(after):
    output = []
    one,four,seven,eight = 0,0,0,0
    for i in range(len(after)):
        after_list = (after[i].split(' '))
        for j in range(len(after_list)):
            output.append(after_list[j])
    for i in range(len(output)):
        if len(output[i]) == 7:
            eight += 1
        elif len(output[i]) == 4:
            four += 1
        elif len(output[i]) == 3:
            seven += 1
        elif len(output[i]) == 2:
            one += 1
    return(one+four+seven+eight)

def problem2(before):
    all_digits = []

    for i in range(len(before)): 
        before_list = (before[i].split(' '))
        before_list.pop()
        nested_digits = []
        for j in range(len(before_list)):
            nested_digits.append(list(before_list[j]))
        all_digits.append(nested_digits)
    for i in range(len(all_digits)):
        ten_digits = all_digits[i] 
        dict_num_keys = [0,1,2,3,4,5,6,7,8,9]
        dict_nums = {}
        for i in dict_num_keys:
            dict_nums[i] = None

        dict_value_keys = ['a','b','c','d','e','f','g']
        dict_values = {}
        for i in dict_value_keys:
            dict_values[i] = None
        
        #step 1: filter for unique values and assign to nums 
        for j in range(len(ten_digits)):
            if len(ten_digits[j]) == 2:
                dict_nums[1] = ten_digits[j]
            elif len(ten_digits[j]) == 3:
                dict_nums[7] = ten_digits[j]
            elif len(ten_digits[j]) == 4:
                dict_nums[4] = ten_digits[j]
            elif len(ten_digits[j]) == 7:
                dict_nums[8] = ten_digits[j]

        #step 2: compare nums 1, 7 to get value a
        one_list = (dict_nums[1])
        seven_list = (dict_nums[7])
        li_dif = [i for i in one_list + seven_list if i not in one_list or i not in seven_list]
        dict_values['a'] = li_dif[0]
        
        #step 3: the number that has 6 digits and does not contain digits from 1 is number 6
        for j in range(len(ten_digits)):
            if len(ten_digits[j]) == 6:
                li1 = ten_digits[j]
                li2 = dict_nums[1]
                bool_list = [x in li1 for x in li2]
                if bool_list[0] == False or bool_list[1] == False:
                    dict_nums[6] = ten_digits[j]
        print(dict_nums)

        #step 4: the digit in 6 that is in one is f, the remaining digit in 1 is c    

        
        print('seperator')
def main():
    print(problem1(parser()[1]))
    print(problem2(parser()[0]))

if __name__ == '__main__':
    main()