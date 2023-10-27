use std::collections::HashMap;

fn main() {
    /*
    给定一系列数字，使用 vector 并返回这个列表的平均数（mean, average）、
    中位数（排列数组后位于中间的值）和众数（mode，出现次数最多的值；
    这里哈希 map 会很有帮助）。
    */

    let number_list = vec![1, 2, 3, 4, 5, 6, 7];
    let mut mean = 0;
    for i in &number_list {
        mean += i;
    }

    mean /= number_list.len();
    println!("mean: {}", mean);

    let len = number_list.len();
    let median: usize;
    // print_type_of(&number_list[len/2-1]);
    if  len % 2 == 1 {
        median = *&number_list[len/2];
    } else {
        median = (*&number_list[len/2-1] + *&number_list[len/2]) / 2
    }
    println!("median: {}", median);
    // println!("{:?}", number_list);

    let mut mode_map = HashMap::new();
    for i in &number_list {
        mode_map.entry(i).and_modify(|counter| *counter += 1).or_insert(1);
    }
    println!("mode: {:?}", mode_map);
}

fn _print_type_of<T>(_: &T) {
    println!("{}", std::any::type_name::<T>() );
}
