/*
 异步与同步

 同步是指当发起一个请求时，如果未得到请求结果，代码将会等待，直到结果出来，才会
 执行后面的代码。

 异步是指当发起一个请求时，不会等待请求结果，而是直接继续执行后面的代码。

 */


/*
 Promise

 1. new Promise 表示创建一个Promise实例对象
 2. Promise函数中的第一个参数为一个回调函数，也可以称之为executor。通常情况下，
 在这个函数中，会执行发起请求操作，并修改结果的状态值。
 3. 请求结果有三种状态，分别是pending（等待中，表示还没有得到结果）、resolved（
 得到了我们想要的结果，可以继续执行），以及rejected（得到了错误的，或者不是我们
 期望的结果，拒绝继续执行）。请求结果的默认状态为pending。在executor函数中，可以
 分别使用resolve与reject将状态修改为对应的resolved与rejected。resolve、reject是
 executor函数的两个参数，它们能够将请求结果的具体数据传递出去。
 4. Promise实例拥有的then方法，可用来处理当请求结果的状态变成resolved时的逻辑。then
 的第一个参数为一个回调函数，该函数的参数是resolve传递出来的数据。
 5. Promise实例拥有的catch方法，可用来处理当请求结果的状态变成rejected时的逻辑。catch
 的第一个参数为一个回调函数，该函数的参数是reject传递出来的数据。

 */


function fn() {
	return new Promise(function(resolve, reject) {
		setTimeout(function() {
			resolve(30);
		}, 1000);
	})
}


// 在该函数的基础上，可以使用async/await语法来模拟同步的效果
var foo = async function() {
	var t = await fn();
	console.log(t);
	console.log('next code');
}

foo();


// 异步
var foo = function() {
	fn().then(function(resp) {
		console.log(resp);
	});

	console.log('next code');
}

foo();


var tag = true;
var p = new Promise(function(resolve, reject) {
	if (tag) {
		resolve('tag is true');
	} else {
		reject('tag is false');
	}
})

p.then(function(result) {
	console.log(result);
})
.catch(function(err) {
	console.log(err);
});




// example 2
function fn2(num) {
	// 创建一个Promise实例
	return new Promise(function(resolve, reject) {
		if (typeof num == 'number') {
			// 修改结果状态为resolved
			resolve();
		} else {
			// 修改结果状态为rejected
			reject();
		}
	}).then(function() {
		console.log('参数是一个number值');
	}).catch(function() {
		console.log('参数不是一个number值');
	})
}


// 修改参数的类型，观察输出结果
fn2('12');

// 注意观察该语句的执行顺序
console.log('next code');



// example 3
function fn3(num) {
	return new Promise(function(resolve, reject) {
		// 模拟一个请求行为，2s以后得到结果
		setTimeout(function() {
			if (typeof num == 'number') {
				resolve(num);
			} else {
				var err = num + ' is not a number';
				reject(err);
			}
		}, 2000);
	})
	.then(function(resp) {
		console.log(resp);
	})
	.catch(function(err) {
		console.log(err);
	})
}

// 修改传入的参数类型，观察结果变化
fn3('abc');

// 注意观察该语句的执行顺序
console.log('next code');


// then 方法可以接收两个参数，第一个参数用来处理resolved状态的逻辑，第二个参数
// 用来处理rejected状态的逻辑
/*
fn('abc')
.then(function(resp) {
	console.log(resp);
}, function(err) {
	console.log(err);
});

// 因此，catch方法其实与下面的写法等价
fn('abc').then(null, function(err) {
	console.log(err);
});


then方法因为返回的仍是一个Promise实例对象，因此then方法可以嵌套使用。
在这个过程中，通过在内部函数末尾return的方式，能够将数据持续往后传递。
*/


function fn4(num) {
	return new Promise(function(resolve, reject) {
		setTimeout(function() {
			if (typeof num == 'number') {
				resolve(num);
			} else {
				var err = num + ' is not a number.'
				reject(err);
			}
		}, 2000);
	});
}

fn4(20)
.then(function(result) {
	console.log(result); // 20
	return result + 1;
})
.then(function(result) {
	console.log(result); // 21
	return result + 1;
})
.then(function(result) {
	console.log(result); // 22
	return result + 1;
})
.then(function(result) {
	console.log(result); // 23
	return result + 1;
})
.then(function(result) {
	console.log(result); // 24 
});

// 注意观察语句的执行顺序
console.log('next code');
