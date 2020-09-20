use std::iter::FromIterator;

#[derive(Debug)]
struct Node<T> {
    pub data: T,
    pub next: Option<Box<Node<T>>>,
}

impl<T> Node<T> {
    pub fn new(e: T, next: Option<Box<Node<T>>>) -> Self {
        Node {
            data: e,
            next: next,
        }
    }
}

pub struct SimpleLinkedList<T> {
    // Delete this field
    // dummy is needed to avoid unused parameter error during compilation
    // dummy: ::std::marker::PhantomData<T>,
    head: Option<Box<Node<T>>>,
    length: usize,
}

impl<T> SimpleLinkedList<T> {
    pub fn new() -> Self {
        SimpleLinkedList {
            head: None,
            length: 0,
        }
    }

    pub fn len(&self) -> usize {
        self.length
    }

    pub fn push(&mut self, _element: T) {
        let node = Box::new(Node::new(_element, self.head.take()));
        self.head = Some(node);
        self.length += 1;
    }

    pub fn pop(&mut self) -> Option<T> {
        match self.length {
            0 => None,
            _ => {
                self.length -= 1;
                self.head.take().map(|x| {
                    let node = *x;
                    self.head = node.next;
                    node.data
                })
            }
        }
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_ref().map(|x| &x.data)
    }
}

impl<T: Clone> SimpleLinkedList<T> {
    pub fn rev(self) -> SimpleLinkedList<T> {
        let mut rev_list: SimpleLinkedList<T> = SimpleLinkedList::new();
        self.head.as_ref().map(|node| {
            rev_list.push(node.data.clone());
        });
        rev_list
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I: IntoIterator<Item = T>>(_iter: I) -> Self {
        unimplemented!()
    }
}

// In general, it would be preferable to implement IntoIterator for SimpleLinkedList<T>
// instead of implementing an explicit conversion to a vector. This is because, together,
// FromIterator and IntoIterator enable conversion between arbitrary collections.
// Given that implementation, converting to a vector is trivial:
//
// let vec: Vec<_> = simple_linked_list.into_iter().collect();
//
// The reason this exercise's API includes an explicit conversion to Vec<T> instead
// of IntoIterator is that implementing that interface is fairly complicated, and
// demands more of the student than we expect at this point in the track.

impl<T> Into<Vec<T>> for SimpleLinkedList<T> {
    fn into(self) -> Vec<T> {
        unimplemented!()
    }
}
