/**
 * 实现无权图的邻接矩阵
 */

public class NoRightGraph {
	private static final int conn = 1; // 无权图中连通的标记

	private GraphNode[] nodes; // 图中的顶点

	private int[][] adjacMatrix; // 矩阵

	public NoRightGraph(GraphNode[] nodes, GraphBorder[] borders) {
		this.nodes = nodes;
		this.adjacMatrix = new int[nodes.length][nodes.length];
		connect(borders);
	}

	// 把顶点之间的邻接关系写入矩阵
	private void connect(GraphBorder[] borders) {
		GraphNode fNode;
		GraphNode lNode;

		for (int i = 0; i < borders.length; i++) {
			fNode = borders[i].fNode;
			lNode = borders[i].lNode;

			int fIndex = findIndex(fNode);
			int lIndex = findIndex(lNode);
			if (fIndex >= 0 && fIndex < nodes.length && lIndex >= 0 && lIndex < nodes.length) {
				adjacMatrix[fIndex][lIndex] = conn;
				adjacMatrix[lIndex][fIndex] = conn;
			}
		}
	}

	// 查找顶点的位置
	private int findIndex(GraphNode node) {
		for (int i = 0; i < nodes.length; i++) {
			if (node.equals(nodes[i])) {
				return i;
			}
		}

		return -1;
	}

	// 打印邻接矩阵
	public void printMatrix() {
		System.out.println("该无权图的邻接矩阵为: ");
		System.out.print("   ");
		for (int i = 0; i < nodes.length; i++) {
			System.out.print(nodes[i] + "   ");

		}

		System.out.println();
		for (int j = 0; j < nodes.length; j++) {
			System.out.print(nodes[j] + "   ");
			for (int k = 0; k < nodes.length; k++) {
				System.out.print(adjacMatrix[j][k] + "   ");
			}
			System.out.println();
		}
	}
}
