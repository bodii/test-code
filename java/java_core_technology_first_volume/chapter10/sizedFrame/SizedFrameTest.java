package sizedFrame;

import java.awt.*;
import javax.swing.*;

/**
 * @version 1.1 2019-09
 * @author wong
 */
public class SizedFrameTest {
	public static void main(String[] args) {
		EventQueue.invokeLater(
			() -> {
				JFrame frame = new SizedFrame();
				frame.setTitle("SizedFrame");
				frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
			}
		);
	}
}

class SizedFrame extends JFrame {
	public SizedFrame() {
		Toolkit kit = Toolkit.getDefaultToolkit();
		Dimension screenSize = kit.getScreenSize();
		int screenHeight = screenSize.height;
		int screenWidth = screenSize.width;

		setSize(screenWidth / 2, screenWidth / 2);
		setLocationByPlatform(true);

		Image img = new ImageIcon("icon.gif").getImage();
		setIconImage(img);
	}
}
