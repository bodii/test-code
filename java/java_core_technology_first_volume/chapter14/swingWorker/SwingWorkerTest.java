package swingWorker;

import java.awt.*;
import java.io.*;
import java.util.*;
import java.util.List;
import java.util.concurrent.*;

import javax.swing.*;

/**
 * This program demonstrates a worker thread that runs a potentially time-consuming task.
 * @version 1.1 2019-10
 * @author wong
 */
public class SwingWorkerTest {
    public static void main(String[] args) throws Exception {
        EventQueue.invokeLater(
            () -> {
                JFrame frame = new SwingWorkerFrame();
                frame.setTitle("Swing worker test");
                frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
                frame.setVisible(true);
            }
        );
    }
}

class SwingWorkerFrame extends JFrame {
    private JFileChooser chooser;
    private JTextArea textArea;
    private JLabel statusLine;
    private JMenuItem openItem;
    private JMenuItem cancelItem;
    private SwingWorker<StringBuilder, ProgressData> textReader;
    public static final int TEXT_ROWS = 20;
    public static final int TEXT_COLUMNS = 60;

    public SwingWorkerFrame() {
        chooser = new JFileChooser();
        chooser.setCurrentDirectory(new File("."));

        textArea = new JTextArea(TEXT_ROWS, TEXT_COLUMNS);
        add(new JScrollPane(textArea));

        statusLine = new JLabel(" ");
        add(statusLine, BorderLayout.SOUTH);

        JMenuBar menuBar = new JMenuBar();
        setJMenuBar(menuBar);

        JMenu menu = new JMenu("File");
        menuBar.add(menu);

        openItem = new JMenuItem("Open");
        menu.add(openItem);
        openItem.addActionListener(event -> {
            int result = chooser.showOpenDialog(null);

            if (result == JFileChooser.APPROVE_OPTION) {
                textArea.setText("");
                openItem.setEnabled(false);
                textReader = new TextReader(chooser.getSelectedFile());
                textReader.execute();
                cancelItem.setEnabled(true);
            }
        });

        cancelItem = new JMenuItem("Cancel");
        menu.add(cancelItem);
        cancelItem.setEnabled(false);
        cancelItem.addActionListener(event -> textReader.cancel(true));

        pack();
    }

    private class ProgressData {
        public int number;
        public String line;
    }

    private class TextReader extends SwingWorker<StringBuilder, ProgressData> {
        private File file;
        private StringBuilder text = new StringBuilder();

        public TextReader(File file) {
            this.file = file;
        }

        @Override
        public StringBuilder doInBackground() throws IOException, InterruptedException {
            int lineNumber = 0;
            try (Scanner in = new Scanner(new FileInputStream(file), "UTF-8")) {
                while (in.hasNextLine()) {
                    String line = in.nextLine();
                    lineNumber++;
                    text.append(line).append("\n");
                    ProgressData data = new ProgressData();
                    data.number = lineNumber;
                    data.line = line;
                    publish(data);
                    Thread.sleep(1);
                }
            }
            return text;
        }

        @Override
        public void process(List<ProgressData> data) {
            if (isCancelled()) return;
            StringBuilder b = new StringBuilder();
            statusLine.setText("" + data.get(data.size() -1).number);
            for (ProgressData d: data) b.append(d.line).append("\n");
            textArea.append(b.toString());
        }

        @Override
        public void done() {
            try {
                StringBuilder result = get();
                textArea.setText(result.toString());
                statusLine.setText("Done");
            } catch (InterruptedException ex) {

            } catch(CancellationException ex) {
                textArea.setText("");
                statusLine.setText("Cancelled");
            } catch (ExecutionException ex) {
                statusLine.setText("" + ex.getCause());
            }

            cancelItem.setEnabled(false);
            openItem.setEnabled(true);
        }
    }


}