package snake;

import java.awt.Graphics;
import java.awt.Font;
import java.awt.Dimension;
import java.awt.Color;
import java.awt.Point;
import java.awt.BorderLayout;
import java.awt.GridLayout;

import javax.swing.JLabel;
import javax.swing.JPanel;

/**
 * SnakeGameHeadPanel class
 * 贪吃蛇 - 主窗口 - 头部组件
 */
public class SnakeGameHeadPanel extends JPanel {
    public static final int DEFULT_WIDTH = 850;
    public static final int DEFAULT_HEIGHT = 100;

    public void paintComponent(Graphics g) {
        // super.paintComponent(g);
        // 底色
        g.fillRect(25,11, DEFULT_WIDTH, DEFAULT_HEIGHT);

        // 标题
        JPanel titlePanel = new JPanel();
        titlePanel.setLayout(new GridLayout(1, 1));
        JLabel title = new JLabel("snake");
        title.setForeground(Color.WHITE);
        title.setOpaque(true);
        title.setFont(new Font("arial", Font.BOLD, 25));
        titlePanel.add(title);
        add(titlePanel, BorderLayout.CENTER);
    }

}