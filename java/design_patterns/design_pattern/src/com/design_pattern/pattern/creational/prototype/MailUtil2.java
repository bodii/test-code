package com.design_pattern.pattern.creational.prototype;

import java.text.MessageFormat;

public class MailUtil2 {
    public static void sendMail(Mail2 mail) {
        String sendContent = "向{0}同学的邮箱地址{1},发送邮件内容是:{2}";
        System.out.println(
            MessageFormat.format(
                sendContent, mail.getName(), mail.getMailAdderss(), mail.getContent()
                )
        );
    }

    public static void saveOrginMailContent(Mail2 mail) {
        System.out.println("保存原始邮件内容: " + mail.getContent());
    }
}