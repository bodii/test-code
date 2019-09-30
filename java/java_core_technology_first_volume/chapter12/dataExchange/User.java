package dataExchange;

public class User {
    private String username;
    private char[] password;

    User(String username, char[] password) {
        this.username = username;
        this.password = password;
    }

    public User getUser() {
        return this;
    }

    public String getName() {
        return username;
    }

    public char[] getPassword() {
        return password;
    }
}