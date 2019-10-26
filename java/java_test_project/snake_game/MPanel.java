package snake;

import java.awt.Graphics;
import java.awt.Color;
import java.awt.Font;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;

import javax.swing.JPanel;
import javax.swing.ImageIcon;


/**
 * MPanel class
 * 贪吃蛇画布
 */
public class MPanel extends JPanel {
    // 食物
    private ImageIcon food = new ImageIcon("food.png");

    // 头的方向
    private ImageIcon headDirectIon  = Snake.getHead();

    // 开始前的提示信息
    public boolean isStarted = false;

    public MPanel() {

        // 获取焦点
        this.setFocusable(true);
        // 添加键盘监听事件
        this.addKeyListener(new GetKeyEvent(this));
    }

    /**
     * 画笔
     */
    public void paintComponent(Graphics g) {
        super.paintComponent(g); // 先调用父类
        this.setBackground(Color.WHITE); // 设置背景色

        // set head
        g.fillRect(25, 11, 850, 50); 

        // set title
        g.setColor(Color.WHITE);
        g.setFont(new Font("arial", Font.BOLD, 25));
        g.drawString("snake", 420, 45);

        // set body
        g.setColor(Color.darkGray);
        g.fillRect(25, 70, 850, 600); 

        //
        if (!isStarted) {
            // set start message
            g.setColor(Color.WHITE);
            g.setFont(new Font("arial", Font.BOLD, 25));
            g.drawString("Press Space to Start", 350, 350);
        }

        Snake.getHead().paintIcon(this, g, 100, 100);
        Snake.getBody().paintIcon(this, g, 75, 100);
        Snake.getBody().paintIcon(this, g, 50, 100);


    }

}


/**
 * Snake class
 * 蛇的相关类
 */
class Snake {
    // 身体
    private static final ImageIcon body = new ImageIcon("body.png");
    // 头
    private static final ImageIcon up = new ImageIcon("up.png");
    private static final ImageIcon down = new ImageIcon("down.png");
    private static final ImageIcon left = new ImageIcon("left.png");
    private static final ImageIcon right = new ImageIcon("right.png");

    // 方向
    private static ImageIcon direction = right;

    public static ImageIcon setHead(String d) {
        switch (d) {
            case "up" :
                direction = up;
                break;
            case "down":
                direction = down;
                break;
            case "left":
                direction = left;
                break;
            case "right":
                direction = right;
                break;
            default:
                direction = right;
        }

        return direction;
    }

    public static ImageIcon getHead() {
        return direction;
    }

    public static ImageIcon getBody() {
        return body;
    }

}


/**
 * GetKeyEvent class
 * 获取监听键盘事件
 */
class GetKeyEvent implements KeyListener {
    private MPanel s;
    // private String direction = "right"; // 方向

    public GetKeyEvent(MPanel ms) {
        s = ms;
    }

    /**
     * 点击
     * @param e
     */
    @Override
    public void keyTyped(KeyEvent e) {

    }

    /**
     * 击中
     * @param e
     */
    @Override
    public void keyPressed(KeyEvent e) {
        int keyCode  = e.getKeyCode();
        if (keyCode == KeyEvent.VK_SPACE) {
            s.isStarted = !s.isStarted; // 设置开始的值取反
            s.repaint(); // 重绘内容
        }

        if (s.isStarted) {
            // 如果是点中的左键或H键
            if (keyCode == KeyEvent.VK_LEFT || keyCode == KeyEvent.VK_H) {
                Snake.setHead("left");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_RIGHT || keyCode == KeyEvent.VK_L) {
                Snake.setHead("right");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_UP || keyCode == KeyEvent.VK_K) {
                Snake.setHead("up");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_DOWN || keyCode == KeyEvent.VK_J) {
                Snake.setHead("down");
                s.repaint(); // 重绘内容
            }
        }
    }

    /**
     * 抬起
     * @param e
     */
    @Override
    public void keyReleased(KeyEvent e) {

    }
}