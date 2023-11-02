use std::fmt::Display;

fn longest_with_an_announcement<'a, T>(x: &'a str, y: &'a str, ann: T) -> &'a str
    where T: Display
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let x = "this is x";
    let y = "this is y";
    let t = format!("this is t, x is {}, y is {}", x, y);
    let result = longest_with_an_announcement(x, y, t);
    println!("result: {}", result);
}
