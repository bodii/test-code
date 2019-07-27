// 使用String.format方法创建格式字符串

class StringFormat2 {
	public static void main(String[] args) {
		for (int i = 1; i <= 4; i++) 
			System.out.printf(String.format("%%%dd\n", i), 5);
	}
}
