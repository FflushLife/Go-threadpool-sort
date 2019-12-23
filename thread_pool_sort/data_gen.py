#/usr/bin/python3

import random

r_size = 1000000
size = random.randint(1, 10000)
sample = [str(x) for x in random.sample(list(range(1, size)), size-1)]
with open("rand_data.txt", "a+") as out:
    out.write(','.join(sample))
    out.write('\n')
    # Second writing for checking rhread pull acceleration
    out.write(','.join(sample))
    out.write('\n')
