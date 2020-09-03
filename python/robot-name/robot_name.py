import random
from string import ascii_uppercase

ALPHABET_SIZE = len(ascii_uppercase)

class Robot:
    name_used = set()

    def __init__(self):
        self.reset()

    def reset(self):
        name = self.__random_name()
        while name in Robot.name_used:
            name = self.__random_name()
        Robot.name_used.add(name)
        self.name = name

    def __random_name(self):
        return self.__random_letters(2) + self.__random_digits(3)

    def __random_letters(self, n):
        return ''.join([ ascii_uppercase[random.randint(0,ALPHABET_SIZE-1)] for _ in range(n) ])

    def __random_digits(self, n):
        return ''.join([ str(random.randint(0,9)) for _ in range(n) ])