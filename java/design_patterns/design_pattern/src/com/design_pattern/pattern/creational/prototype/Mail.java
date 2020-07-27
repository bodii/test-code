package com.design_pattern.pattern.creational.prototype;

public class Mail {
    private String name;
    private String mailAdderss;
    private String content;

    public Mail() {
        System.out.println("Mail class Constructor");
    }

    public String getName() {
        return name;
    }

    public String getMailAdderss() {
        return mailAdderss;
    }

    public String getContent() {
        return content;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setMailAdderss(String mailAdderss) {
        this.mailAdderss = mailAdderss;
    }

    public void setContent(String content) {
        this.content = content;
    }

    @Override
    public String toString() {
        return "Mail{name: '" + name + "'" +
                " mailAdderss: '" + mailAdderss + "'" +
                " content: '" + content + "'}";
    }
    
}