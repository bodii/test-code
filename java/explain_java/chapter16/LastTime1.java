// 显示上次的心情

import java.io.*;
import java.util.Scanner;

class LastTime1 {

	// -- 读入上次的心情
	static void init() {
		BufferedReader br = null;

		try {
			br = new BufferedReader(new FileReader("LastTime.txt"));
			String kibun = br.readLine();
			System.out.println("上次的心情" + kibun + "。");
		} catch (IOException e) {
			System.out.println("这是你第1次运行本程序。");
		} finally {
			if (br != null) {
				try {
					br.close();
				} catch (IOException e) {
					System.out.println("文件关闭失败。");
				}
			}
		}
	}

	// -- 读入此次的心情
	static void term(String kibun) {
		FileWriter fw = null;

		try {
			fw = new FileWriter("LastTime.txt");
			fw.write(kibun);
		} catch (IOException e) {
			System.out.println("发生错误!!");
		} finally {
			if (fw != null) {
				try {
					fw.close();
				} catch (IOException e) {
					System.out.println("文件关闭失败。");
				}
			}
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		init();

		System.out.println("当前的心情: ");
		String kibun = stdIn.next();

		term(kibun);
	}
}
