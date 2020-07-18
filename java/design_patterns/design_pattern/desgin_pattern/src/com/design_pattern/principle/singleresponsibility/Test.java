package com.design_pattern.principle.singleresponsibility;

public class Test {
    public static void main(String[] args) {
        IBird dayan = new DayanBird();
        IBird ostrich = new OstrichBird();

        System.out.println(dayan.name());
        System.out.println(ostrich.name());
    }
}