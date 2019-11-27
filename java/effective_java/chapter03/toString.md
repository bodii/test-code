#### toString
始终要覆盖toString
虽然Object提供了toString方法的一个实现，但它返回的字符串通常并不是类的用户所期望看到的。
它包含类的名称，以及一个"@"符号，接着是散列码的无符号十六进制表示法，例如PhoneNumber@163b91。
toString的通用约定指出，被返回字符串应该是一个“简洁的但信息丰富，并且易于阅读的表达形式"。尽管有人
认为PhoneNumber@163b91算得上是简洁和易于阅读了，但是与707-867-5309比起来，它还算不上是信息丰富
的。toString约定进一步指出，"建议所有的子类都覆盖这个方法。"

> 提供好的toString实现可以使类用起来更加舒适，使用了这个类的系统当更易于调试。
当对象被传递给println、printf、字符串联操作符(+)以及assert, 或者被调试器打印出来时，toString方法会被自动调用。即使你永远不调用对象的toString方法，但是其他人也许可能需要。例如，带有对象引用的一个组件，在它记录的错误消息中，可能包含该对象的字符串表示法。如果你没有覆盖toString，这条消息可能就毫无用处。

如果为PhoneNumber提供了好的toString方法，那么要产生有用的诊断消息会非常容易:
```java
System.out.println("Failed to connect to " + phoneNumber);
```
> 在实际应用中，toString方法应该返回对象中包含的所有值关注的信息。

```java
@Override public String toString() {
    return String.format("%03d-%03d-%04d", areaCode, prefix, lineNum);
}
```
如果你决定不指定格式，那么文档注释部分也应该有如下所示的指示信息:
```java
/**
* Returns a brief description of this potion. The exact details
* of the representation are unspecified and subject to change,
* but the following may be regarded as typical:
* "[Potion #9: type=love, smell=turpentine, look=india ink]"
*/
@Override public String toString() { ... }
```
> toString 返回值中包含的所有信息提供一种可以通过编程访问之的途径。
