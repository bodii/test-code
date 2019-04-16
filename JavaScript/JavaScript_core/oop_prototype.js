// 在原型上添加方法和属性
function Person() {}
Person.prototype.getName = function() {}
Person.prototype.getAge = function() {}
Person.prototype.sayHello = function() {}

// 对象字面量的写法来添加原型方法
Person.prototype = {
	constructor: Person,
	getName: function() {},
	getAge: function() {},
	sayHello: function() {}
}

// 当使用Person.prototype = {}时，其实是将Person的原型指向一个新的对象{}.
// 如果不做特殊处理，那么将会导致原型对象丢失。因此在这个新的对象中，需要
// 将它的constructor属性指向构造函数Person，这样就重新建立了正确的对应关系


/*
	原型链

	当一个对象A作为原型时，它有一个constructor属性指向它的构造函数，即A.constructor.
	当一个对象B作为构造函数时，它有一个prototype属性指向它的原型，即B.prototype.
	当一个对象C作为实例时，它有一个__proto__属性指向它的原型，即C.__proto__.
	当想要判断一个对象foo是否是构造函数Foo的实例时，可以使用instanceof关键字
	foo instanceof Foo

	当创建一个对象时，可以使用new Object()来创建。因此Object其实是一个构造函数，而
	其对应的原型Object.prototype则是原型链的终点。
*/
// 所有函数与对象都有一个toString与valueOf方法,
// 就是来自于Object.prototype
Object.prototype.toString = function() {}
Object.prototype.valueOf = function() {}


// 当创建函数时，除可以使用function关键字外，还可以使用Function对象
var add = new Function('a', 'b', 'return a + b');

// 等价
var add = function(a, b) {
	return a + b;
}


// add方法是一个实例，它对应的构造函数时Function,它的原型是Function.prototype
console.log(add.__proto__ === Function.prototype); // true

// Function同时是Function.Prototype的构造函数与实例
console.log(add.__proto__ === Function.prototype); // true
console.log(Function.prototype.constructor == Function ); // true
console.log(Function.prototype === Function.prototype);
console.log(Function.__proto__ === Function.prototype);
console.log(add instanceof Function);

console.log(Function.prototype.__proto__ === Object.prototype);
console.log(Object.__proto__ === Function.prototype);
console.log(Object instanceof Function);


function add() {}
console.log(add.toString === Object.toString);
