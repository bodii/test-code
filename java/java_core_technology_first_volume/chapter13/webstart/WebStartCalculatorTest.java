package webstart;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

import webstart.WebStartCalculatorFrame;

public class WebStartCalculatorTest {
    public static void main(String[] args) {
        EventQueue.invokeLater(
            () -> {
                JFrame frame = new WebStartCalculatorFrame();
                frame.setTitle("web start calculator test");
                frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
                frame.setVisible(true);
            }
        );
    }
}

