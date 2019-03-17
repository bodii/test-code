/**
 * 图的节点类
 */
public class GraphNode {
	public Object Node;

	public GraphNode(Ojbect value) {
		this.Node = value;
	}

	public String toString() {
		return value == null ? "" : value.toString();
	}
}
