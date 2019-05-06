// ES6
const person = {
	name,
	age
}

// 等价于ES5
var person = {
	name: name,
	age: age
};

// 例如在一个模块中对外提供接口时
const getName = () => person.name;
const getAge = () => person.age;

// commonJs的方式
module.exports = { getName, getAge }

// ES6 modules的方式
export default { getName, getAge }



// 对象中方法的简写
// ES6
const person = {
	name,
	age,
	getName() {
		return this.name
	}
}

// ES5
var person = {
	name: name,
	age: age,
	getName: function getName() {
		return this.name;
	}
};


// 可以使用变量作为对象的属性，只需用中括号[]包裹即可
const name = 'Jane';
const age = 20;

const person = {
	[name]: true,
	[age]: true
}



// class 

//ES5
// 构造函数
function Person(name, age) {
	this.name = name;
	this.age = age;
}

// 原型方法
Person.prototype.getName = function() {
	return this.name
}


// ES6
class Person {
	constructor(name, age) { // 构造函数
		this.name = name;
		this.age = age;
	}

	getName() { // 原型方法
		return this.name;
	}
}



class Person {
	constructor(name, age) {
		this.name = name;
		this.age = age;
	}

	getName() {
		return this.name
	}

	static a = 20; // 等同于Person.a = 20
	c = 20; // 表示在构造函数中添加属性，在构造函数中等同于this.c = 20

	// 箭头函数的写法表示在构造函数中添加方法，在构造函数中等同于this
	getAge = function() {}
	getAge = () => this.age
}


// 继承
class Student extends Person {
	constructor(name, age, gender, classes) {
		super(name,age);
		this.gender = gender;
		this.classes = classes;
	}

	getGender() {
		return this.gender;
	}

}


const s = new Student('Tom', 20, 1, 3);
s.getName();
s.getGender();

// 只需用一个extends关键字与super方法，就能够实现继承

// 构造函数中
// ES6
super(name, age);

// 等价于ES5
Pserson.call(this);
