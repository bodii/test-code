package staticMethod;

import java.lang.Boolean;

public class StaticMethodTest2 {
	public static void main(String[] args) {
		if (args.length > 0) 
			System.out.printf(
					"argument %s is %s\n", 
					args[0], 
					StaticMethod2.valueOf(true) ? "true" : "false"
			);
		else
			System.out.println("error!");
	}
}

class StaticMethod2 {
	public static Boolean valueOf(boolean b) {
		return b ? Boolean.TRUE : Boolean.FALSE;
	}
}