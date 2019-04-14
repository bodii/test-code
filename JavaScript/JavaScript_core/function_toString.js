/**
 * 当函数直接参与其他计算时，函数会默认调用toString方法，直接将函数
 * 体转换为字符串参与计算
 */

function fn() { return 20; }
console.log(fn + 10);
// out: function() { return 20; }10


// 可以重写toString方法
fn.toString = function() { return 20; }
console.log(fn + 10);


// 重写vauleOf方法
fn.valueOf = function() { return 30; }
console.log(fn + 10);

// 当同时重写函数toString方法与valueOf方法时，最终的终果会取valueOf
fn.valueOf = function() { return 1; }
fn.toString = function() { return 100; }
console.log(fn + 10);


function add() {
	// 第一次执行时，定义一个数组专门用来存储所有的参数
	var _args = [].slice.call(arguments);

	// 在内部声明一个函数，利用闭包的特性保存_args并收集所有的参数值
	var adder = function() {
		var _adder = function() {
			// [].push.apply(_args, [].slice.call(arguments)); // ES5
			_args.push(...arguments);
			return _adder;
		};

		_adder.toString = function() {
			return _args.reduce(function (a, b) {
				return a + b;
			});
		}

		return _adder;
	}
	// return adder.apply(null, _args); // ES5
	return adder(..._args); // ES6
}

var a = add(1)(2)(3)(4);
console.log('a: ', a);

var b = add(1, 2, 3, 4);
console.log('b: ', b);

var c = add(1, 2)(3, 4);
console.log('c: ', c);

var d = add(1, 2, 3)(4);
console.log('d: ', d);

console.log(a + 10);

console.log(a(10) + 100);
