/**
 * bind方法也能指定函数内部的this指向，但是它与call/apply有所不同.
 * 当函数调用call/apply时，函数的内部this被显式指定，并且函数会立即执行。而
 * 当函数调用bind时，函数并不会即执行，而是返回一个新的函数，这个新的函数与
 * 原函数有共同的函数体，但它并非原函数，并且新函数的参数与this指向都已经绑
 * 定，参数为bind的后续参数.
 */

function fn(num1, num2) {
	return this.a + num1 + num2;
}

var a = 20;
var object = { a: 40 };

var _fn = fn.bind(object, 1, 2);
console.log(_fn === fn); 
console.log(_fn());
console.log(_fn(1, 4));
