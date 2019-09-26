package toolBar;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import toolBar.ToolBarFrame;

public class ToolBarTest {
    public static void main(String[] args) {
        JFrame frame = new ToolBarFrame();
        frame.setTitle("tool bar test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}