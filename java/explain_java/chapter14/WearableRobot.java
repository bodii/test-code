// -- 可穿戴机器人类
public class WearableRobot implements Wearable, Color {
	private int color;

	public WearableRobot(int color) { changeColor(color); }

	public void changeColor(int color) { this.color = color; }

	public String toString() {
		switch (color) {
			case RED	: return "红色机器人"; 
			case GREEN  : return "绿色机器人";
			case BLUE	: return "蓝色机器人";
		}

		return "机器人";
	}

	public void putOn() {
		System.out.println(toString() + " 戴上!!");
	}

	public void putOff() {
		System.out.println(toString() + " 脱下!!");
	}
}
