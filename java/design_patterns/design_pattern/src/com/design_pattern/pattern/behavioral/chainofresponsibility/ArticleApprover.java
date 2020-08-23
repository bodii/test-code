package com.design_pattern.pattern.behavioral.chainofresponsibility;

public class ArticleApprover extends Approver {
    @Override
    public void deploy(Course course) {
        if (!(course.getArticle() == null || "".equals(course.getArticle()))) {
            System.out.println(course.getName() + "的手记不为空，批准！");
            if (approver != null) {
                approver.deploy(course);
            }
        } else {
            System.out.println("手记为空，不批准，流程结束！");
            return;
        }
    }
}
