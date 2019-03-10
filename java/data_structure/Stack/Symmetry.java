/**
 * 验证符号是否平衡对称
 */

public class Symmetry {
	private static SequenceStack ifSymmetry(String args) {
		SequenceStack ts = new SequenceStack(10);
		int size = args.length();

		for (int i = 0; i < size; i++) {
			char charset = args.charAt(i);
			switch (charset) {
			case '(':
			case '[':
			case '{':
				ts.push(charset);
				break;
			case ')':
			case ']':
			case '}':
				ts.pop();
				break;
			default:
				continue;
			}
		}

		return ts;
	}

	public static void main(String[] args) {
		String arg = "(aa[bb{cc}])";
		SequenceStack st = ifSymmetry(arg);

		System.out.println("size is: " + st.size());
	}
}
