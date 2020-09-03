from __future__ import division


class Rational:
    def __init__(self, numer, denom):
        self.numer = numer
        self.denom = denom
        self.__normalize__()

    def __eq__(self, other):
        return self.numer == other.numer and self.denom == other.denom

    def __repr__(self):
        return '{}/{}'.format(self.numer, self.denom)

    def __add__(self, other):
        return Rational(self.numer * other.denom + self.denom * other.numer, self.denom * other.denom)

    def __sub__(self, other):
        return Rational(self.numer * other.denom - self.denom * other.numer, self.denom * other.denom)

    def __mul__(self, other):
        return Rational(self.numer * other.numer, self.denom * other.denom)

    def __truediv__(self, other):
        return Rational(self.numer * other.denom, self.denom * other.numer)

    def __abs__(self):
        return Rational(abs(self.numer), abs(self.denom))

    def __pow__(self, power):
        if str(power).isdigit():
            return Rational(pow(self.numer, abs(power)), pow(self.denom, abs(power)))
        else:
            return Rational(pow(self.numer, power), pow(self.denom, power))

    def __rpow__(self, base):
        return pow(pow(base, self.numer), 1/self.denom)

    def __normalize__(self):
        if self.numer==0:
            self.denom = 1
            pass

        if abs(self.numer) > abs(self.denom):
            smaller = abs(self.denom)
        else:
            smaller = abs(self.numer)
        
        hcf = 1
        for i in range(1,int(smaller) + 1):
            if((self.numer % i == 0) and (self.denom % i == 0)):
                hcf = i
        
        self.numer /= hcf
        self.denom /= hcf

        if self.numer<0 and self.denom<0:
            self.numer = abs(self.numer)
            self.denom = abs(self.denom)
        elif self.numer>0 and self.denom<0:
            self.numer *= -1
            self.denom *= -1
