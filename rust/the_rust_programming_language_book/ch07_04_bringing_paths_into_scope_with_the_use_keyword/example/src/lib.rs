mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
    }
}

// use crate::front_of_house::hosting;

pub fn eat_at_restaurant() {
    hosting::add_to_waitlist();
}

mod customer {
    use crate::front_of_house::hosting;

    pub fn eat_at_restaurant() {
        hosting::add_to_waitlist();
    }
}

// 使用父模块可以区分这两个Result类型
// use std::fmt;
// use std::io;
// fn function1() -> fmt::Result {}
// fn function2() -> io::Result<()> {}

// 通过as重命名同名类型
// use std::fmt::Result;
// use std::io::Result as IoResult;
// fn function1() -> Result {}
// fn function2() -> IoResult<()> {}

// 使用pub use 重导出名称
pub use crate::front_of_house::hosting;
pub fn eat_at_restaurant2() {
    hosting::add_to_waitlist();
}
