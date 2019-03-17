/**
 * 树节点类
 */
public class TreeNode {
	public Object nodeValue;
	public TreeNode left, right;

	// 类构造方法
	public TreeNode() {
		this(null, null, null);
	}

	// 类构造方法
	public TreeNode(Object o) {
		this(o, null, null);
	}

	// 类构造方法
	public TreeNode(Object o, TreeNode left, TreeNode right) {
		this.nodeValue = o;
		this.left = left;
		this.right = right;
	}

	// 判断是否是叶子节点
	public boolean isLeaf() {
		if (left == null && right == null) {
			return true;
		} else {
			return false;
		}
	}

	// 重写类的toString方法
	public String toString() {
		// 如果这个一个空节点, 直接返回空字符串
		if (nodeValue == null) {
			return "";
		}

		// 返回节点信息
		String result = " (node " + nodeValue.toString();

		if (left != null) {
			result += "   left child node:" + left.toString();
		}

		if (right != null) {
			result += "   right child node:" + right.toString();
		}

		result += ")\n";

		return result;
	}
}
