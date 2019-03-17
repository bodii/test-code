/**
 * 图的边类
 */
public class GraphBorder {
	public GraphNode fNode; // 边的第一个顶点
	public GraphNode lNode; // 边的第二个顶点

	public GraphBorder(GraphNode fNode, GraphNode lNode) {
		this.fNode = fNode;
		this.lNode = lNode;
	}
}
