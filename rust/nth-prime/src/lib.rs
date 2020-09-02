pub fn nth(n: u32) -> u32 {
    (2..).filter(|x| is_prime(*x)).nth(n as usize).unwrap()
}

fn is_prime(n: u32) -> bool {
    if n<2 {
        return false;
    }else if n==2 {
        return true;
    }else if n%2==0 {
        return false;
    }
    let x = n/2;
    for i in 1..x {
        if n%(2*i+1)==0 {
            return false;
        }
    }
    return true;
}