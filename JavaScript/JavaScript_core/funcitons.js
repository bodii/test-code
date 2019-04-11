/**
 * 函数是一等公民
 *
 * 所谓的"一等公民", 其实就是普通公民。也就是说，函数其实没有什么特殊的，我们可
 * 以像对待任何其他数据类型一样对待函数
 */

// 把函数赋值给一个变量
var fn = function() {};

// 把函数存在数组里
function fn(callback) {
	var a = 20;
	return callback(20, 30) + a;
}

function add(a, b) {
	return a + b;
}

fn(add);

// 还可以把函数作为另一个函数运行的返回值
function add(x) {
	var y = 20;
	return function() {
		return x + y;
	}
}

var _add = add(100);
_add();
