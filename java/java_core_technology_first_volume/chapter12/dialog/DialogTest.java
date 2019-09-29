package dialog;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import dialog.DialogFrame;

public class DialogTest extends DialogFrame {
    public static void main(String[] args) {
        JFrame frame = new DialogFrame();
        frame.setTitle("dialog test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}