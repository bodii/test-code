/**
 * 在map函数的基础上，进行二次封装，以简化在项目中的使用
 *
 */

function getNewArray(array) {
	return array.map(function(item) {
		return item * 100 + '%';
	});
}

console.log(getNewArray([1, 2, 3, 0.12]));


// 借助柯里化，来二次封装
function createCurry(func, arity, args) {

	var arity = arity || func.length;

	var args = args || [];

	var wrapper = function() {

		var _args = [].slice.call(arguments);

		[].push.apply(args, _args);

		if (args.length < arity) {
			arity -= _args.length;

			return createCurry(func, arity, args);
		}

		return func.apply(func, args);
	}

	return wrapper;
}


function _map(func, array) {
	return array.map(func);
}

var _getNewArray = createCurry(_map);
var getNewArray = _getNewArray(function(item) {
	return item * 100 + '%';
});
console.log(getNewArray([1, 2, 3, 0.12]));

var _getNewArray = createCurry(_map);
var getNewArray = _getNewArray(function(item) {
	return item * 100 + '%';
});
console.log(getNewArray([0.01, 1]));

