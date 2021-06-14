use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn guess() {
    println!("guess game");

    let num = rand::thread_rng().gen_range(1, 101);
    println!("num: {}", num);

    loop {
        println!("input you guess: ");
        let mut guess = String::new();
        io::stdin().read_line(&mut guess).expect("fail");
        println!("you guess: {}", guess);
        // shadowing: change var type
        let guess: u32 = match guess.trim().parse() {
            // handle err
            Ok(n) => n,
            Err(_) => continue,
        };

        match guess.cmp(&num) {
            Ordering::Less => {
                println!("less");
            }
            Ordering::Equal => {
                println!("win");
                break;
            }
            Ordering::Greater => {
                println!("big");
            }
        }
    }
}
