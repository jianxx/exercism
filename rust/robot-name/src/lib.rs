use lazy_static::lazy_static;
use rand::Rng;
use std::collections::HashSet;
use std::sync::Mutex;

lazy_static! {
    static ref NAME_USED: Mutex<HashSet<String>> = {
        let set = HashSet::new();
        Mutex::new(set)
    };
}
pub struct Robot {
    pub name: String,
}

fn generate_name() -> String {
    let mut rng = rand::thread_rng();
    let letter: char = rng.gen_range(b'A', b'Z' + 1) as char;
    let digits: u16 = rng.gen_range(0, 1000);
    let mut name = format!("{}{}{:03}", letter, letter, digits);
    let mut set = NAME_USED.lock().unwrap();
    while !set.insert(name.clone()) {
        name = format!("{}{}{:03}", letter, letter, digits);
    }
    name
}

impl Robot {
    pub fn new() -> Self {
        Robot {
            name: generate_name(),
        }
    }

    pub fn name(&self) -> &str {
        &self.name
    }

    pub fn reset_name(&mut self) {
        self.name = generate_name()
    }
}
