#[derive(Debug)]
#[allow(dead_code)]
enum Status {
    Value(u32),
    Stop,
}

fn main() {
    let list_of_numbers = vec![1, 2, 3];
    let list_of_strings: Vec<String> = list_of_numbers
        .iter()
        .map(|i| i.to_string())
        .collect();
    println!("{:?}", list_of_strings);

    let list_of_numbers = vec![1, 2, 3];
    let list_of_strings: Vec<String> = list_of_numbers
        .iter()
        .map(ToString::to_string)
        .collect();
    println!("{:?}", list_of_strings);

    let list_of_status: Vec<Status> = 
        (0u32..20)
        .map(Status::Value)
        .collect();
    println!("{:?}", list_of_status);
}
