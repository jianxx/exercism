def response(hey_bob):
    remark = hey_bob.strip()
    is_question = remark.endswith("?")
    is_yell = remark.isupper()
    if remark == "":
        return "Fine. Be that way!"
    if is_question and is_yell:
        return "Calm down, I know what I'm doing!"
    elif is_question and not is_yell:
        return "Sure."
    elif not is_question and is_yell:
        return "Whoa, chill out!"
    return "Whatever."
