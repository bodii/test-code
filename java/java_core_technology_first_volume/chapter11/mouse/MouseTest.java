package mouse;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

import mouse.MouseFrame;

public class MouseTest {
    public static void main(String[] args) {
        MouseFrame frame = new MouseFrame();
        frame.setTitle("Mouse Test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}