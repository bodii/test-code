package factory;

public class Jetty implements JVMServers {
	public void startServer() {
		System.out.println("Starting Jetty Server...");
	}

	public void stopServer() {
		System.out.println("Shutting Down Jetty Server...");
	}

	public String getName() {
		return "Jetty";
	}
}
