package com.design_pattern.pattern.behavioral.chainofresponsibility;

public abstract class Approver {
    protected Approver approver;

    protected void setNextApprover(Approver approver) {
        this.approver = approver;
    }

    public abstract void deploy(Course course);
}
