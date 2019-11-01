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
public class MPanel extends JPanel {
    // 蛇
    private Snake snake;
    // 食物
    private Food food;

    // 开始前的提示信息
    public boolean isStarted = false;

    public MPanel() {

        // 获取焦点
        this.setFocusable(true);
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
        g.setColor(GameArea.titleBackground);
        g.setFont(GameArea.titleFont);
        g.drawString(GameArea.titleText, GameArea.titleX, GameArea.titleY);

        // set body
        g.setColor(GameArea.bodyBackground);
        g.fillRect(GameArea.bodyX, GameArea.bodyY, GameArea.bodyWdith, GameArea.bodyHeight); 

        //
        if (!isStarted) {
            // set start message
            g.setColor(Color.WHITE);
            g.setFont(new Font("arial", Font.BOLD, 25));
            g.drawString("Press Space to Start", 350, 350);
        }
        
        // 初始一条蛇
        snake = new Snake(this, g);
        snake.init();

        // 初始化一个食物
        // food = new Food(this, g);
        // food.init();

        // 添加键盘监听事件
        addKeyListener(new SnakeKeyEvent(this, snake));
    }

}

/**
 * GameArea class 
 * 公共游戏区域类
 */
class GameArea {
    public static final int latticeSize = 25; // 每个格子的大小

    // 标题字体区域
    public static final Font titleFont = new Font("arial", Font.BOLD, 25); // 标题字体
    public static final Color titleBackground = Color.WHITE;
    public static final String titleText = "snake";
    public static final int titleX = 420;
    public static final int titleY = 45;

    // 游戏主区域
    public static final Color bodyBackground  = Color.darkGray;
    public static final int bodyX  = 25;
    public static final int bodyY  = 70;
    public static final int bodyWdith  = 850;
    public static final int bodyHeight  = 600;

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
    private MPanel p;
    private Graphics g;

    // 头部方向
    private ImageIcon headDirection = right;
    // 蛇头坐标
    private Coordinate headCoordinate;
    // 默认蛇的身体长度
    private int bodyLen = 2;
    // 蛇身坐标
    private ArrayList<Coordinate> bodyCoordinate = new ArrayList<>(); 

    Snake(MPanel mp, Graphics gp) {
        p = mp;
        g = gp;
    }

    public void setBodyLen(int l) {
        bodyLen = l;
    }

    public void addBody(int l, Coordinate[] cs) {
        bodyLen += l;
        for (int i = 0 ; i < l; i++) {
            setBody(cs[i]);
        }
    }

    public int getBodyLen() {
        return bodyLen;
    }

    /**
     * 设置蛇头的位置和方向
     * @param d
     * @return
     */
    public void setHead(Coordinate c) {
        headDirection.paintIcon(p, g, c.x, c.y);
    }

    /**
     * 设置蛇头的位置
     * 
     * @param d
     * @return
     */
    public void setHeadDirection(String d) {
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
        
        setHead(getHeadCoordinate());
    }

    /**
     * 获取头的方向
     * 
     * @return
     */
    public ImageIcon getHeadDirection() {
        return headDirection;
    }

    public Coordinate getHeadCoordinate() {
        return headCoordinate;
    }

    public ImageIcon getBody() {
        return body;
    }

    public void setBody(Coordinate c) {
        body.paintIcon(p, g, c.x, c.y);
    }

    public void init() {
        bodyLen = 2;
        int headCoordinateInt = GameArea.latticeSize *4;
        headCoordinate = new Coordinate(headCoordinateInt);
        setHead(headCoordinate);
        for (int i = 0; i < bodyLen; i++) {
            bodyCoordinate.add(
                new Coordinate(
                    headCoordinateInt - GameArea.latticeSize * (i + 1), 
                    headCoordinateInt
                )
            );
            setBody(bodyCoordinate.get(i));
        }
    }
}

/**
 * Food class 
 * 食物类
 */
class Food {
    // 食物
    private ImageIcon food = new ImageIcon("food.png");
    private Coordinate coordinate;
    private MPanel p;
    private Graphics g;

    public Food(MPanel panel, Graphics grap) {
        p = panel;
        g = grap;
    }

    /**
     * @param coordinate the coordinate to set
     */
    public void setCoordinate(Coordinate c) {
        coordinate = c;
    }

    public Coordinate getCoordinate() {
        return coordinate;
    }

    public void hide() {
        coordinate = new Coordinate(0);
    }

    public void setFood(Coordinate coordinate) {
        food.paintIcon(p, g, coordinate.x, coordinate.y);
    }

    public void randFood() {
        Random rand = new Random();
        int cx = rand.nextInt(GameArea.bodyWdith / GameArea.latticeSize) * GameArea.latticeSize + GameArea.bodyX;
        int cy = rand.nextInt(GameArea.bodyHeight / GameArea.latticeSize) * GameArea.latticeSize + GameArea.bodyY;
        setFood(new Coordinate(cx, cy));
    }

    public void unset() {
        food = null;
    }

    public void init() {
        randFood();
    }
}

/**
 * SnakeKeyEvent class
 * 获取监听键盘事件
 */
class SnakeKeyEvent implements KeyListener {
    private MPanel p;
    private Snake snake;

    public SnakeKeyEvent(MPanel mp, Snake s) {
        p = mp;
        snake = s;
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
            p.isStarted = !p.isStarted; // 设置开始的值取反
            p.repaint(); // 重绘内容
        }
        
        if (p.isStarted) {
            // 如果是点中的左键或H键
            if (keyCode == KeyEvent.VK_LEFT || keyCode == KeyEvent.VK_H) {
                snake.setHeadDirection("left");
            } else if (keyCode == KeyEvent.VK_RIGHT || keyCode == KeyEvent.VK_L) {
                snake.setHeadDirection("right");
            } else if (keyCode == KeyEvent.VK_UP || keyCode == KeyEvent.VK_K) {
                snake.setHeadDirection("up");
            } else if (keyCode == KeyEvent.VK_DOWN || keyCode == KeyEvent.VK_J) {
                snake.setHeadDirection("down");
            }
            System.out.println(snake.getHeadCoordinate().x);
            p.repaint(); // 重绘内容
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

    public Coordinate(int ox ,int oy) {
        // 设置边界minX
        if ((ox - GameArea.bodyX) < 0 ) {
            ox = GameArea.bodyX;
        }
        // 设置边界minY
        if ((oy - GameArea.bodyY) < 0) {
            oy = GameArea.bodyY;
        }
        // 设置边界maxX
        if ((ox - (GameArea.bodyX + GameArea.bodyWdith - GameArea.latticeSize)) > 0) {
            ox = GameArea.bodyX + GameArea.bodyWdith - GameArea.latticeSize;
        }
        // 设置边界maxY
        if ((oy - (GameArea.bodyY + GameArea.bodyHeight - GameArea.latticeSize)) > 0) {
            oy = GameArea.bodyY + GameArea.bodyHeight - GameArea.latticeSize;
        }

        x = ox;
        y = oy;
    }

    public Coordinate(int o) {
        this(o, o);
    }

    public Coordinate() {
    }
}