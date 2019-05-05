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
