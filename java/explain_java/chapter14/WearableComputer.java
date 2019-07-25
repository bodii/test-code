// -- 可穿戴的计算机类
public class WearableComputer implements Wearable {
	private String name;

	public WearableComputer(String name) { this.name = name; }

	public void putOn() { System.out.println(name + " ON!!"); }
	public void putOff() { System.out.println(name + " OFF!!"); }
}
