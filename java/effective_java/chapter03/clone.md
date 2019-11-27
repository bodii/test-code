#### 谨慎地覆盖clone
Cloneable 接口的目的是作为对象的一个mixin接口(mixin interface),表明这样的对象允许克隆(clone)。
遗憾的是，它并不没有成功地达到这个目的。它的主要缺陷在于缺少一个clone方法，而Object的clone方法
是受保护的。如果不借助于反射(reflection)，就不能仅仅因为一个对象实现了Cloneable，就调用clone方法，
即使是反射调用也可能会失败，因为不能保证该对象一定具有可访问的clone方法。

如果一个类实现了Cloneable, Object的clone方法就返回该对象的逐域拷贝，否则就会抛出CloneNotSuppotedException异常。这是接口的一种极端非典型的用法，也不值得仿效。通常情况下，实现接口
是为了表明类可以为它的客户做些什么。然而，对于Cloneable接口，它改变了超类中受保护的方法和行为。

clone方法的通用约定是非常弱的，下面是来自Object规范中的约定内容:
创建和返回该对象的一个拷贝。这个“拷贝”的精确含义取决于该对象的类。一般的含义是，对于任何对象x，表达
式
```java
x.clone() != x
```
将返回true，并且表达式
```java
x.clone().getClass() == x.getClass()
```
将返回结果true,但这些都不是绝对的要求。
```java
x.clone().equals(x)
```
将返回结果true, 但是，这也不是一个绝对的要求。
按照约定，这个方法返回的对象应该通过调用super.clone获得。如果类及其超类(Object除外)遵守这一约定，那么:
```java
x.clone().getClass() == x.getClass()
```

> 不可变的类永远都不应该提供clone方法，因为它只会激发不必要的克隆。
```java
// Clone method for class with no references to mutable state
@Override public PhoneNumber clone() {
    try {
        return (PhoneNumber) super.clone();
    } catch(CloneNotSupportedException e) {
        throw new AssertionError(); // Can't happen
    }
}
```
为了让这个方法生效，应该修改PhoneNumber的类声明为实现Cloneable接口。
虽然Object的clone方法返回的是Object，但这个clone方法返回的却是PhoneNumber。
Java支持协变返回类型(covariant return type)。

如果对象中包含的域引用了可变的对象，使用简单的clone实现可能会导致灾难性的后果。

> 实际上，clone方法就是另一个构造器；必须确保它不会伤害到原始的对象，并确保正确地创建被克隆对象中的
约束条件(invariant)。