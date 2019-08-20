#### java.lang.Class

* Field[] getFields()
* Field[] getDeclaredFields()
	getFields 方法将返回一个包含Field对象的数组，这些对象记录了这个类或其超类的公有域。
	getDeclaredField方法也将返回包含Field对象的数组，这些对象记录了这个类的全部域。如
	果类中没有域，或者Class对象描述的是基本类型或数组类型，这些方法将返回一个长度为0的
	数组。

* Method[] getMethods()
* Method[] getDeclaredMethods()
	返回包含Method对象的数组: 
		getMethods将返回所有的公有方法，包括从超类继承来的公有方法；
		getDeclaredMethods返回这个类或接口的全部方法，但不包括由超类继承了的方法。

* Constructor[] getConstructors()
* Constructor[] getDeclaredConstructors()
	返回包含Constructor对象的数组，其中包含了Class对象所描述的类的所有公有构造器(
	getConstructors)或所有构造器(getDeclaredConstructors)。



#### java.lang.reflect.Field
#### java.lang.reflect.Method
#### java.lang.reflect.Constructor

* Class getDeclaringClass()
	返回一个用于描述类中定义的构造器、方法或域的Class对象

* Class[] getExceptionTypes()(在Constructor和Method类中）
	返回一个用于描述方法抛出的异常类型的Class对象数组。

* int getModifiers()
	返回一个用于描述构造器、方法或域的修饰符的整型数值。使用Modifier类中的这个方法可以
	分析这个返回值。

* String getName()
	返回一个用于描述构造器、方法或域名的字符串。

* Class[] getParameterTypes()(在Constructor和Method类中)
	返回一个用于描述参数类型的Class对象数组

* Class getReturnType()(在Method类中)
	返回一个用于描述返回类型的Class对象。


#### java.lang.reflect.Modifier
* static String toString(int modifiers)
返回对应modifiers中位设置的修饰符的字符串表示。
* static boolean isAbstract(int modifiers)
* static boolean isDinal(int modifiers)
* static boolean isInterface(int modifiers)
* static boolean isNative(int modifiers)
* static boolean isPrivate(int modifiers)
* static boolean isProtected(int modifiers)
* static boolean isPublic(int modifiers)
* static boolean isStatic(int modifiers)
* static boolean isStrict(int modifiers)
* static boolean isSynchronized(int modifiers)
* static boolean isVclatile(int modifiers)
