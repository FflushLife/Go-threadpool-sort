#/usr/bin/python3

import random

size = random.randint(1, 1000000)
sample = [str(x) for x in random.sample(list(range(1, size)), size-1)]
with open("rand_data.txt", "w+") as out:
    out.write(','.join(sample))
    out.write('\n')
