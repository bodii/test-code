#### 枚举集与映射
EnumSet是一个枚举类型元素集的高效实现。由于枚举类型只有有限个实例，所以EnumSet内部用位
序列实现。如果对应的值在集中，则相应的位被置为1。

EnumSet类没有公共的构造器。可以使用静态工厂方法构造这个集:
enum Weekday { MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY };
EnumSet<Weekday> always = EnumSet.allOf(Weekday.class);
EnumSet<Weekday> never = EnumSet.noneOf(Weekday.class);
EnumSet<Weekday> workday = EnumSet.range(Weekday.MONDAY, Weekday.FRIDAY);
EnumSet<Weekday> mwf = EnumSet.of(Weekday.MonDAY, Weekday.WEDNSDAY, Weekday.FRIDAY);

可以使用Set接口的常用方法来修改EnumSet。
EnumMap是一个键类型为枚举类型的映射。它可以直接且高效地用一个值数组实现。在使用时，
需要在构造器中指定键类型:
```java
EnumMap<Weekday, Employee> personInCharge = new EnumMap<>(Weekday.class);
```

> 在EnumSet的API文档中，将会看到E extends Enum<E> 这样奇怪的类型参数。
> 简单地说，它的意思是"E是一个枚举类型"所有的枚举类型都扩展于泛型Enum类。
例如， Weekday扩展Enum<Weekday>。
