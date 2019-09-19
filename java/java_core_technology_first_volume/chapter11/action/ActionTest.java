package action;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import action.ActionFrame;


/**
 * @version 1.1 2019-09
 * @author wong
 */
public class ActionTest {
	public static void main(String[] args) {
		EventQueue.invokeLater(
			() -> {
				JFrame frame = new ActionFrame();
				frame.setTitle("ActionTest");
				frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
				frame.setVisible(true);
			}
		);
	}
}
