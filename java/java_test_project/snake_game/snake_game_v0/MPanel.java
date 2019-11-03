package snake;

import java.awt.Graphics;
import java.awt.Color;
import java.awt.Font;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.util.Random;
import java.util.ArrayList;

import javax.swing.JPanel;
import javax.swing.ImageIcon;


/**
 * MPanel class
 * 贪吃蛇画布
 */
public class MPanel extends JPanel implements KeyListener {

    private ImageIcon snakeBody = new ImageIcon("body.png");
    private ImageIcon snakeHeadUp = new ImageIcon("up.png");
    private ImageIcon snakeHeadDown = new ImageIcon("down.png");
    private ImageIcon snakeHeadLeft = new ImageIcon("left.png");
    private ImageIcon snakeHeadRight = new ImageIcon("right.png");
    private ImageIcon food = new ImageIcon("food.png");

    // 开始前的提示信息
    public boolean isStarted = false;

    public final int latticeSize = 25; // 每个格子的大小
    public int bodyLen = 2;
    public int[] snakeHeadCoordinate = new int[2];
    public int[] snakeBodyCoordinateX = new int[100];
    public int[] snakeBodyCoordinateY = new int[100];
    public ImageIcon snakeHeadDerection = snakeHeadRight;

    public MPanel() {
        initSnake();
        setFocusable(true); // 添加焦点
        addKeyListener(this); // 添加键盘事件监听
    }

    /**
     * 画笔
     */
    public void paintComponent(Graphics g) {
        super.paintComponent(g); // 先调用父类
        setBackground(Color.WHITE); // 设置背景色

        // title 
        g.fillRect(25, 11, 850, 50);
        g.setColor(Color.white);
        g.setFont(new Font("arial", Font.BOLD, 25));
        g.drawString("snake", 420, 45);

        // body
        g.setColor(new Color(50, 50, 50));
        g.fillRect(25, 70, 850, 600);

        if (!isStarted) {
            // set start message
            g.setColor(Color.WHITE);
            g.setFont(new Font("arial", Font.BOLD, 25));
            g.drawString("Press Space to Start", 350, 350);
        }

        snakeHeadDerection.paintIcon(this, g, snakeHeadCoordinate[0], snakeHeadCoordinate[1]);
        for (int i = 0; i < bodyLen; i++) {
            snakeBody.paintIcon(this, g, snakeBodyCoordinateX[i], snakeBodyCoordinateY[i]);
        }
    }

    public void initSnake() {
        bodyLen = 2;
        snakeHeadCoordinate[0] = 100;
        snakeHeadCoordinate[1] = 100;
        for (int i = 0; i < bodyLen; i++) {
            snakeBodyCoordinateX[i] =  100  -  25 * (i +1);
            snakeBodyCoordinateY[i] = 100;
        }
    }

    @Override
    public void keyTyped(KeyEvent e) {

    }

    @Override
    public void keyPressed(KeyEvent e) {
        int keyCode = e.getKeyCode();
        if (keyCode == KeyEvent.VK_SPACE) {
            isStarted = !isStarted;
            repaint();
        }

        if (isStarted) {
            if (keyCode == KeyEvent.VK_UP || keyCode == KeyEvent.VK_K) {
                snakeHeadDerection = snakeHeadUp;
                repaint();
            } else if (keyCode == KeyEvent.VK_DOWN || keyCode == KeyEvent.VK_J) {
                snakeHeadDerection = snakeHeadDown;
                repaint();
            } else if (keyCode == KeyEvent.VK_LEFT || keyCode == KeyEvent.VK_H) {
                snakeHeadDerection = snakeHeadLeft;
                repaint();
            } else if (keyCode == KeyEvent.VK_RIGHT || keyCode == KeyEvent.VK_L) {
                snakeHeadDerection = snakeHeadRight;
                repaint();
            }
        }
    }

    @Override
    public void keyReleased(KeyEvent e) {

    }
}