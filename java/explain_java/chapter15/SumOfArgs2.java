// 将通过命令行参数传入的所有数值相加并显示

class SumOfArgs2 {
	public static void main(String[] args) {
		double sum = 0.0;

		for (String d : args) 
			sum += Double.parseDouble(d);
		System.out.println("合计值: " + sum + "。");
	}
}

