use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("Here's vector: {:?}", v);
    });

    // error: value used here after move
    // drop(v);

    handle.join().unwrap();

    // error: value used here after move
    // drop(v);
}
