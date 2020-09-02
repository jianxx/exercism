pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let mut r = vec!();
    for (row_num, row) in input.iter().enumerate() {
        for (col_num, v) in row.iter().enumerate() {
            if row.iter().all(|x| x<=v) && input.iter().all(|x| *v<=x[col_num]) {
                r.push((row_num, col_num));
            }
        }
    }
    return r;
}
