use gui::Draw;
use gui::{Screen, Button};

#[derive(Debug)]
#[allow(dead_code)]
struct SelectBox {
    width: u32,
    height: u32,
    options: Vec<String>,
}

impl Draw for SelectBox {
    // SelectBox实现了Draw trait，就意味着它实现了draw方法
    fn draw(&self) {
        println!("{:?}", &self);
    }
}

fn main() {
    let screen = Screen {
        components: vec![
            Box::new(SelectBox{
                width: 75,
                height: 10,
                options: vec![
                    String::from("Yes"),
                    String::from("Maybe"),
                    String::from("No")
                ],
            }),
            Box::new(Button{
                width: 50,
                height: 10,
                label: String::from("Ok"),
            }),
        ],
    };

    screen.run();
}
