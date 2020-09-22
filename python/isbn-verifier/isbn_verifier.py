import re


def is_valid(isbn):
    r = re.search("^([0-9]{1})-?([0-9]{3})-?([0-9]{5})-?([0-9,X]{1})$", isbn)
    if not r:
        return False
    else:
        return (
            int(r.group(1)) * 10
            + int(r.group(2)[0]) * 9
            + int(r.group(2)[1]) * 8
            + int(r.group(2)[2]) * 7
            + int(r.group(3)[0]) * 6
            + int(r.group(3)[1]) * 5
            + int(r.group(3)[2]) * 4
            + int(r.group(3)[3]) * 3
            + int(r.group(3)[4]) * 2
            + (10 if r.group(4) == "X" else int(r.group(4)))
        ) % 11 == 0
