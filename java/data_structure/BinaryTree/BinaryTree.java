import java.util.Vector;

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

    // 先序遍历二叉树
    public static Vector rootFirst(TreeNode root) {
        Vector result = new Vector();
        if (root == null) {
            return result;
        }

        result.add(root);
        Vector leftChild = rootFirst(root.left);
        Vector rightChild = rootFirst(root.right);
        result.addAll(leftChild);
        result.addAll(rightChild);

        return result;
    }

    // 中序遍历二叉树
    public static Vector rootMid(TreeNode root) {
        Vector result = new Vector();
        if (root == null) {
            return result;
        }

        Vector leftChild = rootFirst(root.left);
        result.addAll(leftChild);

        result.add(root);

        Vector rightChild = rootFirst(root.right);
        result.addAll(rightChild);

        return result;
    }

    // 后序遍历二叉树
    public static Vector rootLast(TreeNode root) {
        Vector result = new Vector();
        if (root == null) {
            return result;
        }

        Vector leftChild = rootFirst(root.left);
        result.addAll(leftChild);

        Vector rightChild = rootFirst(root.right);
        result.addAll(rightChild);

        result.add(root);

        return result;
    }
}