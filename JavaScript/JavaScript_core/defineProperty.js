/*
	Object.defineProperty只能设置一个属性的属性特性，
	当想要同时设置多个属性的特性时，需要使用之前介绍过的Object.defineProperties
 */


var Person = {};

Object.defineProperties(Person, {
	name: {
		value: 'Jake',
		configurable: true,
	},
	age: {
		get: function() {
			return this.value || 22;
		},
		set: function(value) {
			this.value = value;
		},
	},
});

console.log(Person.name);
console.log(Person.age);


var Person = {};

Object.defineProperty(Person, 'name', {
	value: 'alex',
	writable: false,
	configurable: false,
});

console.log(Person.name);


// 读取属性的特性值 Object.defineProperty

var descripter = Object.getOwnPropertyDescriptor(Person, 'name');
console.log(descripter);
