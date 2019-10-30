package snake;

import java.awt.Color;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Point;
import java.awt.Toolkit;
import java.awt.BorderLayout;

import javax.swing.JButton;
import javax.swing.JFrame;

import snake.SnakeGameHeadPanel;
import snake.SnakeGameBodyPanel;


/**
 * 贪吃蛇游戏主窗口
 */
public class SnakeGameFrame extends JFrame {
    public static final int DEFAULT_WIDTH = 900;
    public static final int DEFAULT_HEIGHT = 720;

    public SnakeGameFrame() {
        setTitle("snake game");
        // 设置窗口默认关闭按钮
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        // 设置不可拖拽大小
        setResizable(false);     

        // 设置窗口大小
        setSize(new Dimension(DEFAULT_WIDTH, DEFAULT_HEIGHT));
        // 获取屏幕尺寸
        Toolkit kit = Toolkit.getDefaultToolkit();
        Dimension screenSize = kit.getScreenSize();
        int screenWidth = screenSize.width;
        int screenHeight  = screenSize.height;
        // 设置窗口距离左上角屏幕的距离(中间)
        setLocation(
            new Point((screenWidth - DEFAULT_WIDTH) / 2, (screenHeight - DEFAULT_HEIGHT) / 2)
        );
        //设置背景为白色
        setBackground(Color.WHITE);
        // 设置窗口布局,并设置距离窗口的左右间距
        setLayout(new BorderLayout());
        // 添加主窗口内 － 头部组件
        add(new SnakeGameHeadPanel(), BorderLayout.NORTH); 
        // 添加主窗口内 － 游戏域
        // add(new SnakeGameBodyPanel(), BorderLayout.SOUTH);

        // 是否显示
        setVisible(true);
        
    }
}