import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

public class TryWithResource {

    public static final int BUFFER_SIZE = 100;
    public static void main(final String[] args) {

    }

    public static String firstLineOfFile(final String path) throws IOException {
        try (BufferedReader br = new BufferedReader(new FileReader(path))) {
            return br.readLine();
        }
    }

    public static void copy(final String src, final String dst) throws IOException {
        try (
            InputStream in = new FileInputStream(src); 
            OutputStream out = new FileOutputStream(dst)
        ) {
            final byte[] buf = new byte[BUFFER_SIZE];
                int n;
                while((n = in.read(buf)) >= 0)
                    out.write(buf, 0, n);
            }
    }

    public static String firstLineOfFile(String path, String defaultVal) {
        try (BufferedReader br = new BufferedReader(new FileReader(path))) {
            return br.readLine();
        } catch (IOException e) {
            return defaultVal;
        }
    }
}