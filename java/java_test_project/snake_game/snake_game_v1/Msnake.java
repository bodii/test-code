package snake;

import javax.swing.JFrame;

import snake.MPanel;

/**
 * Msnake class
 * 贪吃蛇游戏
 */
public class Msnake {
    public static void main(String[] args) {
        // 设置窗口
        JFrame frame = new JFrame();
        frame.setBounds(10, 10, 900, 720);
        frame.setTitle("snake game");
        frame.setResizable(false);  // 禁止拖动
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        frame.add(new MPanel());  // 将画布添加到窗口内

        frame.setVisible(true);
    }
}