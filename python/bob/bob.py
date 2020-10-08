import re


def response(hey_bob):
    remark = "".join(str.split(hey_bob))
    if len(remark) == 0:
        return "Fine. Be that way!"
    is_question = remark.endswith("?")
    is_yell = re.search("[a-zA-Z]", remark) != None and remark.upper() == remark
    if is_question and is_yell:
        return "Calm down, I know what I'm doing!"
    elif is_question and not is_yell:
        return "Sure."
    elif not is_question and is_yell:
        return "Whoa, chill out!"
    return "Whatever."
