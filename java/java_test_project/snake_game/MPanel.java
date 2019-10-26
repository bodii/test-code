package snake;

import javax.swing.JPanel;
import java.awt.Graphics;
import java.awt.Color;

public class MPanel extends JPanel {
    public MPanel() {

    }

    public void paintComponent(Graphics g) {
        super.paintComponent(g);
        this.setBackground(Color.GREEN);
    }
}