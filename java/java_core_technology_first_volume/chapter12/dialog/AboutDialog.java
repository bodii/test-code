package dialog;

import java.awt.BorderLayout;
import javax.swing.JButton;
import javax.swing.JDialog;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;

/**
 * A sample modal dialog that displays a message and waits for the user to click the OK button.
 */
public class AboutDialog extends JDialog {
    public AboutDialog(JFrame owner) {
        super(owner, "About DialogTest", true);

        add(
            new JLabel(
                "<html><h1><i>Core Java</i><hr> By Cay Horstmann</html>"
            ),
            BorderLayout.CENTER
        );
        JButton ok = new JButton("OK");
        ok.addActionListener(event -> setVisible(false));

        JPanel panel = new JPanel();
        panel.add(ok);
        add(panel, BorderLayout.SOUTH);

        pack();
    }
}