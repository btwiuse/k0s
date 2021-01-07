extern crate rand;

fn print_rand() {
    println!("char: {}", rand::random::<char>());
}

fn main() {
    println!("Hello, world!");
    print_rand::<>();
}
