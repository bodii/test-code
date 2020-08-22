package com.design_pattern.pattern.behavioral.memento;

public class Test {
    public static void main(String[] args) {
        ArticleMementoManager articleMementoManager = new ArticleMementoManager();
        Article article = new Article("文章文章标题测试","content", "imgs_url");

        ArticleMemento articleMemento  = article.saveToMemento();
        articleMementoManager.addArticleMemento(articleMemento);
        System.out.println("文章标题:" + article.getTitle() + " 文章内容：" + article.getContent() + " 图片: " + article.getImgs());

        System.out.println("文章完整内容： " + article);
    }
}
