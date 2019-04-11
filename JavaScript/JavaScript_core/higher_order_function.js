/**
 * 高阶函数
 *
 */


function Person(name, age) {
	this.name = name;
	this.age = age;
}

Person.prototype.getName = function() {
	return this.name;
}

var p1 = new Person('Jake', 18);
console.log(p1.getName());

/*
	构造函数其实就是普通的函数，而this是在函数运行时才确认的。
	new关键字
*/

function New(func) {

	// 声明一个中间对象，该对象为最终返回的实例
	var res = {};
	if (func.prototype !== null) {
		// 将实例的原型指向构造函数的原型
		res.__proto__ = func.prototype;
	}

	// ret为构造函数执行的结果，这里通过apply
	// 将构造函数内部的this指向修改为指向res,即为实例对象
	var ret = func.apply(res, Array.prototype.slice.call(arguments, 1));

	// 当在构造函数中明确指定了返回对象时，那么new的执行结果就是该返回对象
	if ((typeof ret === 'object' || typeof ret === 'function') && ret !== null) {
		return ret;
	}

	// 如果没有明确指定返回对象，则默认返回res,这个res就是实例对象
	return res;
}


// 使用上例中封装的New方法来创建实例
var p1 = New(Person, 'Jake', 18);
var p2 = New(Person, 'Tom', 20);
console.log(p1.getName()); 
console.log(p2.getName());

/*
	如果把当前函数看成基础函数，那么高阶函数，就是让当前函数获得额外能力的函数。如果
	把构造函数看成基础函数，那么New方法，就是构造函数的高阶函数。构造函数本就和普通函
	数一样，没有什么区别。但是因为New的存在，它获得了额外的能力。New方法每次执行都会
	创建一个新的中间对象，并将中间对象的原型指向构造函数的原型，将构造函数的this指向该
	中间对象。这样统一逻辑的封装，就是高阶函数的运用。

	如果简单来解释，凡是接收一个函数作为参数的函数，就是高阶函数。高阶函数其实是一个高
	度封装的过程。
 */
