/*
 事件循环机制

 事件循环机制(Event Loop) 是全面了解JavaScript代码执行顺序绕不开的一个
 重要知识点。

 */


// demo01
setTimeout(function() {
	console.log(1);
}, 0);

console.log(2);

for(var i = 0; i < 5; i++) {
	console.log(3);
}

console.log(4);


// demo02
console.log(1);
for (var i = 0; i < 5; i++) {
	setTimeout(function() {
		console.log('2-' + i);
	}, 0);
}

console.log(3);

/*
 即使设置了setTimeout的延迟时间为0, 它里面的代码仍然是最后执行。

 通常情况下，决定代码执行顺序的是函数调用栈。很显然这里的setTimeout中的执行
 顺序已经不是用函数调用栈能够解释清了。
 
 JavaScript的一个特点就是单线程，但是很多时候我们仍然需要在不同的时间去执行不
 同的任务，例如给元素添加点击事件，设置一个定时器，或者发起Ajax请求。因此需要
 一个异步机制来达到这样的目的，事件循环机制也因此而来。

 每一个JavaScript程序都拥有唯一的事件循环，大多数代码的执行顺序是可以根据函数调
 用栈的规则执行的，而setTimeout/setInterval或者不同的事件绑定(click、mousedown等)
 中的代码，则通过队列来执行。

 setTimeout为任务源，或者任务分发器，由它们将不同的任务分发到不同的任务队列中去。
 每一个任务源都有对应的任务队列。
 
 任务队列又分为宏任务(macro-task)与微任务(micro-task)两种，在浏览器中，包括:
 1. macro-task: script(整体代码)、setTimeout/setInterval、I/O、UI rendering等
 2. micro-task: Promise

 注意：在node.js中还包括更多的任务队列。来自不同任务源的任务会进到不同的任务队列中，
 其中setTimeout与setInterval是同源的。

 事件循环的顺序，决定了JavaScript代码的执行顺序。
 
 它从macro-task中的script开始第一次循环。此时全局上下文进入函数调用栈，直到调用栈清空
 （只剩下全局上下文），在这个过程中，如果遇到任务分发器，就会将任务放入对应队列中去。

 第一次循环时，macro-task中其实只有script,因此函数调用栈清空之后，会直接执行所有的micro-task.
 当所有可执行的micro-task执行完毕之后，就表示第一次事件循环已经结束。

 第二次循环会再次从macro-task开始执行。此时macro-task中的script队列中已经没有任务了，
 但是可能会有其他的队列任务，而micro-task中暂时还没有任务。此时会先选择其中一个宏任务
 队列，例如setTimeout，将该队列中的所有任务全部执行完毕，然后再执行此过程中可能产生的微
 任务。微任务执行完毕之后，再回过头来执行其他宏任务队列中的任务。依次类推，直到所有宏任务
 队列中的任务都被执行一遍，并且清空了微任务，第二次循环就会结束。
 
 如果在第二次循环过程中，产生了新的宏任务队列，或者之前宏任务队列中的任务暂时没有满足执行
 条件，例如延迟时间不够或者事件没有触发，那么将会继续以同样的顺序重复循环。

 首先，macro-task script任务队列最先执行。执行过程中遇到setTimeout,setTimeout会将它的任务分
 发到setTimeout任务队列中去，因此里面的代码是不执行的。
 */


// example
setTimeout(function() {
	console.log('timeout1');
})

new Promise(function(resolve) {
	console.log('promise1');
	for (var i = 0; i < 1000; i++) {
		i == 99 && resolve();
	}
	console.log('promise2');
}).then(function() {
	console.log('then1');
})

console.log('global1');
