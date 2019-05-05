/*
 ES6变量声明

 一般来说，声明一个引用可以被改变的变量时用let,声明一个引用不能被改变的变量时用const
 */

let a = 20;
a = 30;
console.log(a);


const PI = 3.1415;
const MAX_LENGTH = 100;

const a = [];
a.push(1);
console.log(a);


const b = {};
b.max = 20;
b.min = 0;
console.log(b);


const arr = [1, 2, 3, 4];
arr.forEach(function(item) {
	const temp = item + 1;
	console.log(temp);
});

