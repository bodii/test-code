import java.nio.file.Paths;
import java.util.Scanner;
import java.io.PrintWriter;
import java.io.IOException;

public class FileInOut {
	public static void main(String[] args) throws IOException {
		// 定义输出流位置，如果文件不存在则创建
		PrintWriter stdOut = new PrintWriter("myfile.txt", "UTF-8");
		//  
		Scanner stdIn = new Scanner(Paths.get("myfile.txt"), "UTF-8");

		stdOut.println("hello"); // 写入内容
		stdOut.close(); // 关闭流
	}
}
