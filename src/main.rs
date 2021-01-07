extern crate rand;
extern crate serde_json;

fn print_rand() {
    println!("char: {}", rand::random::<char>());
}

// taken from https://crates.io/crates/serde_json # Constructing JSON values
fn print_json() {
    // The type of `john` is `serde_json::Value`
    let john = serde_json::json!({
        "name": "John Doe",
        "age": 43,
        "phones": [
            "+44 1234567",
            "+44 2345678"
        ]
    });

    println!("first phone number: {}", john["phones"][0]);

    // Convert to a string of JSON and print it out
    println!("{}", john.to_string());
}

fn main() {
    println!("Hello, world!");
    print_rand::<>();
    print_json::<>();
}
