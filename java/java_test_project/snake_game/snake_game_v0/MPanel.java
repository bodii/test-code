package snake;

import java.awt.Graphics;
import java.awt.Color;
import java.awt.Font;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.awt.event.ActionListener;
import java.awt.event.ActionEvent;
import java.util.Random;
import java.util.ArrayList;
import java.util.Random;

import javax.swing.JPanel;
import javax.swing.ImageIcon;
import javax.swing.Timer;



/**
 * MPanel class
 * 贪吃蛇画布
 */
public class MPanel extends JPanel implements KeyListener, ActionListener {

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
    private Timer timer = new Timer(350, this);
    public String deretion = "right";
    public int[] foodCoordinat = new int[2];
    private Random rand = new Random();
    public Boolean isFoodEaten = true;
    public long startTime = System.currentTimeMillis();
    public long endTime = System.currentTimeMillis();



    public MPanel() {
        init();
        setFocusable(true); // 添加焦点
        addKeyListener(this); // 添加键盘事件监听
        timer.start();
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

        // sorce
        g.setFont(new Font("arial", Font.BOLD, 15));
        g.drawString("time: " + (endTime - startTime) / 1000, 750, 25);
        g.drawString("food: " + (bodyLen - 2), 750, 40);
        g.drawString("sorcce: " + ((bodyLen - 2) * latticeSize), 750, 55);

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

        food.paintIcon(this, g, foodCoordinat[0], foodCoordinat[1]);
    }

    public void init() {
        bodyLen = 2;
        snakeHeadCoordinate[0] = 100;
        snakeHeadCoordinate[1] = 100;
        for (int i = 0; i < bodyLen; i++) {
            snakeBodyCoordinateX[i] =  100  -  25 * (i +1);
            snakeBodyCoordinateY[i] = 100;
        }
        isFoodEaten = true;
        getFood();
        startTime = System.currentTimeMillis();
        endTime = System.currentTimeMillis();
    }

    public void getFood() {
        if (isFoodEaten) {
            foodCoordinat[0] = rand.nextInt(850 / 25) * 25 + 25;
            foodCoordinat[1] = rand.nextInt(645 / 25) * 25 + 70;
            isFoodEaten = !isFoodEaten;
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
                deretion = "up";
                repaint();
            } else if (keyCode == KeyEvent.VK_DOWN || keyCode == KeyEvent.VK_J) {
                snakeHeadDerection = snakeHeadDown;
                deretion = "down";
                repaint();
            } else if (keyCode == KeyEvent.VK_LEFT || keyCode == KeyEvent.VK_H) {
                snakeHeadDerection = snakeHeadLeft;
                deretion = "left";
                repaint();
            } else if (keyCode == KeyEvent.VK_RIGHT || keyCode == KeyEvent.VK_L) {
                snakeHeadDerection = snakeHeadRight;
                deretion = "right";
                repaint();
            }
        }
    }

    @Override
    public void keyReleased(KeyEvent e) {

    }

    @Override
    public void actionPerformed(ActionEvent e) {
        if (isStarted) {

            if (deretion == "right") {
                snakeHeadCoordinate[0] += latticeSize;
                if (snakeHeadCoordinate[0] > 850) {
                    snakeHeadCoordinate[0] = latticeSize;
                }
                for(int i = 0; i < bodyLen; i++) {
                    snakeBodyCoordinateX[i] = snakeHeadCoordinate[0]  -  ((i + 1) * latticeSize);
                    snakeBodyCoordinateY[i] = snakeHeadCoordinate[1];
                    if (snakeBodyCoordinateX[i] > 850) {
                        snakeBodyCoordinateX[i] = latticeSize;
                    } 
                    if ( snakeBodyCoordinateX[i] < latticeSize) {
                        snakeBodyCoordinateX[i] = 850;
                    }
                }
            } else if (deretion == "left") {
                snakeHeadCoordinate[0] -= latticeSize;
                if (snakeHeadCoordinate[0]  < latticeSize) {
                    snakeHeadCoordinate[0] = 850;
                }
                for(int i = 0; i < bodyLen; i++) {
                    snakeBodyCoordinateX[i] = snakeHeadCoordinate[0]  + ((i  + 1) * latticeSize);
                    snakeBodyCoordinateY[i] = snakeHeadCoordinate[1];
                    if (snakeBodyCoordinateX[i] < latticeSize) {
                        snakeBodyCoordinateX[i] = 850;
                    }

                    if (snakeBodyCoordinateX[i] > 850) {
                        snakeBodyCoordinateX[i] = latticeSize;
                    }
                }
            }  else if (deretion == "up") {
                snakeHeadCoordinate[1] -= latticeSize;
                if (snakeHeadCoordinate[1] < 70) {
                    snakeHeadCoordinate[1] = 645;
                }
                for(int i = 0; i < bodyLen; i++) {
                    snakeBodyCoordinateX[i] = snakeHeadCoordinate[0];
                    snakeBodyCoordinateY[i] = snakeHeadCoordinate[1] + (i +1) * latticeSize;
                    if (snakeBodyCoordinateY[i] < 70) {
                        snakeBodyCoordinateY[i] = 645;
                    }
                    if (snakeBodyCoordinateY[i] > 645) {
                        snakeBodyCoordinateY[i] = 70;
                    }
                }
            } else {
                snakeHeadCoordinate[1] += latticeSize;
                if (snakeHeadCoordinate[1] > 645) {
                    snakeHeadCoordinate[1] = 70;
                }
                for(int i = 0; i < bodyLen; i++) {
                    snakeBodyCoordinateX[i] = snakeHeadCoordinate[0];
                    snakeBodyCoordinateY[i] = snakeHeadCoordinate[1] - (i +1) * latticeSize;
                    if (snakeBodyCoordinateY[i] > 645) {
                        snakeBodyCoordinateY[i] = 70;
                    }
                    if (snakeBodyCoordinateY[i] < 70) {
                        snakeBodyCoordinateY[i] = 645;
                    }
                }
            }

            // 如果蛇头与食物的x、y坐标相等，重新生成一个食物，并且蛇的身体加1
            if (snakeHeadCoordinate[0] == foodCoordinat[0] && snakeHeadCoordinate[1] == foodCoordinat[1]) {
                isFoodEaten = true;
                bodyLen += 1;
            }

            // 如果食物被吃掉，生成一个食物
            if (isFoodEaten) {
                getFood();
            }
            endTime =System.currentTimeMillis();
            repaint();
        }
        // timer.start();

    }
}