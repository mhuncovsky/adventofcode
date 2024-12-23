import string
from advent import readlines


lines = readlines(r'./input/01-input.txt')
lines_as_digits = [[char for char in line if char.isdigit()] for line in lines]
print(sum(int(digits[0] + digits[-1]) for digits in lines_as_digits))


digits = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
}

def replace_first(line: str, words: dict[str, str]):
    found = {}
    for word, _ in words.items():
        if (index := line.find(word)) >= 0:
            if not index in found:
                found[index] = []
            found[index].append(word)
    if not found:
        return line
    keys = sorted(found.keys())
    first = found[keys[0]][0]
    return line.replace(first, words[first], 1)
    



def process(line: str):
    reverse_digits = {k[::-1]:v for k, v in digits.items()}
    first = getfirst(line, digits)
    last = getfirst(line[::-1], reverse_digits)
    # print(f'{line} -> {first}, {last}')
    return first + last

def getfirst(line, table) -> str:
    for i, ch in enumerate(line):
        if ch.isdigit():
            return ch
        else:
            for word, digit in table.items():
                if line[i:].startswith(word):
                    return digit
    return ''

lines = readlines(r'./input/01-input.txt')
print(sum(int(process(line)) for line in lines))

# for i, line in enumerate(lines):
#     if 'eightwo' in line:
#         print(f'{i}, {line} -> {process(line)}')

# 53532 too low