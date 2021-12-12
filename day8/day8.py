#!/usr/bin/env python3
from os import sep


def parser():
    with open('/Users/jblau407/Code/adventofcode/day8/input.txt', 'r') as f:
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

def problem2(before, after):
    all_digits = []
    after_digits = []

    for i in range(len(after)):
        after_list = after[i].split(' ')
        nested_digits = []
        for j in range(len(after_list)):
            nested_digits.append(list(after_list[j]))
        after_digits.append(nested_digits)


    for i in range(len(before)): 
        before_list = (before[i].split(' '))
        before_list.pop()
        nested_digits = []
        for j in range(len(before_list)):
            nested_digits.append(list(before_list[j]))
        all_digits.append(nested_digits)

    output_value = 0
    for i in range(len(all_digits)):
        ten_digits = all_digits[i] 
        sep_digits = after_digits[i]

        dict_num_keys = [0,1,2,3,4,5,6,7,8,9]
        dict_nums = {}
        for i in dict_num_keys:
            dict_nums[i] = None

        dict_value_keys = ['A','B','C','D','E','F','G']
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
        dict_values['A'] = li_dif[0]
        
        #step 3: the number that has 6 digits and does not contain digits from 1 is number 6
        for j in range(len(ten_digits)):
            if len(ten_digits[j]) == 6:
                li1 = ten_digits[j]
                li2 = dict_nums[1]
                bool_list = [x in li1 for x in li2]
                if bool_list[0] == False or bool_list[1] == False:
                    dict_nums[6] = ten_digits[j]
        
        #step 4: the digit in 6 that is in 1 is f, the remaining digit in 1 is c
        a_set = set(dict_nums[1])
        b_set = set(dict_nums[6])
        f_val = list(a_set&b_set)
        dict_values['F'] = f_val[0]
        for i in range(len(dict_nums[1])):
            if dict_nums[1][i] != dict_values['F']:
                dict_values['C'] = dict_nums[1][i]

        #step 5: grab unknown vals in 4 and check to see if they exist in 0,9
        ### if it exists in ten_digits[i] == 6 that value is 9
        ### the remaining value in ten_digits[i] == 6 is 0
        current_values = [i for i in dict_values.values() if i]
        unknown_four =list(set(dict_nums[4]) - set(current_values))          
        
        for j in range(len(ten_digits)):
            if len(ten_digits[j]) == 6 and ten_digits[j] != dict_nums[6]:
                if len((list(set(unknown_four) - set(ten_digits[j])))) == 0:
                    dict_nums[9] = ten_digits[j]
                else:
                    dict_nums[0] = ten_digits[j]

        #step 6: compare 6,9 the different value that isn't present in 1 is e
        ### if solution doesn't work, start over here
        six_nine_diff = list(set(dict_nums[6]) - set(dict_nums[9]))
        dict_values['E'] = six_nine_diff[0]
        
        #step 7: combine values of 6 & 9 into a single dict, the value that isn't present in 0 is d
        six_nine_same = list(set(dict_nums[6] + dict_nums[9]))
        dict_values['D'] = list(set(six_nine_same) - set(dict_nums[0]))[0]

        #step 8: the value that isn't present in the values list from 4 is b && the remaining value is g
        current_values = [i for i in dict_values.values() if i]
        dict_values['B'] = list(set(dict_nums[4]) - set(current_values))[0]
        a = list(map(lambda x: x.lower(), dict_value_keys))
        dict_values['G'] = list(set(a) - set([i for i in dict_values.values() if i]))[0] 

        #step 9: assign the rest of nums using values
        two = list(map(dict_values.get, ['A','C','D','E','G']))
        three = list(map(dict_values.get, ['A','C','D','F','G']))
        five = list(map(dict_values.get, ['A','B','D','F','G']))

        for i in range(len(ten_digits)):
            if len(ten_digits[i]) == 5:
                if len(list(set(ten_digits[i])-set(two))) == 0:
                    dict_nums[2] = ten_digits[i]
                elif len(list(set(ten_digits[i])-set(three))) == 0:
                    dict_nums[3] = ten_digits[i]
                elif len(list(set(ten_digits[i])-set(five))) == 0:
                    dict_nums[5] = ten_digits[i]

        #step 10: solve for after digit values and append int to list
        after_ints = []

        zero = list(map(dict_values.get, ['A','B','C','E','F','G'])) 
        six = list(map(dict_values.get, ['A','B','D','E','F','G']))
        nine = list(map(dict_values.get, ['A','B','C','D','F','G']))        

        for i in range(len(sep_digits)):
            if len(sep_digits[i]) == 7:
                after_ints.append(8)
            elif len(sep_digits[i]) == 2:
                after_ints.append(1)
            elif len(sep_digits[i]) == 4:
                after_ints.append(4)
            elif len(sep_digits[i]) == 3:
                after_ints.append(7)
            elif len(sep_digits[i]) == 5:
                if len(list(set(sep_digits[i]) - set(two))) == 0:
                    after_ints.append(2)
                elif len(list(set(sep_digits[i]) - set(three))) == 0:
                    after_ints.append(3)
                elif len(list(set(sep_digits[i]) - set(five))) == 0:
                    after_ints.append(5)
            elif len(list(set(sep_digits[i]) - set(zero))) == 0:
                after_ints.append(0)
            elif len(list(set(sep_digits[i]) - set(six))) == 0:
                after_ints.append(6)
            elif len(list(set(sep_digits[i]) - set(nine))) == 0:
                after_ints.append(9) 
                
        s = [str(i) for i in after_ints]
        res = int("".join(s))
        output_value += res  

    return output_value
def main():
    print(problem1(parser()[1]))
    before, after = parser()
    print(problem2(before, after))

if __name__ == '__main__':
    main()