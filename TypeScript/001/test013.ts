/**
 * 模块
 * 
 *内部模块 现在称做 [ 命名空间 ]。
 *外部模块 简称 [ 模块 ]
 *与ECMAScript 2015里的术语保持一致，（也就是说module X{} 相当于现在推荐的写法
 *namespace X {}).
 */
/*
模块在其自身的作用域里执行，而不是在全局作用域里；这意味着定义在一个模块里的变量、函数、
类等等在模块外部是不可见的，除非你明确地使用export形式之一导出它们。相反，如果想使用其它
模块导出的变量，函数，类，接口等的时候，你必须要导入它们，可以使用import形式之一。

模块是自声明的；两个模块之间的关系是通过在文件级别上使用imports和exports建立的。
模块使用模块加载器去导入其它的模块。在运行时，模块加载器的作用是在执行此模块代码前去查找
并执行这个模块的所有依赖。

任何包含顶级import或export的文件都被当成一个模块。相反地，如果一个文件不带有顶级的import
或者export声明，那么它的内容被视为全局可见的

*/


// 导出
// 导出声明（任何声明（比如变量，函数，类，类型别名或接口）都能够通过添加export关键字来导出）

// Validation.ts
export interface StringValidator {
    isAcceptable(s: string): boolean;
}
// ZipCodeValidator.ts
export const numberRegexp = /^[0-9]+$/;
export class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}


// 导出语句
// 导出语句很便利，因为我们可能需要对导出的部分重命名，
class ZipCodeValidator2 implements StringValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}
export { ZipCodeValidator2 };
export { ZipCodeValidator2 as mainValidator };


// 重新导出
// 扩展其它模块，并且只导出那个模块的部分内容。重新导出功能并不会在当前模块导入那个模块或定义一个新局部变量

// ParseIntBasedZipCodeValidator.ts
export class ParseIntBasedZipCodeValidator {
    isAcceptable(s: string) {
        return s.length === 5 && parseInt(s).toString() === s;
    }
}
// 导出原先的验证器但做了重命名
// export { ZipCodeValidator as RegExpBasedZipCodeValidator } from './ZipCodeValidator';


// 或者一个模块可以包裹多个模块，并把他们导出的内容联合在一起通过语法： export * form 'module'.
//AllValidators.ts
// export * from './StringValidator';
// export * from './LettersOnlyValidator';
// export * from './ZipCodeValidator';



/*
    导入

模块的导入， 可以用import形式之一来导入其它模块中的导出内容。
*/
// import { ZipcodeValidator } from './ZipCodeValidator';
// let myValidator = new ZipCodeValidator();

// 可以对导入的内容重命名
// import { ZipCodeValidator as ZCV } from './ZipCOdeValidator';
// let myValidator = new ZCV();

// 将整个模块导入到一个变量，并通过它来访问模块的导出部分
// import * as validator from './ZipCodeValidator';
// let myValidator = new validator.ZipCodeValidator();


// 具有副作用的导入模块
// 尽管不推荐这么做，一些模块会设置一些全局状态供其它模块使用。这些模块可能没有任何的导出或用户根本就
// 不关注它的导出。使用下面的方法来导入这类模块
// import './my-module.js';



/*
    默认导出
    每个模块都可以有一个default导出。默认导出使用default关键字标记；并且一个模块只能够有一个default
    导出。需要使用一种特殊的导入形式来

*/

// JQuery.d.ts
// declare let $: JQuery;
// export default $;

// App.ts
// import $ from "JQuery";
// $("button.continue").html('Next Step...');


//  ！！！ 类和函数声明可以直接被标记为默认导出。标记为默认导出的类和函数的名字是可以省略的。


// ZipCodeValidator.ts
// export defautl class ZipCodeValidator {
//     static numberRegexp = /^[0-9]+$/;
//     isAcceptable(s: string) {
//         return s.length === 5 && this.numberRegexp.test(s);
//     }
// }

//　Test.ts
// import validator from "./ZipCodeValidator";
// let myValidator = new validator();



// 或者
// StaticZipCodeValidator.ts
const numberRegexp2 = /^[0-9]+$/;
export default function (s: string) {
    return s.length === 5 && numberRegexp2.test(s);
}

// Test.ts
// import validate from './StaticZipCodeValidator';

// let strings = [ 'Hello', '98052', '101'];

// strings.forEach(s => {
//     console.log(`"${s}" ${validate(s) ? " matches " : " does not match "}`);
// });


// !!! default 导出也可以是一个值

// OneTwoThree.ts
// export default '123';

// Log.ts
// import num from './OneTwoThree';
// console.log(num); // 123



/*
export = 和 import = require()

exprot = 语法定义一个模块的导出对象。它可以是类，接口，命名空间，函数或枚举。
若要导入一个使用了export=的模块时，必须使用TypeScript提供的特定语法import module = require("module")
*/

// ZipCodeValidator.ts
let numberRegexp3 = /^[0-9]+$/;
class ZipCodeValidator3 {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp3.test(s);
    }
}
// export = ZipCodeValidator3;

// import zip = require('./ZipCodeValidator');
// let strings = ['Hello', '98052', '101'];
// let validator = new zip();
// strings.forEach(s => {
//     console.log(`"${ s }" - ${ validator.isAcceptable(s) ? 'matches': 'does not match'}`);
// })

// SimpleModule.ts
// import m = require('mod');
// export let t = m.something + 1;


/**
 * 外部模块
 * 
 * 可以使用顶级的export声明来为每个模块都定义一个.d.ts文件，但最好还是写在一个大的.d.ts文件里。
 * 使用与构造一个外部命名空间相似的方法，但是这里使用module关键字并且把名字用引号括起来，方便之后import
 * 
 */
/*
declare module 'url' {
    export interface Url {
        protocol?: string;
        hostname?: string;
        pathname?: string;
    }

    export function parse(urlStr: string, parseQueryString?, slashesDenotoHost?): Url;
}

declare module 'path' {
    export function normalize(p: string): string;
    export function join(...paths: any[]): string;
    export let sep: string;
}

// 可以使用 /// <reference> node.d.ts并且使用import url = require('url'); 或import * as URL
// from 'url' 加载模块


/// <reference path='node.d.ts'/>
import * as URL form 'url';
let myUrl = URL.parse('http://www.typescriptlang.org');

*/



/**
 * 外部模块简写
 * 
 * 可以采用声明的简写形式以便能够快速使用它
 * 
 */
/*
    declare module 'hot-new-module';

    简写模块里所有导出的类型将是any
    import x, {y} from 'hot-new-module';
    x(y);
*/

/*
  引入非ts文件
 
导入匹配'*!text' 或 'json!*'的内容
import fileContent from './xyz.txt!text';
import data from 'json!http://example.com/data.json';
console.log(data, fileContent);
 */


 /**
  * UMD模块
  * 
  * 有些模块被设计成兼容多个模块加载器，或者不使用模块加载器(全局变量)。 
  * 它们以UMD模块为代表。这些库可以通过导入的形式或全局变量的形式访问。
  */
 // math-lib.d.ts
//  export function isPrime(x: number): boolean;
//  export as namespace mathLib;
// 之后，这个库可以在某个模块里通过导入来使用：
// import { isPrime } from 'math-lib';
// isPrime(2);
// 不能在模块内使用全局定义：mathLib.isPrime(2);
// 它可以通过全局变量的形式使用，但只能在某个脚本(指不带有模块导入或导出的脚本文件)里。



/**
 * 创建模块结构指导
 * 
 * 尽可能地在顶层导出
 * 模块中导出一个命名空间就是一个增加嵌套的例子。虽然命名空间有时候有它们的用处，在使用模块
 * 的时候它们额外地增加了一层。这对用户来说是很不方便的并且通常是多余的。
 * 
 * 导出类的静态方法也同样的问题-这个类本身就增加了一层嵌套。除非它能方便表达或便于清晰使用，
 * 否则请考虑直接导出一个辅助方法。
 * 
 * !! 如果仅导出单个class或function,使用export default
 * 
 */

 // MyClass.ts
//  export defualt class SomeType {
//      constructor() { ... }
//  }

// MyFunc.ts
// export default function getThing() { return 'thing'; }

// Consumer.ts
// import t from './MyClass';
// import f from './MyFunc';
// let x = new t();
// console.log(f());


// 如果要导出多个对象，把它们放在顶层导出
// MyThings.ts
// export class SomeType { /* ... */ }
// export function someFunc() { /* ... */ }

// 当导入时，明确地列出导入的名字
// Consumer.ts
// import { SomeType, someFunc } from './MyThings';
// let x = new SomeType();
// let y = someFunc();


// 使用命名空间导入模式当你要导出大量内容的时候
// MyLargeModule.ts
/*
export class Dog { ... }
export class Cat { ... }
export class Tree { ... }
export class Flower { ... }
*/
// Consumer.ts
// import * as myLargeModule from './MyLargeModule.ts';
// let x = new myLargeModule.Dog();


/**
 * 使用重新导出进行扩展
 * 
 */

 // Calculator.ts
 export class Calculator {
     private current = 0;
     private memory = 0;
     private operator: string;

     protected processDigit(digit: string, currentValue: number) {
         if (digit >= '0' && digit <= '9') {
             return currentValue * 10 + (digit.charCodeAt(0) - '0'.charCodeAt(0))
         }
     }

     protected processOperator(operator: string) {
         if (['+', '-', '*', '/'].indexOf(operator) >= 0) {
             return operator;
         }
     }

     protected evaluateOperator(operator: string, left: number, right: number): number {
         switch (this.operator) {
             case '+': return left + right;
             case '-': return left - right;
             case '*': return left * right;
             case '/': return left / right;
         }
     }

     private evaluate() {
         if (this.operator) {
             this.memory = this.evaluateOperator(this.operator, this.memory, this.current);
         }
         else {
             this.memory = this.current;
         }
     }

     public handelChar(char: string) {
         if (char === '=') {
             this.evaluate();
             return;
         }
         else {
             let value = this.processDigit(char, this.current);
             if (value !== undefined) {
                 this.current = value;
                 return;
             }
             else {
                 let value = this.processOperator(char);
                 if (value !== undefined) {
                     this.evaluate();
                     this.operator = value;
                     return;
                 }
             }
         }
         throw new Error(`Unsupported input: '${char}'`);
     }

     public getResult() {
         return this.memory;
     }
 }

 export function test(c: Calculator, input: string) {
     for (let i = 0; i < input.length; i++) {
         c.handelChar(input[i]);
     }

     console.log(`result of '${input}' is '${c.getResult}'`);
 }


 // 下面使用导出的test函数来测试计算器
 // TestCalculator.ts
 // import { Calculator, test } from './Calculator';
 // let c = new Calculator();
 // test(c, '1+2*33/11='); // 9

 // ProgrammerCalculator.ts
// import { Calculator } from './Calculator';

class ProgrammerCalculator extends Calculator {
    static digits = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'];

    constructor(public base: number) {
        super();
        if (base <= 0 || base > ProgrammerCalculator.digits.length) {
            throw new Error('base has to be within 0 to 16 inclusive.');
        }
    }

    protected processDigit(digit: string, currentValue: number) {
        if (ProgrammerCalculator.digits.indexOf(digit) >= 0) {
            return currentValue * this.base + ProgrammerCalculator.digits.indexOf(digit);
        }
    }
}

// export { ProgrammerCalculator as Calculator };
// export { test } from './Calculator';

// 新的ProgrammerCalculator模块导出的API与原先的Calculator模块很相似，但却没有改变原模块里的对象。

// TestProgrammerCalculator.ts
// import { Calculator, test } from './ProgrammerCalculator';
// let c = new Calculator(2);
// test(c, '001+010='); // 3



// 模块里不要使用命名空间
