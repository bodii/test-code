```java
// Typical hashcode method
@Override public int hashCode() {
    int result = Short.hashCode(areaCode);
    result = 31 * result + Short.hashCode(prefix);
    result = 31 * result + Short.hashCode(lineNum);
    return result;
}
```
虽然本条目中前面给出的hashCode实现方法能够获得相当好的散列函数，但它们并不是最先进的。
它们的质量堪比Java平台类库的值类型中提供的散列函数，这些方法对于绝大多数应用程序而言已经
足够了。

```java
// One-line hashCode method - mdiocre performance
@Override public int hashCode() {
    return Objects.hash(lineNum, prefix, areaCode);
}
```

> 如果一个类是不可变的，并且计算散列码的开销也比较大，就应该考虑把散列码缓存在对象内部，而
> 不是每次请求的时候都重新计算散列码。如果你觉得这种类型的大多数对象会被用作`散列键(hash keys)`，
> 就应该在创建实例的时候计算散列码。否则，可以选择"延迟初始化"(lazily initialize) 散列码，即一直到
> hashCode被第一次调用的时候才初始化。
```java
// hashCode method with lazily initialized cached hash code
private int hashCode;  // Automatically initialized to 0

@Override public int hashCode() {
    int result = hashCode;
    if (result == 0) {
        result = Short.hashCode(areaCode);
        result = 31 * result + Short.hashCode(prefix);
        result = 31 * result + Short.hashCode(lineNum);
        hashCode = result;
    }
    return result;
}
```

> 不要试图从散列码计算中排除一个对象的关键域来提高性能。

> 不要对hashCode方法的返回值做出具体的规定，因此客户端无法理所当然地信赖它；这样可以为修改提供
> 灵活性。