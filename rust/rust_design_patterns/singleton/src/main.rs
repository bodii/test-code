mod safe;
mod lazy;
mod mutex;


fn main() {
    safe::run_safe();
    lazy::run_lazy();
    mutex::run_mutex();
}
