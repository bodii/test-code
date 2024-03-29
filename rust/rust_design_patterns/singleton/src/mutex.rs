use std::sync::Mutex;

static ARRAY: Mutex<Vec<i32>> = Mutex::new(Vec::new());

fn do_a_call() {
    ARRAY.lock().unwrap().push(1);
}

pub fn run_mutex() {
    do_a_call();
    do_a_call();
    do_a_call();

    println!("Called {} times", ARRAY.lock().unwrap().len());
}
