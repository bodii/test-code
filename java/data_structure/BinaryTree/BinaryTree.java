
/**
 * 二叉树节点类
 */
public class BinaryTree {
    public TreeNode root; // 根节点

    public BinaryTree() {
        this.root = null;
    }

    public BinaryTree(TreeNode root) {
        this.root = root;
    }

    // 判断二叉树是否为空
    public boolean isEmpty() {
        if (root == null) {
            return true;
        } else {
            return false;
        }
    }

    public String toString() {
        return root.toString();
    }
}