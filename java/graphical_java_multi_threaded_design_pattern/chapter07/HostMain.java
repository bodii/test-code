package chapter07;

public class HostMain {
	public static void main(String[] args) {
		System.out.println("main BEGIN");

		Host host = new Host();
		host.request(10, 'A');
		host.request(20, 'B');
		host.request(30, 'C');

		System.out.println("main END");
	}
}
