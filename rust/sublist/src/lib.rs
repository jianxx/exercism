#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list == _second_list {
        return Comparison::Equal;
    }
    if _first_list.len() == 0 {
        return Comparison::Sublist;
    }
    if _second_list.len() == 0 {
        return Comparison::Superlist;
    }

    if _first_list.len() < _second_list.len() {
        let r = sublist(_second_list, _first_list);
        return match r {
            Comparison::Sublist => Comparison::Superlist,
            Comparison::Superlist => Comparison::Sublist,
            _ => r,
        };
    } else {
        if _first_list
            .windows(_second_list.len())
            .any(|w| w == _second_list)
        {
            return Comparison::Superlist;
        }
    }

    Comparison::Unequal
}
