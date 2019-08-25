package innerClass;

import java.awt.*;
import java.awt.event.*;
import java.util.*;
import javax.swing.*;
import javax.swing.Timer;


/**
 * This program demonstrates the use of inner classes.
 * @version 1.1 2019-08
 * @author wong
 */

 public class InnerClassTest {
     public static void main(String[] args) {
         TalkingClock clock = new TalkingClock(1000, true);
         clock.start();

         JOptionPane.showMessageDialog(null, "Quit rogram?");
         System.exit(0);
     }
 }

 /**
  * A clock that prints the time in regular intervals.
  */
 class TalkingClock {
    private int interval;
    private boolean beep;

    public TalkingClock(int interval, boolean beep) {
        this.interval = interval;
        this.beep = beep;
    }

    public void start(){
        ActionListener listener = new TimePrinter();
        Timer t = new Timer(interval, listener);
        t.start();
    }

    public class TimePrinter implements ActionListener {
        public void actionPerformed(ActionEvent event) {
            System.out.println("At the tone, the time is " + new Date());
            if (beep) Toolkit.getDefaultToolkit().beep();
        }
    }
 }