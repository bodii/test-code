package com.design_pattern.pattern.behavioral.memento;

import java.util.Stack;

public class ArticleMementoManager {
    private final Stack<ArticleMemento> ARTICLIE_MEMENTS = new Stack<>();

    public ArticleMemento getArticleMemento() {
        return ARTICLIE_MEMENTS.pop();
    }

    public void addArticleMemento(ArticleMemento articleMemento) {
        ARTICLIE_MEMENTS.push(articleMemento);
    }
}
