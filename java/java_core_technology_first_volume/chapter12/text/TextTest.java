package text;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import text.TextComponentFrame;

public class TextTest {
    public static void main(String[] args) {
        JFrame textFrame = new TextComponentFrame();
        textFrame.setTitle("TextTest");
        textFrame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        textFrame.setVisible(true);
    }
}