import random

guess = random.randint(1, 101)

i = 1
while True:
    print('num game: %d times, int num please' % i)
    try:
        t = int(input())
        i += 1
    except ValueError:
        print('wrong input')
        continue
    if t == guess:
        print('right!', guess)
        break
    elif t > guess:
        print('too big')
    else:
        print('too small')
