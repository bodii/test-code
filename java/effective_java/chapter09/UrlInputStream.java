package chapter09;

import java.io.IOException;
import java.io.InputStream;
import java.net.URL;

public class UrlInputStream {
    public static void main(String[] args) throws IOException {
        try (InputStream in = new URL(args[0]).openStream()) {
            in.transferTo(System.out);
        }
    }
}