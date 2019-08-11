import java.io.Console;

public class InputPasswordTester {
	public static void main(String[] args) {
		Console cons = System.console();
		String username = cons.readLine("User name: ");
		char[] passwd = cons.readPassword("Password: ");
		System.out.println("username: " + username + " password： " + passwd);
	}
}
