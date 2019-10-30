package snake;

import java.awt.EventQueue;

import javax.swing.JFrame;

import snake.SnakeGameFrame;

public class SnakeGame {
    public static void main(String[] args) {
        EventQueue.invokeLater(
            () -> {
                JFrame frame = new SnakeGameFrame();
                
            }
        );
    }
}