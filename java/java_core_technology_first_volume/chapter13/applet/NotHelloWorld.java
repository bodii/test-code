package applet;

import java.awt.*;
import javax.swing.*;

/**
 * @version 1.0 2019-10
 * @author wong
 */
public class NotHelloWorld extends JApplet {
    public void init() {
        EventQueue.invokeLater(
            () -> {
                Label label = new JLabel("Not a Hello, World applet", SwingConstants.CENTER);
                add(label);
            }
        );
    }
}