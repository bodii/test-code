// 将通过命令行参数传入的所有数值相加并显示

class SumOfArgs {
	public static void main(String[] args) {
		double sum = 0.0;
		for (int i = 0; i < args.length; i++)
			sum += Double.parseDouble(args[i]);

		System.out.println("合计值为" + sum + "。");
	}
}
