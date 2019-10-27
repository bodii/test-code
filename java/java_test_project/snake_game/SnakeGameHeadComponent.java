package snake;

import java.awt.Graphics;
import java.awt.Font;
import java.awt.Dimension;
import java.awt.Color;

import javax.swing.JComponent;

/**
 * SnakeGameHeadPanel class
 * 贪吃蛇 - 主窗口 - 头部组件
 */
public class SnakeGameHeadConponent extends JComponent {
    public static final int DEFULT_WIDTH = 800;
    public static final int DEFAULT_HEIGHT=50;

    public void paintComponent(Graphics g) {
        super.paintComponent(g);
        
        // 底色
        g.fillRect(25, 11, DEFULT_WIDTH, DEFAULT_HEIGHT);

        // 标题
        g.setColor(Color.DARK_GRAY);
        setBackground(Color.DARK_GRAY);
        g.setFont(new Font("arial", Font.BOLD, 25));
        g.drawString("snake", 420, 45);
    }

}