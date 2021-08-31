extern crate rand;
extern crate serde;
extern crate serde_json;

use serde::{Deserialize, Serialize};

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

#[derive(Serialize, Deserialize, Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn print_serde() {
    let point = Point { x: 1, y: 2 };

    // Convert the Point to a JSON string.
    let serialized = serde_json::to_string(&point).unwrap();

    // Prints serialized = {"x":1,"y":2}
    println!("serialized = {}", serialized);

    // Convert the JSON string back to a Point.
    let deserialized: Point = serde_json::from_str(&serialized).unwrap();

    // Prints deserialized = Point { x: 1, y: 2 }
    println!("deserialized = {:?}", deserialized);
}

// https://bastion.turbofish.rs/
fn turbofish() {
    let (oh, woe, is, me) = ("the", "Turbofish", "remains", "undefeated");
    let (_, _): (bool, bool) = (oh < woe, is > (me));
}

fn main() {
    println!("Hello, world!");
    turbofish();
    print_rand();
    print_json();
    print_serde();
}
