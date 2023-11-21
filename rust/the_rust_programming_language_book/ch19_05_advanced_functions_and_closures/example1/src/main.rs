fn add_one(x: i32) -> i32 {
    x + 1
}

type AddOne = fn(i32) -> i32;

fn do_twice(f: AddOne, arg: i32) -> i32 {
    f(arg) + f(arg)
}

fn main() {
    let answer = do_twice(add_one, 5);

    println!("The answer is: {}", answer);
}
