package slider;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import slider.SliderFrame;

public class SliderTest {
    public static void main(String[] args) {
        JFrame frame = new SliderFrame();
        frame.setTitle("slider test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}