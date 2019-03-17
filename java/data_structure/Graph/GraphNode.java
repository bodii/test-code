/**
 * 图的节点类
 */
public class GraphNode {
	public Object Node;

	public GraphNode(Object value) {
		this.Node = value;
	}

	public String toString() {
		if (Node == null) {
			return "";
		} else {
			return Node.toString();
		}
	}
}
