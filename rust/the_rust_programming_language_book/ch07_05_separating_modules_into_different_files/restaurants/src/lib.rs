pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}

mod front_of_hosts;

pub use crate::front_of_hosts::hosting;

pub fn eat_at_restaurant() {
    hosting::add_to_waitlist();
}
