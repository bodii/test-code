<?php

/*

    [ 创建型设计模式杂谈 ]

    单例模式，工厂模式，建造者模式，原型模式都属于创建型模式。使用创建型模式的目的，
    就是为了创建一个对象。


    创建型模式的优点，在于如何把复杂的创建过程封装起来，如何降低系统的内销。

    我认为创建型模式的一个重要的思想其实就是封装，利用封装，把直接获得一个对象改为
    通过一个接口获得一个对象。这样最明显的优点，在于我们可以把一些复杂的操作也封装
    到接口里去，我们使用时直接调这个接口就可以了。具体的实现，我们在主线程序中就不
    再考虑。这样使得我们的代码看上去更少，更简洁。

    单例模式，我们把对象的生成从new改为通过一个静态方法，通过静态方法的控制，使得
    我们总是返回同一个实例给调用者，确保了系统只有一个实例

    工厂模式，也是一样，生成对象改为接口，还可以通过传参实例化不同的类。如果我们通
    过直接new的话，那么我们 在主线代码中少不了要写if condition new 一个加法类，
    else new一个减法类。封装了之后，我们通过接口传参，还能利用多态的特性去替代if
     else语句。
    而且我们遵循了单一原则，让类的功能单一。我们如果需要一个新功能，只需添加一个类，
    不用修改其他类的功能。这样使得代码的扩展性更好了。

    建造者模式，我们把初始化的工作和顺序，封装给了一个建造者和指挥者。如果，我们下
    次要建造的类属性，或是顺序 不同。我们只需新建对应的建造者类或添加对应的指挥者
    法，不必再去修改原代码。而且我们也省去了，这new对象后，还要 写$attribut=array();
    这种一大串数组，然后调好几个方法去初始化的工作。

    原型模式，通过先创建一个原型对象，然后直接克隆，省去了new大对象带来的开销浪费。
    当然我们同样可以通过，封装clone这个动作。使得我们在clone的同时还可以做一些其他
    的准备工作。

*/
