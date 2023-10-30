use std::fs;
use std::fs::File;
use std::io;
use std::io::Read;

fn main() {
    let f = File::open("hello.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

// 传播错误的简写： ？运算符
/*
    Result值之后?被定义的处理Result值的match表达式有着完全相同的工作方式。
    如果Result的值是Ok,这个表达式将会返回Ok中的值而程序将继续执行。
    如果值是Err, Err将作为整个函数的返回值，就好像使用了return关键字一样，这样错误
    值就被传播给了调用者。
*/
fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

/* ？之后直接使用链式方法调用来进一步缩短代码 */
fn read_username_from_file2() -> Result<String, io::Error> {
    let mut s = String::new();

    File::open("hello.txt")?.read_to_string(&mut s)?;

    Ok(s)
}

fn read_username_from_file3() -> Result<String, io::Error> {
    fs::read_to_String("hello.txt")
}

// 一个有效的返回值是()，同时出于方便，另一个有效的返回值是Result<T, E>
// Box<dyn Error>为使用?时允许返回的“任意类型的错误”
use std::error::Error;
fn read04() -> Result<(), Box<dyn Error>> {
    let f = File::open("hello.txt")?;

    Ok(())
}
