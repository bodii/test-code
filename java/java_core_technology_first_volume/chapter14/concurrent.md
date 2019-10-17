#### 同步器
java.util.concurrent包包含了几个能帮助管理相互合作的线程集的类。
这些机制具有线程之间的`共同集结点模式(common rendezvous patterns)`提供的"预置功能"(canned 
functionality)。如果有一个相互合作的线程集满足这些行为模式之一，那么应该直接重用合适的库类
而不要试图提供手工的锁与条件的集合。

