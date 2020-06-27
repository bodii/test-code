package chapter08;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;

import java.awt.event.ActionListener;
import java.awt.FlowLayout;
import java.awt.event.ActionEvent;

public class MyFrame extends JFrame implements ActionListener {
    private final JLabel label = new JLabel("Event Dispatching Thread Sample");
    private final JButton button = new JButton("countUp");

    public MyFrame() {
        super("MyFrame");
        getContentPane().setLayout(new FlowLayout());
        getContentPane().add(label);
        getContentPane().add(button);
        button.addActionListener(this);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        pack();
        setVisible(true);
    }

    /**
     * actionPerformed()方法及paintComponent()方法都是事件分派Swing线程中被调用，所以
     * 如果在actionPerformed()方法中有大量的耗时操作，将会导致用户界面无法响应用户请求，造
     * 成用户界面锁死。在这种情况下，最好在actionPerformed()方法中将耗时操作的程序部分移至
     * 一个新的Swing线程中实现，这样程序在进行大量计算的同时又可以及时地响应用户请求。
     */
    @Override
    public void actionPerformed(ActionEvent e) {
        if (e.getSource() == button) {
            new Thread(){
                @Override
                public void run() {
                    countUp();
                }
            }.start();
        }
    }

    private void countUp() {
        for (int i = 0; i < 10; i++) {
            System.out.println(
                Thread.currentThread().getName() + ":countUp:setText(" + i + ")"
                );
            label.setText("" + i);

            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {

            }
        }
    }
}