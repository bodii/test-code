package mouse;

import javax.swing.*;
import mouse.MouseComponent;

/**
 * A frame containing a panel for testing mouse operations
 */
public class MouseFrame extends JFrame {
	public MouseFrame() {
		add(new MouseComponent());
		pack();
	}
}
