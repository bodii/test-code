enum IpAddrKind {
    V4,
    V6,
}

struct IpAddr {
    kind : IpAddrKind,
    address String,
}

// 更简洁的方式表达 IpAddrKind 和 IpAddr 相同的功能
enum IpAddrEnum {
    V4(String),
    V6(String),
}

enum IpAddrEnum2 {
    V4(u8, u8, u8, u8),
    V6(String),
}

// = struct Ipv4Addr 可以不加括号
struct Ipv4Addr {}
struct Ipv6Addr {}

enum IpAddrEnum3 {
    V4(Ipv4Addr),
    V6(Ipv6Addr),
}

enum Message {
    // Quit 没有关联任何数据
    Quit,
    // Move包含一个匿名结构体
    Move {x: i32, y: i32},
    // Write 包含单独一个String
    Write(String),
    // ChangeColor包含三个i32
    ChangeColor(i32, i32, i32),
}

// 枚举可以使用impl来为枚举定义方法
impl Message {
    fn call(&self) {}
}

fn main() {
    let four = IpAddrKind::V4;
    let six = IpAddrKind::V6;

    let home = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };

    let loopback = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };

    let home = IpAddrEnum::V4(String::from("127.0.0.1"));
    let loopback = IpAddrEnum::V6(String::from("::1"));

    let home = IpAddrEnum2::V4(127, 0, 0, 1);
    let loopback = IpAddrEnum2::V6(String::from("::1"));

    let m = Message::Write(String::from("hello"));
    m.call();


    // 一些包含数字类型和字符串类型Option值的例子:
    let some_number = Some(5);
    let some_string = Some("a string");

    let absent_number: Option<i32> = None;

    // Option<T>和T(T可以是任何类型)是不同的类型
    let x: i8 = 5;
    let y: Option<i8> = Some(5);

    // error
    // let sum = x + y;
}

fn route(ip_type: IpAddrKind) {}