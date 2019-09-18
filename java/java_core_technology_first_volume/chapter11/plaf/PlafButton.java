package plaf;

import java.awt.*;
import javax.swing.*;
import plaf.PlafFrame;


/**
 * @version 1.1 2019-09
 * @author wong
 */
public class PlafButton {
	public static void main(String[] args) {
		EventQueue.invokeLater(
			() -> {
				JFrame frame = new PlafFrame();
				frame.setTitle("PlafButton");
				frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
				frame.setVisible(true);
			}
		);
	}
}
