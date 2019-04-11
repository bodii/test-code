// 命令式编程
var array = [1, 3, 'h', 5, 'm', '4'];
var res = [];
for(var i=0; i < array.length; i++) {
	if (typeof array[i] === 'number')
		res.push(array[i]);
}


// 函数式编程
function getNumbers(arr) {
	var res = [];
	array.forEach(function(item) {
		if (typeof item === 'number') {
			res.push(item);
		}
	});

	return res;
}

var res = getNumbers(array);
console.log(res);
