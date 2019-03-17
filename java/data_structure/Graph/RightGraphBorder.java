/**
 * 有权图的边类
 */
public class RightGraphBorder {
	public GraphNode fNode; // 边的第一个顶点
	public GraphNode lNode; // 边的第二个顶点
	public int power; // 权值

	public RightGraphBorder(GraphNode fNode, GraphNode lNode, int power) {
		this.fNode = fNode;
		this.lNode = lNode;
		this.power = power;
	}
}
