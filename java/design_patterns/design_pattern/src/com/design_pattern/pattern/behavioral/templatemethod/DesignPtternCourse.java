package com.design_pattern.pattern.behavioral.templatemethod;

public class DesignPtternCourse extends ACourse {
    private boolean writeArticleFlag;

    public DesignPtternCourse() {
        this(false);
    }

    public DesignPtternCourse(boolean writeArticleFlag) {
        this.writeArticleFlag = writeArticleFlag;
    }

    @Override
    protected boolean needWriteArticle() {
        return writeArticleFlag;
    }

    @Override
    public void packageCourse() {
        System.out.println("提供课程Java源代码");
    }
}
