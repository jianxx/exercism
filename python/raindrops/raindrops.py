def convert(number):
    if(number%3!=0 and number%5!=0 and number%7!=0):
        return str(number)
    else:
        r = ''
        if number%3==0:
            r += 'Pling'
        if number%5==0:
            r += 'Plang'
        if number%7==0:
            r += 'Plong'

        return r
