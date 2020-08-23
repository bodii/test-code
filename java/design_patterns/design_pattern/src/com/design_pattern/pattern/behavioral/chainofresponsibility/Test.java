package com.design_pattern.pattern.behavioral.chainofresponsibility;

public class Test {
    public static void main(String[] args) {
        Approver articleApprover = new ArticleApprover();
        Approver videoApprover = new VideoApprover();

        Course course = new Course();
        course.setName("Java设计精讲");
        course.setArticle("文章内容");
//        course.setVideo("视频内容");
        articleApprover.setNextApprover(videoApprover);
        articleApprover.deploy(course);
    }
}
