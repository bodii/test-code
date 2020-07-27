package com.design_pattern.pattern.creational.prototype;

public class Test {
    public static void main(String[] args) {
        Mail mail = new Mail();
        mail.setContent("初始化邮件内容");
        for (int i = 0; i < 10; i++) {
            mail.setName("姓名: " + i);
            mail.setMailAdderss("mail_" + i + "@163.com");
            mail.setContent("恭喜您,这是一个测试邮件!");
            MailUtil.sendMail(mail);
        }
        MailUtil.saveOrginMailRecord(mail);
    }
}