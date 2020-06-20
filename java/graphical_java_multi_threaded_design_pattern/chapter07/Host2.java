package chapter07;

public class Host2 {

	private final Helper helper = new Helper();

	public void request(final int count, final char c) {
		System.out.println("	request(" + count + ", " + c + ") BEGIN");

		new Thread (
			new Runnable() {
				@Override
				public void run() {
					helper.handle(count, c);
				}
			}
		).start();

		System.out.println("	request(" + count + ", " + c + ") END");
	}

}
