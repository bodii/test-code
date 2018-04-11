/**
 * 命名空间
 * 
 * "内部模块"现在称做“命名空间”。“外部模块”现在则简称为“模块”。
 * 与ECMAScript2015里的术语保持一致，（也就是说module x { 相当于现在推荐的写法namespace x{ }
 */

 interface StringValidator {
     isAcceptable(s: string): boolean;
 }

 let lettersRegexp = /^[A-Za-z]+$/;
 let numberRegexp = /^[0-9]+$/;

 class LettersOnlyValidator implements StringValidator {
     isAcceptable(s: string) {
         return lettersRegexp.test(s);
     }
 }

 class ZipCodeValidator implements StringValidator {
     isAcceptable(s: string) {
         return s.length === 5 && numberRegexp.test(s);
     }
 }

 let strings = ['Hello', '98052', '101'];

 let validators: { [s: string]: StringValidator } = {};
 validators['ZIP code'] = new ZipCodeValidator();
 validators['Letters only'] = new LettersOnlyValidator();

 for (let s of strings) {
     for (let name in validators) {
         let isMatch = validators[name].isAcceptable(s);
         console.log(`'${ s }' ${ isMatch ? "matches": "does not match" } '${ name }'`);
     }
 }


 // 使用命名空间的验证器
 namespace Validation {
     export interface StringValidator {
         isAcceptable(s: string): boolean;
     }

     const lettersRegexp = /^[A-Za-z]+$/;
     const numberRegexp = /^[0-9]+$/;

     export class LettersOnlyValidator implements StringValidator {
         isAcceptable(s: string) {
             return lettersRegexp.test(s);
         }
     }

     export class ZipCodeValidator implements StringValidator {
         isAcceptable(s: string) {
             return numberRegexp.test(s);
         }
     }
 }

 let strings2 = ['Hello', '98052', '101'];

 let validators2 : {[s: string]: Validation.StringValidator} = {};

 validators2['ZIP code'] = new Validation.ZipCodeValidator();
 validators2['Letters only'] = new Validation.LettersOnlyValidator();

for (let s of strings2) {
    for (let name in validators2) {
        console.log(`"${ s }" - ${ validators2[name].isAcceptable(s) ? 'matches': 'does not match' } ${ name } `);
    }
}



/**
 * 别名
 * 
 * 另一种简化命名空间操作的方法是使用import q = x.y.z给常用的对象起一个短的名字。不要与用来加载模块的
 * import x = require('name')语法弄混，这里的语法是为指定的符号创建一个别名。可以用这种方法为任意标识
 * 符创建别名， 也包括导入的模块中的对象。
 */

 namespace Shapes {
     export namespace Polygons {
         export class Triangle { }
         export class Square { }
     }
 }

 import polygons = Shapes.Polygons;
 let sq = new polygons.Square();//

 /* 
 没有使用require关键字，而是直接使用导入符号的限定名赋值。这与使用var相似，但它还适用于类型和
 导入的具有命名空间含义的符号。重要的是，对于值来讲，import会生成与原始符号不同的引用，所以改变
 别名的var值并不会影响原始变量的值。
 */


 /*
     使用其它的JavaScript库

     声明是因为它不是外部程序的具体实现。我们通常在.d.ts里写这些声明。
 */


 /**
  * 外部命名空间
  * 
  * 为了让TypeScript编译器识别它的类型，使用用外部命名空间声明
  */
 // D3.d.ts(部分摘录)
 declare namespace D3 {
     export interface Selectors {
         select: {
             (selector: string): Selection;
             (element: EventTarget): Selection;
         };
     }

     export interface Event {
         x: number;
         y: number;
     }

     export interface Base extends Selectors {
         event: Event;
     }
 }

 declare var d3: D3.Base;