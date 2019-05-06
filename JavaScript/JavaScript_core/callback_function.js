// 利用函数调用栈，将想要执行的代码放入回调函数中

function want() {
	console.log('这是你想要执行的代码');
}

function fn(want) {
	console.log('这里表示执行了一大堆各种代码');

	// 其他代码执行完毕后，最后执行回调函数
	want && want();
}


fn(want);


function want1() {
	console.log('这是你想要执行的代码');
}

function fn1(want) {
	want && setTimeout(want, 0);
	console.log('这里表示执行了一大堆各种代码');
}

fn1(want1);
