/**
 * 当函数调用call/apply时，则表示会执行该函数，并且函数内部的this
 * 指向call/apply的第一个参数.
 *
 *
 * call的第一个参数是为函数内部指定this指向，后续的参数则是函数执行
 * 时所需要的参数，一个个传递。
 *
 * apply的第一个参数与call相同，为函数内部this指向，而函数的参数，则
 * 以数组的形式传递，作为apply的第二个参数。
 */

function fn(num1, num2) {
	return this.a + num1 + num2;
}

var a = 20;
var object = { a: 40 };

console.log(fn(10, 10));

// 通过call改变this指向
console.log(fn.call(object, 10, 10));

// 通过apply改变this指向
console.log(fn.apply(object, [10, 10]));
