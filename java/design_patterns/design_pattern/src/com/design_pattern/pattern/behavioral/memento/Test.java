package com.design_pattern.pattern.behavioral.memento;

public class Test {
    public static void main(String[] args) {
        ArticleMementoManager articleMementoManager = new ArticleMementoManager();

        Article article = new Article("文章标题测试","content", "imgs_url");
        ArticleMemento articleMemento  = article.saveToMemento();
        articleMementoManager.addArticleMemento(articleMemento);
        System.out.println("文章标题:" + article.getTitle() + " 文章内容：" + article.getContent() + " 图片: " + article.getImgs());
        System.out.println("文章完整内容： " + article);

        System.out.println("修改文章start");
        article.setTitle("新的标题");
        article.setContent("新的文章内容");
        article.setImgs("新的文章图片");
        System.out.println("修改文章end");
        System.out.println("文章完整内容:" + article);
        articleMemento = article.saveToMemento();
        articleMementoManager.addArticleMemento(articleMemento);

        System.out.println("修改文章start");
        article.setTitle("新的标题02");
        article.setContent("新的内容02");
        article.setImgs("新的图片02");
        System.out.println("修改文章end");
        System.out.println("文章完整内容:" + article);

        System.out.println("暂存回退");
        System.out.println("回退出栈1次");
        articleMemento = articleMementoManager.getArticleMemento();
        article.undoFromMemento(articleMemento);

        System.out.println("回退出栈2次");
        articleMemento = articleMementoManager.getArticleMemento();
        article.undoFromMemento(articleMemento);
        
        System.out.println("暂存回退");
        System.out.println("文章的完整内容" + article);
    }
}
