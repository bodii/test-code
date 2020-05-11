package chapter17;

import java.io.IOException;
import java.util.logging.ConsoleHandler;
import java.util.logging.FileHandler;
import java.util.logging.Formatter;
import java.util.logging.Level;
import java.util.logging.LogRecord;
import java.util.logging.Logger;

public class TestLogger {
    public static void main(String[] args) throws IOException {
        Logger log = Logger.getLogger("test_log");
        log.setLevel(Level.INFO);
        log.info("aa");
        log.info("bb");
        log.fine("fine");

        // 控制台控制器
        ConsoleHandler consoleHandler = new ConsoleHandler();
        consoleHandler.setLevel(Level.ALL);
        log.addHandler(consoleHandler);

        try {
            // 文件控制器
            FileHandler fileHandler = new FileHandler("./test_log_file.log");
            fileHandler.setLevel(Level.INFO);
            fileHandler.setFormatter(new Formatter() { // 匿名类
                @Override
                public String format(LogRecord record) {
                    return record.getLevel() + ":" + record.getMillis()  + ":" + record.getMessage() + "\n";
                }
            });
            log.addHandler(fileHandler);
            log.info("aaaa");
            log.info("bbbb");
            log.fine("fine");
        } catch (SecurityException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}