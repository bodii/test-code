/**
 * 继承
 *
 * 因为封装一个对象是由构造函数与原型共同组成的，所以继承也被分为
 * 有构造函数的继承与原型的继承两种
 */

var Person = function(name, age) {
	this.name = name;
	this.age = age;
}

Person.prototype.getName = function() {
	return this.name;
}

Person.prototype.getAge = function() {
	return this.age;
}

var Student = function(name, age, grade) {
	// 通过call方法还原Person构造函数中的所有处理逻辑
	Student.call(Person, name, age);
	// 等价于
	// Person.call(this, name, age);
	this.grade = grade;
}

// 等价于
var Student = function(name, age, grade) {
	this.name = name;
	this.age = age;
	this.grade = grade;
}


/**
 * 原型继承
 */

function create(proto, options) {
	// 创建一个空对象
	var tmp = {};

	// 让这个新的空对象成为父类对象的实例
	tmp.__proto__ = proto;

	// 传入的方法都挂载到新对象上，新对象将作为子类对象的原型
	Object.defineProperties(tmp, options);
	return tmp;
}

Student.prototype = create(Person.prototype, {
	// 不要忘了重新指定构造函数
	/*
	constructor: {
		value: Student
	},
	*/
	// 等价
	consrtuctor: Student,
	getGrade: {
		value: function() {
			return this.grade;
		}
	}
});


var s1 = new Student('ming', 22, 5);
console.log(s1.getName());
console.log(s1.getAge());
console.log(s1.getGrade());
