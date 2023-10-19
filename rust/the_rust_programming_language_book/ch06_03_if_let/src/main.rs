#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn main() {
    let some_u8_value = Some(0u8);
    match some_u8_value {
        Some(3) => println!("three"),
        _ => (),
    }

    // 使用if 达到一致的效果
    if let Some(3) = some_u8_value {
        println!("three")
    }

    let mut count = 0;
    let coin = Coin::Dime;
    match coin {
        Coin::Quarter(state) => println!("State quarter from {:?}!", state),
        _ => count += 1,
    }

    let coin = Coin::Dime;
    if let Coin::Quarter(state) = coin {
        println!("State quarter from {:?}!", state);
    }else {
        count += 1;
    }
}
