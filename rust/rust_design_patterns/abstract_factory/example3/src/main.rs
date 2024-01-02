trait Logger {
    fn log(&self, message: &str);
}

struct ConsoleLogger;

impl Logger for ConsoleLogger {
    fn log(&self, message: &str) {
        println!("Console Log: {}", message);
    }
}

struct FileLogger;

impl Logger for FileLogger {
    fn log(&self, message: &str) {
        println!("File Log: {}", message);
    }
}

trait LoggerFactory {
    fn create_logger(&self) -> Box<dyn Logger>;
}

struct ConsoleLoggerFactory;

impl LoggerFactory for ConsoleLoggerFactory {
    fn create_logger(&self) -> Box<dyn Logger> {
        Box::new(ConsoleLogger)
    }
}

struct FileLoggerFactory;

impl LoggerFactory for FileLoggerFactory {
    fn create_logger(&self) -> Box<dyn Logger> {
        Box::new(FileLogger)
    }
}

fn main() {
    let console = ConsoleLoggerFactory;
    let console_logger = console.create_logger();

    let file = FileLoggerFactory;
    let file_logger = file.create_logger();

    console_logger.log("This is a console log.");
    file_logger.log("This is a file log.");
}
