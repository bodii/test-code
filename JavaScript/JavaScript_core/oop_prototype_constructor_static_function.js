/**
 * 实例方法、原型方法、静态方法
 
	构造函数中的this指向的是新创建的实例。因此在往this上添加方法与属性时，其实
	是在往新创建的实例上添加属性与方法，所以构造函数中的方法可称之为 实例方法

	通过prototype添加的方法，将会挂载到原型对象上，因此称为 原型方法.

	方法直接挂载在构造函数上，我们称之为 静态方法。
 */

function Foo() {
	// 实例方法
	this.bar = function() {
		return 'bar in Foo';
	}
}

// 静态方法
Foo.bar = function() {
	return 'bar in static';
}

// 原型方法
Foo.prototype.bar = function() {
	return 'bar in prototype';
}
