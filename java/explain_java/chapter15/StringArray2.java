// 字符串数组(逐个显示字符)

class StringArray2 {
	public static void main(String[] args) {
		String[] sx = { "Turbo", "NA", "DOHC" };

		for (int i = 0; i < sx.length; i++) {
			System.out.print("sx[" + i + "] = \"");
			for (int j = 0; j < sx[i].length(); j++)
				System.out.print(sx[i].charAt(j));
			System.out.println('\"');
		}
	}
}
