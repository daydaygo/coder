// https://fasterthanli.me/articles/a-half-hour-to-learn-rust
use std::cmp;

fn quick_start() {
    let x: i32 = 42;
    let x = x + 3; // can redefine; can not: x = x+3;
    let _ = 42;
    {
        // block scope
        let x = "out"; // another x
    }

    let pair = ('a', 17);
    let pair: (char, i32) = ('a', 17);
    pair.0; // tuple
    let (left, right) = slice.split_at(middle); // fn return tuple

    let b: bool = true;
    if b {
    } else {
    }
    match b {
        true => 6,
        _ => {}
    };

    let least = cmp::min(3, 8); // :: for namespace
    let x = "quick start".len(); // type are namespace too
    let v = Vec::new(); // Vec: regular struct
}

fn qs_foo() -> i32 {
    // -> return type
    4
}
