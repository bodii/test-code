package swing;

import java.awt.*;
import java.util.*;

import javax.swing.*;

/**
 * This program demonstrates that a thread that runs in parallel with the event
 * dispatch thread can cause errors in Swing components.
 * @version 1.0 2019-10
 * @author wong
  */
public class SwingThreadTest {
    public static void main(String[] args) {
        EventQueue.invokeLater(
            () -> {
                JFrame frame = new SwingThreadFrame();
                frame.setTitle("SwingThreadTest");
                frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
                frame.setVisible(true);
            }
        );
    }
}


/**
 * This frame has two buttons to fill a combo box from a separate therad.
 * The "Good" button uses the event queue, 
 * The "Bad" button modifies the combo box directly.
 */
class SwingThreadFrame extends JFrame {
    public SwingThreadFrame() {
        final JComboBox<Integer> combo = new JComboBox<>();
        combo.insertItemAt(Integer.MAX_VALUE, 0);
        combo.setPrototypeDisplayValue(combo.getItemAt(0));
        combo.setSelectedIndex(0);

        JPanel panel = new JPanel();

        JButton goodButton = new JButton("Good");
        goodButton.addActionListener(
            event -> new Thread(new GoodWorkerRunnable(combo)).start()
        );
        panel.add(goodButton);
        JButton badButton = new JButton("Bad");
        badButton.addActionListener(
            event -> new Thread(new BadWorkerRunnable(combo)).start()
        );
        panel.add(badButton);

        panel.add(combo);
        add(panel);
        pack();
    }
}

class BadWorkerRunnable implements Runnable {
    private JComboBox<Integer> combo;
    private Random generator;

    public BadWorkerRunnable(JComboBox<Integer> aCombo) {
        combo = aCombo;
        generator = new Random();
    }

    public void run() {
        try {
            while (true) {
                int i = Math.abs(generator.nextInt());
                if (i % 2 == 0)
                    combo.insertItemAt(i, 0);
                else if (combo.getItemCount() > 0)
                    combo.removeItemAt(i % combo.getItemCount());
                Thread.sleep(1);
            }
        } catch (InterruptedException e) {
            
        }
    }
}


class GoodWorkerRunnable implements Runnable {
    private JComboBox<Integer> combo;
    private Random generator;

    public GoodWorkerRunnable(JComboBox<Integer> aCombo) {
        combo = aCombo;
        generator = new Random();
    }

    public void run() {
        try {
            while (true) {
                EventQueue.invokeLater(
                    () -> {
                        int i = Math.abs(generator.nextInt());
                        if (i % 2 == 0)
                            combo.insertItemAt(i, 0);
                        else if (combo.getItemCount() > 0)
                            combo.removeItemAt(i % combo.getItemCount());
                    }
                );
                Thread.sleep(1);
            }
        } catch(InterruptedException e) {

        }
    }
}