// 一个vec！宏定义的简化版本
#[macro_export]
macro_rules! vec2 {
    ( $( $x:expr ),* ) => {
        {
            let mut temp_vec = Vec::new();
            $(
                temp_vec.push($x);
            )*
            temp_vec
        };

    }
}   

#[cfg(test)]
mod tests {
    #[test] 
    fn test_macro() {
        let x = vec2![1, 2, 3];
        assert_eq!(vec![1,2,4], x);
    } 
}
