import operator
import functools
import itertools
from typing import Iterable
from advent import readlines


def solve(lines):
    lines = [f'.{line}.' for line in lines]
    numbers = []
    is_number = False
    prev = '.' * len(lines[0])
    for ln, line in enumerate(lines):
        next = lines[ln + 1] if ln < (len(lines) - 1) else '.' * len(lines[0])

        is_number = False
        is_serial = False
        current_number = []
        for x, ch in enumerate(line):
            if x == 0:  # the bounds I added
                continue

            if ch.isdigit():
                is_number = True
                current_number.append(ch)
                if any(not k.isdigit() and not k=='.' for k in (
                    line[x-1] + line[x+1] + prev[x-1:x+2] + next[x-1:x+2])):
                    is_serial = True
            else:
                if is_number and is_serial:
                    numbers.append(int(''.join(current_number)))
                current_number.clear()
                is_number = False
                is_serial = False
        prev = line
    return sum(numbers)

lines = readlines(r'./input/03-test.txt')
print(f'Part 1 test: {solve(lines)}')

lines = readlines(r'./input/03-prod.txt')
print(f'Part 1 test: {solve(lines)}')



def solve2(lines):
    lines = [f'.{line}.' for line in lines]
    numbers = []
    is_number = False
    prev = '.' * len(lines[0])
    for ln, line in enumerate(lines):
        next = lines[ln + 1] if ln < (len(lines) - 1) else '.' * len(lines[0])

        is_number = False
        is_serial = False
        current_number = []
        for x, ch in enumerate(line):
            if x == 0:  # the bounds I added
                continue

            if ch.isdigit():
                is_number = True
                current_number.append(ch)
                if any(not k.isdigit() and not k=='.' for k in (
                    line[x-1] + line[x+1] + prev[x-1:x+2] + next[x-1:x+2])):
                    is_serial = True
            else:
                if is_number and is_serial:
                    numbers.append((int(''.join(current_number)), ln, x - len(current_number), x))
                current_number.clear()
                is_number = False
                is_serial = False
        prev = line
    
    # find stars and check if numbers are around (number, line, start, end)


lines = readlines(r'./input/03-test.txt')
print(f'Part 2 test: {solve2(lines)}')
