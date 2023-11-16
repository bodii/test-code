fn main() {
    let s = Some(String::from("Hello!"));

    if let Some(ref _s) = s {
        println!("found a string");
    }
    println!("{:?}", s);


    let s = Some(String::from("Hello!"));
    if let Some(_) = s {
        println!("found a string");
    }
    println!("{:?}", s);
}
