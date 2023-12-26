mod conceptual;
mod functional;

use conceptual::run_conceptual;
use functional::run_functional;

fn main() {
    println!("conceptual:");
    run_conceptual();

    println!();

    println!("functional:");
    run_functional();
}
