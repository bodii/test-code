// 确认表达式的求值顺序是左-->右

class EvaluationOrder {
	public static void main(String[] args) {
		int a = 3;
		int x = (a++) * (2 + a);
		System.out.println("a = " + a);
		System.out.println("x = " + x);
	}
}
