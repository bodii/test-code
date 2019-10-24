#### 依赖注入(dependency injection)
当创建一个新的实例时，就将该资源传到构造器中。

这是依赖注入(dependency injection)的一种形式：
词典(dictionary)是拼写检查的一个`依赖(dependency)`, 在创建拼写检查器时就将词典`注入(injected)`其中。

```java
// Dependency injection provides flexibility and testability
public class SpellChecker {
    private final Lexicon dictionary;

    public SpellChecker(Lexicon dictionary) {
        this.dictionary = Objects.requireNonNull(dictionary);
    }

    public boolean isValid(String word) { }
    public List<String> suggestions(String typo) {}
}
```

依赖注入也同样适用于构造器、静态工厂和构建器。

这个程序模式的另一种有用的变体是，将资源`工厂(factory)`传给构造器。工厂是可以被重复调用来创建类型实例的一个对象。这类工厂具体表现为`工厂方法(Factory Method)`模式。
在Java 8中增加的接口Supplier<T>, 最适合用于表示工厂。带有Supplier<T>的方法，通常应该限制输入工厂的类型参数使用有限制的通配符类型(bounded wildcard type),以便客户端能够传入一个工厂，来创建指定类型的任意子类型。

例如，下面是一个生产马赛克的方法，它利用客户端提供的工厂来生产每一片马赛克:
```java
Mosaic create(Supplier<? extends Tile> tileFactory) {  }
```

虽然依赖注入极大地提升了灵活性和可测试性，但它会导致大型项目凌乱不堪，因为它通常包含上千个依赖。不过这种凌乱用一个`依赖注入框架(dependency injection framework)`便可以终结，如Dagger[Dagger]、Guice[Guice]或者Spring[Spring]。这些框架的用法超出了本收的讨论范畴，但是，请注意：设计成手动依赖注入的API，一般都适用于这些框架。

总而言之，不要用Singleton和静态工具类来实现依赖一个或多个底层资源的类，且该资源的行为会影响到该类的行为；也不要直接用这个类来创建这些资源。而应该将这些资源或者工厂传给构造器(或者静态工厂，或者构建者)，通过它们来创建表。这个实践就被称作依赖注入，它极大地提升了类的灵活性、可重用性和可测试性。