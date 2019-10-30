package snake;

import java.awt.Graphics;
import java.awt.Font;
import java.awt.Color;

import javax.swing.JPanel;

/**
 * 贪吃蛇游戏主窗口组件
 */
public class SnakeGameBodyPanel extends JPanel {


    public void paint(Graphics g) {
        // add(new SnakeGameHeadComponent());
        // set head
        g.fillRect(25, 11, 850, 50); 
          // set title
          g.setColor(Color.darkGray);
          g.setFont(new Font("arial", Font.BOLD, 25));
          g.drawString("snake", 420, 45);
          repaint();
  
          // set body
        //   g.setColor(Color.darkGray);
    }
}