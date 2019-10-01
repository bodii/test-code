package fileChooser;

import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

import fileChooser.ImageViewerFrame;

public class ImageViewTest {
    public static void main(String[] args) {
        JFrame frame = new  ImageViewerFrame();
        frame.setTitle("Image view");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);
    }
}