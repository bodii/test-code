package menu;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import menu.MenuFrame;

public class MenuTest {
    public static void main(String[] string) {
        JFrame frame = new MenuFrame();
        frame.setTitle("menu test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}