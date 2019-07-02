// 随机生成并显示大于等于0.0小于1.0的实数值

import java.util.Random;

class RandomDoubleTest01 {
	public static void main(String[] args) {
		Random rand = new Random();

		System.out.println("大于等于0.0小于1.0的实数值:" + rand.nextDouble());
	}
}
