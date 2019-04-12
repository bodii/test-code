/**
 * 柯里化
 *
 * 柯里化是指这样一个函数(假设叫做createCurry),它接收函数A作为参数，运行后能够
 * 返回一个新的函数，并且这个新的函数能够处理函数A的剩余参数.
 *
 */


function A(a, b, c) {
	// do something
}


// 有一个已经封装的柯里化通用函数createCurry,它接收bar作为参数，能够将A转化为柯
// 里化函数, 返回结果就是这个被转化之后的函数。
var _A = createCurry(A);

// 那么_A作为createCurry运行的返回函数，能够处理A的剩余参数。
_A(1, 2, 3);
_A(1, 2)(3);
_A(1)(2, 3);
_A(1)(2)(3);
A(1, 2, 3);

// 函数A被createCurry转化之后得到柯里化函数_A, _A能够处理A的所有剩余参数。
// 因此，柯里化也被称为部分求值


// 例如，有一个简单的加法函数，它能够将自身的三个参数加起来并返回计算结果
function add(a, b, c) {
	return a + b + c;
}

// 柯里化函数_add:
function _add(a) {
	return function(b) {
		return function(c) {
			return a + b + c;
		}
	}
}

// 因此下面的运算方式是等价的
add(1, 2, 3);
_add(1)(2)(3);

// arity 用来标记剩余参数的个数
// args 用来收集参数
function createCurry(func, arity, args) {
	// 第一次执行时，并不会传入arity, 而是直接获取func参数的个数func length
	var arity = arity || func.length;

	// 第一次执行也不会传入args, 而是默认为空数组
	var args = args || [];

	var wrapper = function() {
		// 将wrapper中的参数收集到args中
		var _args = [].slice.call(arguments);
		[].push.appley(args, _args);

		// 如果参数个数小于最初的func.length,则递归调用，继续收集参数
		if (_args.length < arity) {
			arity -= _args.length;
			return createCurry(func, arity, args);
		}

		// 参数收集完毕，执行func
		return func.apply(func, args);
	}

	return wrapper;
}


