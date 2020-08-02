package com.design_pattern.pattern.behavioral.templatemethod;

public class FECourse extends ACourse {
    private boolean writeArticleFlag;

    public FECourse(boolean writeArticleFlag) {
        this.writeArticleFlag = writeArticleFlag;
    }

    public FECourse() {
        this(false);
    }

    @Override
    protected boolean needWriteArticle() {
        return writeArticleFlag;
    }

    @Override
    public void packageCourse() {
        System.out.println("提供前端源码");
        System.out.println("提供课程的图片等多媒体素材");
    }
}
