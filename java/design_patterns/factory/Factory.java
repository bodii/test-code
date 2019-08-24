package factory;

public class Factory {
	public JVMServers getServerByVendor(String name) {
		if (name.equals("Apache")) {
			return new Tomcat();
		}
		else if (name.equals("Eclipse")) {
			return new Jetty();
		} 
		else if (name.equals("RedHat")) {
			return new WildFly();
		}
		else {
			return null;
		}
	}
}
