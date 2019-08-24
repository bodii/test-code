package factory;

public class WildFly implements JVMServers {
	public void startServer() {
		System.out.println("Starting WildFly Server...");
	}

	public void stopServer() {
		System.out.println("Shutting Down WildFly Server...");
	}

	public String getName() {
		return "WildDly";
	}
}
