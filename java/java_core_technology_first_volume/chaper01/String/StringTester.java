public class StringTester {
	public static void main(String[] args) {
		String greeting = "Hello";
		String s = greeting.substring(0, 3);
		System.out.println("s: " + s);

		String expletive = "Expletive";
		String PG13 = "deleted";
		String message = expletive + PG13;
		System.out.println("expletive + PG13 : " + message);

		int age = 13;
		String rating = "PG" + age;
		System.out.println("The answer is " + rating);

		String all = String.join("/ " , "S", "M", "L", "XL");
		System.out.println("all is the string :" + all);
	}
}
