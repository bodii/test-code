/*
	异步问题不仅可以使用前面尝到的Promise来解决，还可以用async/await来解决 

	async/await是ES7中新增的语法。
 */

async function fn() {
	return 30;
}

/*
// 或者
const fn = async () => {
	return 30;
}
*/


console.log(fn());


/*
 
 可以发现fn函数运行后返回的是一个标准的Promise对象，因此可以猜到async其实是
 Promise的一个语法糖，目的是为了让写法更加简单，因此也可以使用Promise的相关
 语法来处理后续的逻辑
 */

fn().then(res => {
	console.log(res);
})


/*
 
 await的含义是等待，意思就是代码需要等待await后面的函数运行完并且有了返回结果
 之后，才继续执行下面的代码。这正是同步的效果.
 但是需要注意的是，await关键字只能在async函数中使用，并且await后面的函数运行后
 必须返回一个Promise对象才能实现同步的效果。
 当使用一个变量去接收await的返回值时，该返回值为Promise中resolve传递出来的值，也
 就是PromiseValue

 */

function fn2() {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve(30);
		}, 1000);
	})
}


const foo = async () => {
	const t = await fn2();
	console.log(t);
	console.log('next code');
}

foo();

// 在async函数中，当运行遇到await时，就会等待await后面的函数运行完毕，而不会直接
// 运行后面的代码

const foo2 = () => {
	return fn2().then(t => {
		console.log(t);
		console.log('next code');
	})
}

foo2();
