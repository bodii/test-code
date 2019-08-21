#### java.lang.reflect.Method
* public Object invoke(Oject implicitParameter, Object[] explicitParamenters)
	调用这个对象所描述的方法，传递给定参数，并返回方法的返回值。对于静态方法，
	把null作为隐式参数传递。在使用包装器传递基本类型的值时，基本类型的返回值必
	须是未包装的。


invoke的参数和返回值必须是Object类型的。这就意味着必须进行多次的类型转换。这样
做将会使用编译器错过检查代码的机会。因此，等到测试阶段才会发现这些错误，找到并
改正它们将会更加困难。不仅如此，使用反射获得方法指针的代码要比仅仅直接调用方法
明显慢一些。

建议不要使用Method对象的回调功能。使用接口进行回调会使得代码的执行速度更快，更
易于维护。
