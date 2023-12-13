use std::io::{BufReader, Cursor, Read};

fn main() {
    let mut buf = [0u8; 10];

    let mut input = BufReader::new(Cursor::new("Input data"));
    input.read(&mut buf).ok();

    println!("Read from a buffered reader: ");

    for byte in buf {
        print!("{}", char::from(byte));
    }

    println!();
}
