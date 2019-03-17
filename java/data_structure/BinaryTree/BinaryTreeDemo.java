public class BinaryTreeDemo {
	public static void main(String[] args) {
		// TreeNode tn = new TreeNode("A");
		System.out.println(create());
	}

	/**
	 * 创建二叉树方法 ..................a ...........b ...................c
	 * ....d..........e .f ....... g
	 * 
	 */
	public static BinaryTree create() {
		// 初始化节点

		TreeNode f = new TreeNode("F");
		TreeNode g = new TreeNode("G");

		TreeNode d = new TreeNode("D");
		TreeNode e = new TreeNode("E", f, g);

		TreeNode b = new TreeNode("B", d, e);
		TreeNode c = new TreeNode("C");

		TreeNode a = new TreeNode("A", b, c);

		BinaryTree bt = new BinaryTree(a);
		return bt;
	}
}
