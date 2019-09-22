package calculator;

import javax.swing.*;
import calculator.*;

public class CalculatorJframe extends JFrame {
    public CalculatorJframe() {
        add (new CalculatorPanel());
        pack();
    }
}