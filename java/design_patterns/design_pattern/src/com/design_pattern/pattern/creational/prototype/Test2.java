package com.design_pattern.pattern.creational.prototype;


public class Test2 {
    public static void main(String[] args) throws CloneNotSupportedException {
        Mail2 mail = new Mail2();
        mail.setContent("原始邮件内容");

        for (int i = 0; i < 10; i++) {
            Mail2 mail2 = (Mail2) mail.clone();
            mail2.setName("姓名: " + i);
            mail2.setMailAdderss("地址: " + i + "_mail@163.com");
            mail2.setContent("发送测试内容");
            MailUtil2.sendMail(mail2);
        }

        MailUtil2.saveOrginMailContent(mail);
    }
}