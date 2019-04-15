/**
 * 构造函数与原型
 */

// 构造函数
var Person = function(name, age) {
	this.name = name;
	this.age = age;
}


// Person.prototype 为Person的原型，这里在原型上添加了一个方法
Person.prototype.getName = function() {
	return this.name;
}

/*
	具体某一个人的特定属性，通常放在构造函数中。

	所有人公共的方法与属性，通常会放在原型对象中。 
 */

var p1 = new Person('Jake', 20);
var p2 = new Person('Tom', 22);

// 创建实例
console.log(p1.getName()); 
console.log(p2.getName());

/**
 * 构造函数其实与普通函数并无区别，首字母大写是一种约定，用来表示这是一个
 * 构造函数。但是new关键字的存在，让构造函数变得与众不同。
 */


function New(func) {
	// 声明一个中间对象，该对象为最终返回的实例
	var res = {};
	if (func.prototype !== null) {
		// 将实例的原型指向构造函数的原型
		res.__proto__ = func.prototype;
	}

	// ret为构造函数的执行结果，这里通过apply,
	// 将构造函数内部的this指向修改为指向res,即实例对象
	var ret = func.apply(res, Array.prototype.slice.call(arguments, 1));

	// 当我们在构造函数中明确指定了返回对象时，
	// 那么new的执行结果就是该返回对象
	if ((typeof ret === 'object' || typeof ret === 'function') && ret !== null) {
		return ret;
	}

	// 如果没有明确指定返回对象，则默认返回res, 这个res就是实例对象
	return res;

}

var p1 = New(Person, 'Jake');
var p2 = New(Person, 'Tom');

console.log(p1.getName());
console.log(p2.getName());

// new关键字在创建
// 1. 先创建一个新的、空的实例对象
// 2. 将实例对象的原型，指向构造函数的原型
// 3. 将构造函数内部的this,修改为指向实例
// 4. 最后返回该实例对象
/*
 表示指向:
 Person.prototype -> PPrototype;
 p1.__proto__ -> PPrototype;
 p2.__proto__ -> PPrototype;
 PPrototype.constructor -> Person
*/

// 构造函数的prototype与所有实例的__proto__都指向原型对象。
// 而原型对象的constructor则指向构造函


function Person(name) {
	this.name = name;
	this.getName = function() {
		return this.name + ', 你正在访问私有方法。';
	}
}

Person.prototype.getName = function() {
	return this.name + ', 你正在访问公共方法。';
}

var p1 = new Person('Tom', 20);
console.log(p1.getName()); 

// 可以通过in来判断一个对象是否拥有某一个方法/属性，无论该方法/属性是否公有
console.log('name' in p1);
console.log('getName' in p1);
console.log('gender' in p1);


// 特性检测，只有移动端环境才支持touchstart事件
// var isMobile = 'ontouchstart' in document;
// console.log(isMobile);
