pub fn brackets_are_balanced(string: &str) -> bool {
    let mut l: Vec<char> = Vec::new();
    for c in string.chars() {
        match c {
            '{' | '[' | '(' => l.push(c),
            '}' => {
                if l.pop() != Some('{') {
                    return false;
                }
            }
            ']' => {
                if l.pop() != Some('[') {
                    return false;
                }
            }
            ')' => {
                if l.pop() != Some('(') {
                    return false;
                }
            }
            _ => {}
        }
    }

    l.is_empty()
}
