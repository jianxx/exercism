pub fn raindrops(n: u32) -> String {
    if n%3!=0 && n%5!=0 && n%7!=0 {
        let s: String = n.to_string();
        return s;
    }else {
        let mut s = String::new();
        if n%3==0 {
            s.push_str("Pling")
        }
        if n%5==0 {
            s.push_str("Plang")
        }
        if n%7==0 {
            s.push_str("Plong")
        }
        return s;
    }
}
