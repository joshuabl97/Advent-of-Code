#!/usr/bin/env python3
def test(count):
    if count < 3:
        count += 1
        print(f'test {count}')
        return test(count)
    else:
        count = 1000
        return count

print(test(0))