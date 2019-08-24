package factory;

public class Tomcat implements JVMServers {
	public void startServer() {
		System.out.println("Starting Tomcat Server...");
	}

	public void stopServer() {
		System.out.println("Shutting Down Tomcat Server...");
	}

	public String getName() {
		return "Tomcat";
	}
}
