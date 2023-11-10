enum List {
    Cons(i32, Box<List>),
    Nil,
}

enum List2 {
    Cons2(i32, Rc<List2>),
    Nil2,
}

use crate::List::{Cons, Nil};
use crate::List2::{Cons2, Nil2};
use std::rc::Rc;

fn main() {
    let a = Cons(
        5, Box::new(
            Cons(
                10, Box::new(Nil
                    )
                )
            )
        );
    
    // error
    // let b = Cons(3, Box::new(a));
    // let c = Cons(4, Box::new(a));

    // ok
    let a = Rc::new(
        Cons2(
            5, Rc::new(
                Cons2(
                    10, Rc::new(Nil2)
                    )
                )
            )
        );

    println!("count after creating a = {}", Rc::strong_count(&a));
    let b = Cons2(3, Rc::clone(&a));
    // = 
    // let b = Cons2(3, a.clone());
    println!("count after creating b = {}", Rc::strong_count(&a));

    {
        let c = Cons2(4, Rc::clone(&a));
        // = 
        // let c = Cons2(4, a.clone());
        println!("count after creating c = {}", Rc::strong_count(&a));
    }

    println!("count after c goes out of scope = {}", Rc::strong_count(&a));
}
