mod conceptual;
mod serdes;

use conceptual::run_conceptual;
use serdes::run_serdes;

fn main() {
    run_conceptual();
    run_serdes();
}
