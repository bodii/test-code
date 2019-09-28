package optionDialog;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import optionDialog.OptionDialogFrame;

public class OptionDialogTest {
    public static void main(String[] args) {
        JFrame frame = new OptionDialogFrame();
        frame.setTitle("option dialog");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}