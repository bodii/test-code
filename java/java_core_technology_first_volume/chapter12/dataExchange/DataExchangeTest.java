package dataExchange;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
import dataExchange.DataExchangeFrame;

public class DataExchangeTest {
    public static void main(String[] args) {
            JFrame frame = new DataExchangeFrame();
            frame.setTitle("data exchange test");
            frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
            frame.setVisible(true);
    }
}