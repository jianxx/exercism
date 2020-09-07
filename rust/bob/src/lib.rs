pub fn reply(message: &str) -> &str {
    let m = message.trim();
    if m.is_empty() {
        return "Fine. Be that way!";
    }
    let is_question = m.ends_with('?');
    let is_yell = m.contains(char::is_alphabetic) && !m.contains(char::is_lowercase);
    match(is_question, is_yell) {
        (true, true) => "Calm down, I know what I'm doing!",
        (true, false) => "Sure.",
        (false, true) => "Whoa, chill out!",
        (false, false) => "Whatever.",
    }
}
