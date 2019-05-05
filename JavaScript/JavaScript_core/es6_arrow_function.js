// es5
var fn = function(a, b) {
	return a + b;
}

// ES6 箭头函数写法
const fn = (a, b) => a + b;

// ES5
var foo = function() {
	var a = 20;
	var b = 30;
	return a + b;
}

// ES6
const foo = () => {
	const a = 20;
	const b = 30;
	return a + b;
};

/*
箭头函数中的this,就是声明函数时所处上下文中的this,它不会被其他方式所改变 
*/

var Mot = function(name) {
	this.name = name;
}

Mot.protoptype = {
	constructor: Mot,
	do: function() {
		console.log(this.name);
		document.onclick = function() {
			console.log(this.name);
		}
	}
}

new Mot('Alex').do();

/*
在这个例子中，当调用do方法时，我们期望单击document时仍然输出'Alex'。但是
很遗憾,在onclick的回调函数中，this的指向其实已经发生了变化，它指向了document,
因此肯定得不到我们想要的结果。
 */

var Mot = function(name) {
	this.name = name;
}

Mot.prototype = {
	constructor: Mot,
	do: function() {
		console.log(this.name);
		document.onclick = () => {
			console.log(this.name);
		}
	}
}

new Mot('Alex').do();


// 在箭头函数中， 没有arguments对象
var add = function(a, b) {
	console.log(arguments);
	return a + b;
}

add(1, 2);


var add = (a, b) => {
	console.log(arguments);
	return a + b;
}

add(1, 2); // arguments is not defined

