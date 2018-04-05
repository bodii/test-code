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
import validate from './StaticZipCodeValidator';

let strings = [ 'Hello', '98052', '101'];

strings.forEach(s => {
    console.log(`"${s}" ${validate(s) ? " matches " : " does not match "}`);
});


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