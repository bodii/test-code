## final的含义

### final类
如果在类的声明中加上final，则表示该类无法扩展。也就是说，无法创建final类的子类。由于无法创建final类的子类，所以final类中声明的方法也就不会被重写。

### final方法 
如果在实例方法的声明中加上final,则表示该方法不会被子类的方法重写。如果在静态方法的声明中加上final，则表示该方法不会被子类的方法隐藏。如果试图重写或隐藏final方法，编译时会提示错误。

### final字段
final字段只能赋值一次
> 对final实例字段赋值的方法有如下两种:
```java
// 1. 声明时赋上初始值
class Something {
    final int value = 123;
}

// 2. 在构造函数中对字段赋值(blank final)
class Something {
    final int value;
    Something() {
        this.value = 123;
    }
}
```
> 对final静态字段赋值的方法有如下两种:
```java
// 1. 在字段声明时赋上初始值
class Something {
    static final int value = 123;
}

// 2. 在static代码块（静态初始化代码块）中对字段赋值(blank final)。
class Something {
    static final int value;
    static {
        value = 123;
    }
}
```
> final字段不可以由setValue这样的setter方法再次赋值