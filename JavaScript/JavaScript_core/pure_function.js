/**
 * 纯函数
 *
 * 相同的输入总会得到相同的输出，并且不会产生副作用的函数，就是纯函数。
 *
 * 在封装一个函数、一个库或一个组件时，其实都期望一次封装，多处使用，而纯函数
 * 刚好具备这样的特性。
 */

// example

function getLast(arr) {
	return arr[arr.length];
}

function getLast_(arr) {
	return arr.pop();
}

var source = [ 1, 2, 3, 4 ];

var last = getLast(source); // 返回结果4, 原数组不变
var last_ = getLast_(source); // 返回结果4,原数组最后一项被删除

// getLast与getLast_虽然都能够获得数组的最后一项值，但是getLast_改变了原数组。而
// 当原数组被改变，我们再次调用该方法时，得到的结果就会变得不一样。



// 获取url的
function getParams(url, param) {
	if (!/\?/.test(url)) {
		return null;
	}

	var search = url.split('?')[1];
	var array = search.split('&');
	
	for(var i = 0; i < array.length; i++) {
		var tmp = array[i].split('=');
		if (tmp['0'] === param) {
			return decodeURIComponent(tmp[1]);
		}
	}

	return null;
}

var url = 'https://www.baidu.com/s?tn=baidu&wd=javascript&rsv_s ug=1';
var wd = getParams(url, 'wd');
console.log(wd);
