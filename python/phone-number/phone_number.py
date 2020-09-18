import re


class PhoneNumber:
    def __init__(self, number):
        r = re.search(
            "^\\+?1?\\s*\\(?([2-9]{1}[0-9]{2})\\)?[-.]?\\s*([2-9]{1}[0-9]{2})[-.]?\\s*([0-9]{4})\\s*$",
            number,
        )
        if not r:
            raise ValueError(number)
        else:
            self.number = r.group(1) + r.group(2) + r.group(3)
            self.area_code = r.group(1)
            self.exchange_code = r.group(2)
            self.phone_code = r.group(3)
        pass

    def pretty(self):
        return "(%s)-%s-%s" % (self.area_code, self.exchange_code, self.phone_code)
