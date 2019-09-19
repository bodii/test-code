package button;

import java.awt.*;
import javax.swing.*;

/**
 * A frame with a button panel
 */
public class ButtonFrame1 extends JFrame {
	private JPanel buttonPanel;
	private static final int DEFAULT_WIDTH = 300;
	private static final int DEFAULT_HEIGHT = 200;

	public ButtonFrame1() {
		setSize(DEFAULT_WIDTH, DEFAULT_HEIGHT);

		makeButton("yellow", Color.YELLOW);
		makeButton("blue", Color.BLUE);
		makeButton("red", Color.RED);

		// JButton exitButton = new JButton("exit");
		// buttonPanel.add(exitButton);
		// exitButton.addActionListener(event -> System.exit(0));

		//pack();
	}

	/**
	 * An action listener that sets the panel's background color.
	 */
	public void makeButton(String name, Color backgroundColor) {
		JButton button = new JButton(name);
		buttonPanel.add(button);
		button.addActionListener(
			event -> buttonPanel.setBackground(backgroundColor)
		);
	}
}
