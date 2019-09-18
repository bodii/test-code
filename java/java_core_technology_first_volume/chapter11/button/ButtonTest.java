package button;

import java.awt.*;
import javax.swing.*;
import button.ButtonFrame;

public class ButtonTest {
	public static void main(String[] args) {
		EventQueue.invokeLater(
			() -> {
				JFrame frame = new ButtonFrame();
				frame.setTitle("ButtonColor");
				frame.setDefaultCloseOperation(frame.EXIT_ON_CLOSE);
				frame.setVisible(true);
			}
		);
	}
}

