#### 服务提供者框架(Service Provider Framework)
服务提供都框架中有三个重要的组件:
1. 服务接口(Service Interface), 这是提供者实现的；
2. 提供者注册API(Provider Registration API),这是提供者用来注册实现的；
3. 服务访问API(Service Access API)，这是客户端用来获取服务的实例。

服务访问API是客户端用来指定某种选择实现的条件。如果没有这样的规定，API就会返回默认实现的一
个实例，或者允许客户端遍历所有可用的实现。
服务访问API是"灵活的静态工厂", 它构成了服务提供者的框架的基础。

服务提供者框架的第四个组件`服务提供者接口(Service Provider Interface)`是可选的，它表示产生
服务接口之实例的工厂对象。如果没有服务提供者接口，实现就通过反射方式进行实例化。

Java平台提供了一个通用的服务提供者框架java.util.ServiceLoader,因此你不需要再自已编写。

