fn main() {
    /*
    将字符串转换为 Pig Latin，也就是每一个单词的第一个辅音字母被移动到单词
    的结尾并增加 “ay”，所以 “first” 会变成 “irst-fay”。元音字母开头的单词则在结
    尾增加 “hay”（“apple” 会变成 “apple-hay”）。牢记 UTF-8 编码！
    */

    let mut first = "first".to_string();
    println!("original word first: {}", first);
    first = modify_contain_vowels_word(first);
    println!("modified word first: {}", first);

    let mut apple = "apple".to_string();
    println!("original word apple: {}", apple);
    apple = modify_contain_vowels_word(apple);
    println!("modified word apple: {}", apple);
}

fn modify_contain_vowels_word(mut s: String) -> String {
    if has_contain_vowels(s.clone()) {
        s.push_str("-hay");
    } else {
        s = format!("{}-{}{}", &s[1..], &s[0..1], "ay");
    }

    s
}

fn has_contain_vowels(s: String) -> bool {
    let vowel = vec!['a', 'e', 'i', 'o', 'u'];
    for v in &vowel {
        if s.find(*v) == Some(0) {
            return true; 
        }
    }

    false
}
