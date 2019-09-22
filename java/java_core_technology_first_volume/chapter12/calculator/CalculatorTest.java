package calculator;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import calculator.*;

public class CalculatorTest {
    public static void main(String[] args) {
        JFrame frame = new CalculatorJframe();
        frame.setTitle("Calculator");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}