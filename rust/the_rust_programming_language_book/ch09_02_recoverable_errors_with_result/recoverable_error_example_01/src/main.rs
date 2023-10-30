use std::fs::File;
use std::io::ErrorKind;

fn main() {
    /*
    let f = File::open("hello.txt");

    let f = match f {
        Ok(file) = file,
        Err(error) => match error.kind() {
            ErrorKind::NetFound => match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(e) => panic!("Problem creating the file: {:?}", e),
            },
            other_error => panic!("Problem opening the file: {:?}", other_error),
        },
    };
    */
    // or
    or_open_file();

    // unwrap or expect use
    or_open_file2();
}

fn or_open_file() {
    let f = File::open("hello.txt").unwrap_or_else(|error| {
        if error.kind() == ErrorKind::NotFound {
            File::create("hello.txt").unwrap_or_else(|error| {
                panic!("Problem creating the file: {:?}", error);
            })
        } else {
            panic!("Problem opening the file: {:?}", error);
        }
    });
}

fn or_open_file2() {
    // 如果文件不存在直接panic!
    let f = File::open("hello02.txt").unwrap();

    // expect 提供一个好的错误信息可以表明你的意图并易于追踪panic的根源
    let f = File::open("hello02.txt").expect("Failed to open hello.txt");
}
