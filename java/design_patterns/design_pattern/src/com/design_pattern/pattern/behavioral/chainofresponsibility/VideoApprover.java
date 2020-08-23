package com.design_pattern.pattern.behavioral.chainofresponsibility;

public class VideoApprover extends Approver {
    @Override
    public void deploy(Course course) {
        if (!(course.getVideo() == null || "".equals(course.getVideo()))) {
            System.out.println("视频内容不为空，批准！");
            if (approver != null) {
                approver.deploy(course);
            }
        } else {
            System.out.println("视频内容为空，不批准，流程结束！");
            return;
        }
    }
}
