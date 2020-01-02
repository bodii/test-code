package chapter04;

public class Test02 {
	public static void main(String[] args) {
		System.out.println(Utensil.NAME + Dessert.NAME);
	}

	private static class Utensil {
		static final String NAME = "pan";
	}

	private static class Dessert {
		static final String NAME = "cake";
	}
}

/**
* ! 永远不要把多个顶级类或者接口放在一个源文件中。
*
*/
