/*

	在Promise中，是通过catch的方式来捕获异常的，而当使用async时，则通过try/catch来捕获
	异常。
 */

function fn() {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			reject('some error.');
		}, 1000);
	})
}

const foo = async () => {
	try {
		await fn();
	} catch (e) {
		console.log(e);
	}
}


foo();


// 如果有多个await函数，那么只返回第一个捕获到的异常
function fn1() {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			reject('some error fn1.');
		}, 1000);
	})
}

function fn2() {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			reject('some error fn2.');
		}, 1000);
	})
}

const foo2 = async () => {
	try {
		await fn1();
		await fn2();
	} catch (e) {
		console.log(e);
	}
}

foo2();
