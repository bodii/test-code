package button;

import java.awt.*;
import javax.swing.*;
import button.ButtonFrame1;

public class ButtonTest1 {
	public static void main(String[] args) {
		EventQueue.invokeLater(
			() -> {
				JFrame frame = new ButtonFrame1();
				frame.setTitle("ButtonColor");
				frame.setDefaultCloseOperation(frame.EXIT_ON_CLOSE);
				frame.setVisible(true);
			}
		);
	}
}

