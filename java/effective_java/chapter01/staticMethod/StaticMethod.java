package staticMethod;

import java.lang.Boolean;

public class StaticMethod {
	public static void main(String[] args) {
		if (args.length > 0) 
			System.out.printf(
					"argument %s is %s\n", 
					args[0], 
					valueOf(true) ? "true" : "false"
			);
		else
			System.out.println("error!");
	}

	public static Boolean valueOf(boolean b) {
		return b ? Boolean.TRUE : Boolean.FALSE;
	}
}
