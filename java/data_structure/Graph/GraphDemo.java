/**
 * 图的demo类
 */
public class GraphDemo {
    public static void main(String[] args) {
        // 创建并打印无权图
        createNoRightGraph();
        // 创建并打印有权图
        createRightGraph();
    }

    // 创建并打印无权图
    public static void createNoRightGraph() {
        // 创建顶点
        GraphNode node1 = new GraphNode("A");
        GraphNode node2 = new GraphNode("B");
        GraphNode node3 = new GraphNode("C");
        GraphNode node4 = new GraphNode("D");
        GraphNode node5 = new GraphNode("E");
        GraphNode node6 = new GraphNode("F");

        // 创建图边
        GraphBorder border1 = new GraphBorder(node1, node2);
        GraphBorder border2 = new GraphBorder(node1, node5);
        GraphBorder border3 = new GraphBorder(node2, node3);
        GraphBorder border4 = new GraphBorder(node3, node4);
        GraphBorder border5 = new GraphBorder(node2, node4);
        GraphBorder border6 = new GraphBorder(node5, node6);

        // 创建图顶点集合
        GraphNode[] nodes = new GraphNode[6];
        nodes[0] = node1;
        nodes[1] = node2;
        nodes[2] = node3;
        nodes[3] = node4;
        nodes[4] = node5;
        nodes[5] = node6;

        // 创建图边集合
        GraphBorder[] borders = new GraphBorder[6];
        borders[0] = border1;
        borders[1] = border2;
        borders[2] = border3;
        borders[3] = border4;
        borders[4] = border5;
        borders[5] = border6;

        // 实例化一个无权图
        NoRightGraph g = new NoRightGraph(nodes, borders);
        g.printMatrix(); // 打印
    }

    // 创建并打印有权图
    public static void createRightGraph() {
        // 创建顶点
        GraphNode node1 = new GraphNode("A");
        GraphNode node2 = new GraphNode("B");
        GraphNode node3 = new GraphNode("C");
        GraphNode node4 = new GraphNode("D");
        GraphNode node5 = new GraphNode("E");
        GraphNode node6 = new GraphNode("F");

        // 创建图边
        RightGraphBorder border1 = new RightGraphBorder(node1, node2, 10);
        RightGraphBorder border2 = new RightGraphBorder(node1, node5, 20);
        RightGraphBorder border3 = new RightGraphBorder(node2, node3, 30);
        RightGraphBorder border4 = new RightGraphBorder(node3, node4, 40);
        RightGraphBorder border5 = new RightGraphBorder(node2, node4, 50);
        RightGraphBorder border6 = new RightGraphBorder(node5, node6, 60);

        // 创建图顶点集合
        GraphNode[] nodes = new GraphNode[6];
        nodes[0] = node1;
        nodes[1] = node2;
        nodes[2] = node3;
        nodes[3] = node4;
        nodes[4] = node5;
        nodes[5] = node6;

        // 创建图边集合
        RightGraphBorder[] borders = new RightGraphBorder[6];
        borders[0] = border1;
        borders[1] = border2;
        borders[2] = border3;
        borders[3] = border4;
        borders[4] = border5;
        borders[5] = border6;

        // 实例化一个有权图
        RightGraph g = new RightGraph(nodes, borders);
        g.printMatrix(); // 打印
    }
}