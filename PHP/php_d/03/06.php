<?php

# 委托
# 一个类将一个任务传递给另一个类时，这就是委托。

# 使用继承时，每一个子类是另一个类或多个类的一部分（IS-A关系）;
# 而采用组合，对象可以使用一个不同的类或一组类完成一系列任务（USE-A关系），
# 这并不是说不能使用继承。实际上，大多数设计模式同时包含有继承和组合。

/*
	要避免使用继承形成一长串子类、孙子类、曾孙子类等，设计模式方法建议使用
	浅继承，另外尽量使用多个类的功能。这种方法有助于避免紧密绑定，另外倘若
	一个具体类有子类，修改这个类设计可能导致程序崩溃，而浅继承可以避免这种
	情况。	
  */


/*
 * 设计模式的作用可以分为3大类：
 * 
 * 创建型
 * 结构型
 * 行为型
 *
 * 按范围可以分为两大类：
 *
 * 类
 * 对象
 */

# 创建型模式
# 创建型模式就是用来创建对象的模式。更确切地讲，这些模式是对实例化过程的抽象。


# 结构型模式
# 结构型模式采用继承来组合接口和实现。结构型对象模式则描述了组合对象来建立
# 新功能的方法。


# 行为型模式
# 到目前为止，绝大多数模式都是行为型对象。这些模式的核心是算法和对象之间职责
# 的分配。
