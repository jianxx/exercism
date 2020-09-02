pub fn verse(n: u32) -> String { 
    if n==0 {
        let s:String = format!("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n");
        return s;
    }else if n==1 {
        let s:String = format!("{} bottle of beer on the wall, {} bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n, n);
        return s;
    }else if n==2 {
        let s:String = format!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottle of beer on the wall.\n", n, n, n-1);
        return s;
    }else {
        let s:String = format!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottles of beer on the wall.\n", n, n, n-1);
        return s;
    }
}

pub fn sing(start: u32, end: u32) -> String {
    let mut s = String::new();
    for i in (end..start+1).rev() {
        let m:String = verse(i);
        s.push_str(&m);
        if i>end {
            s.push_str("\n");
        }
    }
    return s;
}
