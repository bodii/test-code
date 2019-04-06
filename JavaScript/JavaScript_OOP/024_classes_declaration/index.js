/**
 * 方法的声明有两种: function声明和表达式声明
 * 
 * 
 */

 sayHello();  // 在此处调用sayHello不会报错
//  sayGoodbye(); // 在此处调用sayGoodbye会报错
// 1. function declaration
// function声明默认是置顶的
function sayHello() {}



// 2. function expression 
const sayGoodbye = function() {};
// 表达式声明
// 和声明一个常量类似
const number = 1;


/**
 * 类声明也有两种： class declaration 和 表达式声明
 */

// const c = new Circle();  // 类声明没有置顶的概念
// console.log(c);

 // 1. class declaration
class Circle {}

// 2. class experssion
const Square = class {

};