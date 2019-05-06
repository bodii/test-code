'use static'
/*
 模板字符串

 模板字符串使用反引号"`"将整个字符串包裹起来，变量或者表达式则
 使用`${}`来包裹。
 还可用来定义多行字符串，其中所有的空格、缩进、换行都会被保留下来
 */

// ES5
var a = 20;
var b = 30;
var string = a + '+' + b + '=' + (a + b);


// ES6
const a = 20;
const b = 30;
const string = `${a} + ${b} = ${a+b}`; 

var elemString = `<div>
	<p>So you crossed the line, how'd you get that far?<p>
	<p>${word} you give it a try.</p>
<div>`;

const hello = 'hello';
let message = `${hello}, world!`;

const a = 40;
const b = 50;
let result = `the result is: ${a + b}`;

let fn = () => {
	const result = 'you are the best.';
	return result;
}

let str = `he said: ${fn()}`;

// 在一个函数中，当传入的参数是数组或者对象时，也可以使用解析结构
const fn = ({name, age}) => {
	return `${name} is age is ${age}`;
}

fn({name: 'TOM', age: 20 });

// 对象解析结构中的默认值
const {name, age=20} = 'tom';

// 数组解析中的默认值
const [a, b=20] = [1]

// 函数参数中的默认值
const fn = (x=20, y=30) => x + y;
console.log(fn());
