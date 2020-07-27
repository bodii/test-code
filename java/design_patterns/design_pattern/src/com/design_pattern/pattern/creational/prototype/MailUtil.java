package com.design_pattern.pattern.creational.prototype;

import java.text.MessageFormat;

public class MailUtil {
    public static void sendMail(Mail mail) {
        String sendContent = "向{0}同学,邮件地址:{1},邮件内容:{2},发送邮件成功!";
        System.out.println(
            MessageFormat.format(
                sendContent, mail.getName(), mail.getMailAdderss(), mail.getContent()
                )
            );
    }

    public static void saveOrginMailRecord(Mail mail) {
        System.out.println("存储orgin mail记录, orgin mail内容: " + mail.getContent());
    }

}