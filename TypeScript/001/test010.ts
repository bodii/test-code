/**
 * 高级类型
 * 
 */

/* 交叉类型

交叉类型是将多个类型合并为一个类型。
把现有的多种类型叠加到一起成为一种类型，它包含了所需的所有类型的特性。
例如：Person & Serializable & Loggable 同时是Person和Serializable和Loggable。就是说
这个类型的对象同时拥有了这三种类型的成员
*/
function extend<T, U>(first: T, second: U): T & U {
    let result = <T & U>{};
    for (let id in first) {
        (<any>result)[id] = (<any>first)[id];
    }
    for (let id in second) {
        if (!result.hasOwnProperty(id)) {
            (<any>result)[id] = (<any>second)[id];
        }
    }
    return result;
}

class Person {
    constructor(public name: string) {}
}
interface Loggable {
    log(): void;
}
class ConsoleLogger implements Loggable {
    log() {
        // ...
    }
}
var jim = extend(new Person('Jim'), new ConsoleLogger());
var n = jim.name;
jim.log();


/* 联合类型
联合类型与交叉类型很有关联，但是使用上却完全不同。偶尔会遇到这种情况，一个代码希望传入number
或string类型的参数
*/

function padLeft(value: string, padding: any) {
    if (typeof padding === 'number') {
        return Array(padding + 1).join(' ') + value;
    }
    if (typeof padding === 'string') {
        return padding + value;
    }
    throw new Error(`Expected string or number, got '${padding}'.`);
}
padLeft('Hello world', 4);

// let indentedString = padLeft('Hello  world', true); // err: 编译阶段通过，运行时报错


// 代替any, 可以使用 联合类型 做为padding的参数：
function padLeft2(value: string, padding: number | string) {
    if (typeof padding === 'number') {
        return Array(padding + 1).join(' ') + value;
    }
    if (typeof padding === 'string') {
        return padding + value;
    }
    throw new Error(`Exected string or number, got '${padding}'.`);
}

// let indentedString = padLeft2('Hello world', true); // err: 这样就达到了类型检测的目的



// 如果一个值是联合类型，我们只能访问此联合类型的所有类型里共有的成员。
interface Bird {
    fly();
    layEggs();
}
interface Fish {
    swim();
    layEggs();
}

function getSmallPet (): Fish | Bird {
    return;
}

let pet = getSmallPet();
pet.layEggs();
// pet.swim(); // errors



// 类型保护与区分类型
let pet2 = getSmallPet();
/*
if (pet2.swim) {
    pet.swim();
}
else if (pet2.fly) {
    pet.fly();
}
*/ // 每一个成员访问都会报错

//　为了让这段代码工作，我们要使用类型断言：
let pet3 = getSmallPet();
if ((<Fish>pet3).swim) {
    (<Fish>pet3).swim();
}
else if ((<Bird>pet3).fly) {
    (<Bird>pet3).fly();
}



// 自定义的类型保护
/*
    TypeScript的 类型保护机制
    类型保护就是一些表达式，它们会在运行时检查以确保在某个作用域里的类型。要定义一个类型
    保护，我们只要简单地定义一个函数，它的返回值是一个类型谓词
*/
function isFish(pet: Fish | Bird): pet is Fish {
    return (<Fish>pet).swim !== undefined;
}
// pet is Fish就是类型谓词。谓词为paramenterName is Type这种形式，parameterName必须是来
// 自于当前函数签名里的一个参数名。


if (isFish(pet)) {
    pet.swim();
}
else {
    pet.fly();
}
// !!! TypeScript不仅知道在if 分支里pet是Fish类型;它还清楚在else分支里，一定不是Fish类型，
// 一定是Bird类型



// typeof 类型保护
function isNumber(x: any): x is number {
    return typeof x === 'number';
}
function isString(x: any): x is string {
    return typeof x === 'string';
}
function padLeft3(value: string, padding: string | number) {
    if (isNumber(padding)) {
        return Array(padding + 1).join(' ') + value;
    }
    if (isString(padding)) {
        return padding + value;
    }
    throw new Error(`Excepted string or number, got '${padding}'.`);
}

// 可以不必将typeof x === 'number'抽象成一个函数，因为TypeScript可以将它识别为一个类型
// 保护。也就是可以直接在代码里检查类型
function padLeft4(value: string, padding: number | string) {
    if (typeof padding === 'number') {
        return Array(padding + 1).join(' ') + value;
    }
    if (typeof padding === 'string') {
        return padding + value;
    }
    throw new Error(`Excepted string or number, got '${padding}'.`);
}
/*
    typeof类型保护
    只有两种形式能被识别： typeof v === 'typename' 和 typeof v !== 'typename',
    typename 必须时number, string, boolean或symbol.
*/



/* instanceof类型保护
instancefo 类型保护是通过构造函数来细化类型的一种方式。
 */

 interface Padder {
     getPaddingString(): string
 }
class SpaceRepeatingPadder implements Padder {
    constructor(private numSpaces: number) {}
    getPaddingString() {
        return Array(this.numSpaces + 1).join(' ');
    }
}

class StringPadder implements Padder {
    constructor(private value: string) {}
    getPaddingString() {
        return this.value;
    }
}

function getRandomPadder() {
    return Math.random() < 0.5 ?
        new SpaceRepeatingPadder(4) :
        new StringPadder('  ');
}

// 类型为SpaceReatingPadder | StringPadder
let padder: Padder = getRandomPadder();
if (padder instanceof SpaceRepeatingPadder) {
    padder; // 类型细化为'SpaceRepeatingPadder'
}
if (padder instanceof StringPadder) {
    padder; // 类型细化为'StringPadder'
}


/*
!!! instanceof的右侧要求是一个构造函数， TypeScript将细化为：

1. 此构造函数的prototype属性的类型，如果它的类型不为any的话
2. 构造签名所返回的类型的联合
 */
