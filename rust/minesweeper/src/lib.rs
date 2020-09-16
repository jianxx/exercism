pub fn annotate(minefield: &[&str]) -> Vec<String> {
    minefield
        .iter()
        .enumerate()
        .map(|(i, row)| {
            (0..row.len())
                .map(|j| {
                    if minefield[i].chars().nth(j) == Some('*') {
                        '*'
                    } else {
                        let mines = (i.saturating_sub(1)..(i + 1))
                            .flat_map(|x| (j.saturating_sub(1)..(j + 1)).map(move |y| (x, y)))
                            .filter(|&(x, y)| (x, y) != (i, j))
                            .filter_map(|(x, y)| minefield.get(x).and_then(|r| r.chars().nth(y)))
                            .filter(|&c| c == '*')
                            .count();
                        match mines {
                            0 => ' ',
                            1 => '1',
                            2 => '2',
                            3 => '3',
                            4 => '4',
                            5 => '5',
                            6 => '6',
                            7 => '7',
                            8 => '8',
                            9 => '9',
                            _ => unreachable!(),
                        }
                    }
                })
                .collect()
        })
        .collect()
}
