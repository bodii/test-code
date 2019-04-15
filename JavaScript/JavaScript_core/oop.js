/**
 * 面向对象
 *
 * 在ECMAScript-262中，对象被定义为 "无序属性的集合，其属性可以包含基本值、对象
 * 或者函数"
 * 也就是说，对象是由一系列无序的key-value对组成。其中value可以是基本数据类型、
 * 对象、数组、函数等
 */

var person = {
	name: 'Tom',
	age: 18,
	getname: function() {
		return this.name;
	},

	parent: {}
}

// 可以通过关键字new来创建一个对象
var obj = new Object();

// 也可以通过对象字面量的形式创建一个对象
var obj = {};

// 给创建的对象添加属性与方法
var person = {};
person.name = 'Tom';
person.getName = function() {
	return this.name;
}

// 也可以这样
var person = {
	name: 'Tom',
	getName: function() {
		return this.name;
	}
};


// 访问对象的属性
console.log(person.name);
// or
console.log(person['name']);
// or
var _name = 'name';
console.log(person[_name]);

// 访问的属性名是一个变更时，可以使用中括号的方式
['name', 'age'].forEach(function(item) {
	if (person.hasOwnProperty(item)) {
		console.log(person[item]);
	}
})
