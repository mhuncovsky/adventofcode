import operator
import functools
import itertools
from typing import Iterable
from advent import readlines

counts = {
    'red': 12,
    'green': 13,
    'blue': 14,
}

def solve(lines):
    return sum(int(game) for game, pulls in (
            (game.strip().split()[-1],
            (pull.split(', ') for pull in pulls.strip().split('; '))
            ) for game, pulls in (line.split(':', 1) for line in lines)
        ) if all(
            all(counts[cube] >= int(scount)
                for scount, cube in (s.split(' ', 1) for s in pull)
            ) for pull in pulls
        ))

result = solve(readlines(r'./input/02-test.txt'))
print(f'Part 1 test: {result}')

result = solve(readlines(r'./input/02-prod.txt'))
print(f'Part 1 prod: {result}')



def solve2(lines):
    games = [
        (game.strip().split()[-1],
        [pull.split(', ') for pull in pulls.strip().split('; ')]
        ) for game, pulls in (line.split(':', 1) for line in lines)
    ]
    return sum(minimum(pulls) for _, pulls in games)

def minimum(pulls: list[list[str]]):
    counts = {}
    for count, color in (cubes.split(' ', 1) for pull in pulls for cubes in pull):
        count = int(count)
        if color not in counts or count > counts[color]:
            counts[color] = count
    return functools.reduce(operator.mul, counts.values())



result = solve2(readlines(r'./input/02-test.txt'))
print(f'Part 2 test: {result}')

result = solve2(readlines(r'./input/02-prod.txt'))
print(f'Part 2 prod: {result}')
