package timer;

/**
 * @version 1.1 2018-08
 * @author wong
 */

 import java.awt.*;
 import java.awt.event.*;
 import java.util.*;
 import javax.swing.*;
 import javax.swing.Timer;

 public class TimerTest {
     public static void main(String[] args) {
         ActionListener listener = new TimePrinter();

         Timer t = new Timer(10000, listener);
         t.start();

         JOptionPane.showMessageDialog(null, "Quit program?");
         System.exit(0);
     }
 }

 class TimePrinter implements ActionListener {
    public void actionPerformed(ActionEvent event) {
        System.out.println("At the tone, the time is " + new Date());
        Toolkit.getDefaultToolkit().beep();
    }
 }