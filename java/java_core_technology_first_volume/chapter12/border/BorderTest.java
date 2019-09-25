package border;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import border.BorderFrame;

public class BorderTest {
    public static void main(String[] args) {
        JFrame frame = new BorderFrame();
        frame.setTitle("border test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}