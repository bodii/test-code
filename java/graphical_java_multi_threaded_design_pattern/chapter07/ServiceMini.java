package chapter07;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;

public class ServiceMini {
    private final int portnumber;
    public ServiceMini(int port) {
        this.portnumber = port;
    }

    public void execute() throws IOException {

        ServerSocket serverSocket = new ServerSocket(portnumber);
        System.out.println("Listening on " + serverSocket);
        try {
            while (true) {
                System.out.println("Accepting...");
                Socket clientSocket = serverSocket.accept();

                System.out.println("Connected to " + clientSocket);
                Service.service(clientSocket);
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            serverSocket.close();
        }
    }
}