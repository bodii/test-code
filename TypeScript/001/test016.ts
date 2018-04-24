/**
 * 声明合并
 * 
 * 声明合并 是指编译器将针对同一个名字的两个独立声明合并为单一声明。合并后的声明同时
 * 拥有原先两个声明的特性。任何数量的声明都可被合并；不局限于两个声明。
 * 
 * TypeScript中的声明会创建以下三种实体之一: 命名空间，类型或值。创建命名空间，它包
 * 含了用(.)符号来访问时使用的名字。创建类型的声明是：用声明的模型创建一个类型并绑定
 * 到给定的名字上。最后，创建值的声明会创建在JavaScript输出中看到的值。
 */
interface Box {
    height: number;
    width: string;
}

interface Box {
    scale: string;
}

let box: Box = { height: 5, width: '6', scale: '10' };

/*
!!!
接口的非函数的成员应该是唯一的。如果它们不是唯一的，那么它们必须是相同的类型。如果两个接口同时
声明了同名的非函数成员且它们的类型不同，则编译器会报错。
对于函数成员，每个同名函数声明都会被当成这个函数的一个重载。同时需要注意，当接口A与后来的接口A
合并时，后面的接口具有更高的优先级。
*/

interface Animal {
    animal: string;
}
interface Sheep {
    animal: string;
}
interface Dog {
    animal: string;
}
interface Cat {
    animal: string;
}

interface Cloner {
    clone(animal: Animal): Animal;
}

interface Cloner {
    clone(animal: Sheep): Sheep;
}

interface Cloner {
    clone(animal: Dog): Dog;
    clone(animal: Cat): Cat;
}


// 这三个接口合并成一个声明：
interface Cloner {
    clone(animal: Dog): Dog;
    clone(animal: Cat): Cat;
    clone(animal: Sheep): Sheep;
    clone(animal: Animal): Animal;
}

// 注意每组接口里的声明顺序保持不变，但各组接口之间的顺序时后来的接口重载出现的靠前的位置。
// 这个规则有一个例外是当出现特殊的函数签名时。如果签名里有一个参数的类型是单一的字符串字面量
// (比如，不是字符串字面量的联合类型)，那么它将会被提升到重载列表的最顶端。


interface Document {
    createElement(tagName: any): Element;
}
interface Document {
    createElement(tagName: 'div'): HTMLDivElement;
    createElement(tagName: 'span'): HTMLSpanElement;
}
interface Document {
    createElement(tagName: string): HTMLElement;
    createElement(tagName: 'canvas'): HTMLCanvasElement;
}

// 合并后的 Document 将会像下面这样
interface Document {
    createElement(tagName: 'cavas'): HTMLCanvasElement;
    createElement(tagName: 'div'): HTMLDivElement;
    createElement(tagName: 'span'): HTMLSpanElement;
    createElement(tagName: string): HTMLElement;
    createElement(tagName: any): Element;
}


/**
 * 合并全名空间
 * 
 * 与接口相似，同名的命名空间也会合并其成员。命名空间会创建出命名空间和值，我们需要知道
 * 这两者都是怎么合并的。
 * 对于命名空间的合并，模块导出的同名接口进行合并，构成单一命名空间内含合并后的接口。
 * 
 * 对于命名空间里值的合并，如果当前已经存在给定名字的命名空间，那么后来的命名空间的导出
 * 成员会被加到已经存在的那个模块里。
 */

 namespace Animals {
     export class Zebra { }
 }

 namespace Animals {
     export interface Legged { numberOfLegs: number; }
     export class Dog { }
 }

 // 等同于：
//  namespace Animals {
//      export interface Lgged { numberOfLegs: number; }
     
//      export class Zebra { }
//      export class Dog { }
//  }

// 更清晰的说明：
namespace Animal {
    let haveMuscles = true;

    export function animalsHaveMuscles() {
        return haveMuscles;
    }
}

namespace Animal {
    export function doAnimalsHaveMuscles() {
        // return haveMuscles; // err, is not visible here
    }
}
// 错误因为haveMuscles并没有导出，只有animalsHaveMuscles函数共享了原始未合并的命名空间可以
// 访问这个变量。 doAnimalsHaveMuscles函数虽然是合并命名空间的一部分，但是访问不了未导出的成员。


/**
 * 命名空间与类和函数和枚举类型合并
 * 
 * 命名空间可以与其它类型的声明进行合并。只要命名空间的定义符合将要合并类型的定义。合并结果包含两者
 * 的声明类型。TypeScript使用这个功能去实现一些JavaScript里的设计模式。
 */

// 合并命名空间和类
// class Album {
//     lable: Album.AlbumLabel;
// }
// namespace Album {
//     export class AlbumLabel { }
// }
function buildLabel(name: string): string {
    return buildLabel.prefix + name + buildLabel.suffix;
}

namespace buildLabel {
    export let suffix = '';
    export let prefix = 'Hello, ';
}

console.log(buildLabel('Sam Smith'));

// 扩展枚举类型
enum Color {
    red = 1,
    green = 2,
    blue = 4
}

namespace Color {
    export function mixColor(colorName: string) {
        if (colorName == 'yellow') {
            return Color.red + Color.green;
        }
        else if (colorName == 'white') {
            return Color.red + Color.green + Color.blue;
        }
        else if (colorName == 'magenta') {
            return Color.red + Color.blue;
        }
        else if (colorName == 'cyan') {
            return Color.green + Color.blue;
        }
    }
}


// TypeScript并非允许所有的合并。目前，类不能与其它类或变量合并。

// 模块扩展
/* JavaScript
export class Observable<T> {

}

// map.js
import { Observable } from './observable';
Observable.prototype.map = function(f) {

}
*/

// 它也可以很地工作在TypeScript中，但编译器对Observable.prototype.map一无所知。你可以使用
// 扩展模块来将它告诉编译器
//map.ts
/*
import { Observable } from './observable';
declare module './observale' {
    interface Observable<T> {
        map<U>(f: (x: T) => U): Observable<U>;
    }
}
Observable.prototype.map = function (f) {

}

// consumer.ts
import { Observable } from './observable';
import './map';
let o: Observable<number>;
o.map(x => x.toFixed());
*/


// 全局扩展
/*
export class Observable<T> {

}

declare global {
    interface Array<T> {
        toObservable(): Observable<T>;
    }
}

Array.prototype.toObservable = function () {

}
*/