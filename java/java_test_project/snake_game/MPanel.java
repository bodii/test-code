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

    private Snake sanke = new Snake();

    // 头的方向
    private ImageIcon headDirectIon = sanke.getHeadDirection();

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

        sanke.init(this, g);

    }

}


/**
 * Snake class
 * 蛇的相关类
 */
class Snake {
    // 身体
    private final ImageIcon body = new ImageIcon("body.png");
    // 头
    private final ImageIcon up = new ImageIcon("up.png");
    private final ImageIcon down = new ImageIcon("down.png");
    private final ImageIcon left = new ImageIcon("left.png");
    private final ImageIcon right = new ImageIcon("right.png");

    // 头部方向
    private ImageIcon headDirection = right;
    // 蛇头坐标
    private Coordinate headCoordinate;
    // 蛇身坐标
    private Coordinate[] bodyCoordinate; 
    private int bodyLen = 2; // 默认蛇的身体长度
    public final int latticeSize = 25; // 每个格子的大小

    public void setBodyLen(int l) {
        bodyLen = l;
    }

    /**
     * 设置蛇头的位置和方向
     * @param d
     * @return
     */
    public void setHead(MPanel p, Graphics g, Coordinate c) {
        headDirection.paintIcon(p, g, c.x, c.y);
    }


    private ImageIcon getHead() {
        return headDirection;
    }

    /**
     * 设置蛇头的位置
     * 
     * @param d
     * @return
     */
    public ImageIcon setHeadDirection(String d) {
        switch (d) {
            case "up" :
                headDirection = up;
                break;
            case "down":
                headDirection = down;
                break;
            case "left":
                headDirection = left;
                break;
            case "right":
                headDirection = right;
                break;
            default:
                headDirection = right;
        }

        return headDirection;
    }

    /**
     * 获取头的方向
     * 
     * @return
     */
    public ImageIcon getHeadDirection() {
        return headDirection;
    }

    public ImageIcon getBody() {
        return body;
    }

    public void setBody(MPanel p, Graphics g, Coordinate c) {
        body.paintIcon(p, g, c.x, c.y);
    }

    public void init(MPanel p, Graphics g) {
        bodyLen = 2;
        int headCoordinateInt = 100;
        headCoordinate = new Coordinate(headCoordinateInt);
        setHead(p, g, headCoordinate);
        for (int i = 1; i <= bodyLen; i++) {
            Coordinate b = new Coordinate(
                headCoordinateInt - latticeSize * i, 
                headCoordinateInt
            );
            bodyCoordinate[i] = b;
            setBody(p, g, b);
        }
    }
}


class Food {
    // 食物
    private ImageIcon food = new ImageIcon("food.png");
    private Coordinate coordinate = new Coordinate();

    Foo() {
        return food;
    }

    public hide() {

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
                Snake.setHeadDirection("left");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_RIGHT || keyCode == KeyEvent.VK_L) {
                Snake.setHeadDirection("right");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_UP || keyCode == KeyEvent.VK_K) {
                Snake.setHeadDirection("up");
                s.repaint(); // 重绘内容
            } else if (keyCode == KeyEvent.VK_DOWN || keyCode == KeyEvent.VK_J) {
                Snake.setHeadDirection("down");
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

/**
 * Coordinate class
 * 坐标类
 */
class Coordinate {
    public int x = 0;
    public int y = 0;

    Coordinate(int ox ,int oy) {
        x = ox;
        y = oy;
    }

    Coordinate(int o) {
        this(o, o);
    }

    // Coordinate(Coordinate c) {
    //     this(c.x, c.y);
    // }
}