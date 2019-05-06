/*
 Promise.all 接收一个Promise对象组成的数组作为参数，当这个数组中所有
 的Promise对象状态都变成resolved或者rejected时，它才会去调用then方法
 */

function getJSON(url) {
	return new Promise(function(resolve, reject) {
		var XHR = new XMLHtttpRequest();
		XHR.open('GET', url, true);
		XHR.send();

		XHR.onreadystatechange = function() {
			if (XHR.readyState == 4) {
				if (XHR.status == 200) {
					try {
						var response = JSON.parse(XHR.responseText);
						resolve(response);
					} catch (e) {
						reject(e);
					}
				} else {
					reject(new Error(XHR.statusText));
				}
			}
		}
	})
}


var url = 'https://cn.bing.com';
var url1 = 'https://www.baidu.com';

function renderAll() {
	return Promise.all([getJSON(url), getJSON(url)]);
}

renderAll().then(function(value) {
	console.log(value);
});


/*
 
 与Promise.all相似的是，Promise.race也是一个Promise对象组成的数组作为参数，不同的
 是，只要当数组中的其中一个Promise状态变成resolved或者rejected时，就可以调用.then
 方法，而传递给then方法的值也会有所不同
 */

function renderRace() {
	return Promise.race([getJSON(url), getJSON(url1)]);
}

renderRace().then(function(value) {
	console.log(value);
});
