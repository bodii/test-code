use std::cmp::PartialOrd;

// 像i32和char这样的类型是已知大小的并可以储存在栈上，所以他们实现了Copy trait.
// 如果并不希望限制largest函数只能用于实现Copy trait的类型，我们可以在T的trait
// bounds 中指定Clone而不Copy。并克隆slice的每一个值使得largest函数拥有其所有权。
// 使用clone函数意味着对于类型String这样拥有堆上数据的类型，会潜在的分配更多堆上
// 空间，而堆分配在涉及大量数据时可能会相当缓慢。
//
// 另一种largest的实现方式返回在slice中T值的引用。如果我们将函数返回值从T 改为&T
// 并改变函数体使其能够返回一个引用，我们将不需要任何Clone或Copy的trait bounds而且
// 也将不会有任何的堆分配。
fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn main() {
    let number_list = vec![34, 50, 25, 100, 65];
    let result = largest(&number_list);
    println!("The largest number is {}", result);

    let char_list = vec!['y', 'm', 'a', 'q'];
    let result = largest(&char_list);
    println!("The largest char is {}", result);
}
