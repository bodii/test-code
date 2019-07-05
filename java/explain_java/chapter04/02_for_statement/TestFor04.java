// 显示身高和标准体重的对应表。
// 显示的身高范围(开始值、结束值、增量)需要作为整数值读入。
// 标准体重的计算公式为(身高 - 100) x 0.9

class TestFor04 {
	public static void main(String[] args) {
		int start = 150, end = 190, increment = 5;

		System.out.println("从多少cm开始: " + start +
				"\n到多少cm结束: " + end +
				"\n每次增量多少cm: " + increment +
				"\n身高\t标准体重");

		for (int i = 150; i <= 190; i+=5) 
			System.out.println(i + "\t" + ((i - 100) * 0.9));
		

	}
}
