package com.design_pattern.pattern.behavioral.templatemethod;

public abstract class ACourse {
    protected final void makeCourse() {
        makePPT();
        makeVideo();
        if (needWriteArticle())
            writeArticle();
        packageCourse();
    }

    protected final void makePPT() {
        System.out.println("制作PPT");
    }

    protected final void makeVideo() {
        System.out.println("制作视频");
    }

    protected final void writeArticle() {
        System.out.println("编写手记");
    }

    /**
     * 钩子方法
     * @return boolean 是否需要编写手记
     */
    protected boolean needWriteArticle() {
        return false;
    }

    public abstract void packageCourse();

}
