// -- 测试
public class Test {
	public static void main(String[] args) {
		Wearable[] w = {
			new WearableComputer("HAL"),
			new WearableRobot(Color.RED),
			new WearableRobot(Color.GREEN),
		};

		for (Wearable k : w) {
			k.putOn();
			k.putOff();
			System.out.println();
		}
	}
}


