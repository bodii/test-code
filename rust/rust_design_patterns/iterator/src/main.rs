mod users;
use crate::users::UserCollection;

fn main() {
    print!("Iterators are widely used in the standard library: ");

    let array = &[1, 2, 3];
    let iterator = array.iter();

    iterator.for_each(|e| print!("{} ", e));

    println!("\n\nLet's test our own iterator.\n");

    let users = UserCollection::new();
    let mut iterator = users.iter();

    println!("1nd element: {:?}", iterator.next());
    println!("1nd element: {:?}", iterator.next());
    println!("1nd element: {:?}", iterator.next());
    println!("1nd element: {:?}", iterator.next());

    print!("\nAll elements in user collection: ");
    users.iter().for_each(|e| print!("{} ", e));

    println!();
}
