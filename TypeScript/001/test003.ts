
/**
 * 接口
 */

 // TypeScript的核心原则之一是对值所具有的结进行类型检查。它有时被称为“鸭式辨型法”
 // 或 “结构性子类型化”。在TypeScript里，接口的作用就是为这些类型命名和为你的代码
 // 或第三方代码定义契约。

 function printLabel(labelledObj: { label: string }) {
     console.log(labelledObj.label);
 }

 let myObj = { size: 10, label: 'Size 10 object' };
 printLabel(myObj);

 interface LabelledValue {
     label: string;
 }

 function printLabel2(labelledObj: LabelledValue) {
     console.log(labelledObj.label);
 }

 let myObj2 = { size: 10, label: 'Size 10 Object' };
 printLabel2(myObj2);

 // 只会关注值的外形。只要传入的对象满足上面提到的必要条件，那么它就被允许
 // 还有一点，类型检查器不会去检查属性的顺序，只要相应的属性存在并且类型也
 // 是对的就可以


 // 可选属性

 interface SquareConfig {
     color?: string;
     width?: number;
 }

 function createSquare(config: SquareConfig): {color: string; area: number} {
     let newSquare = {color: 'white', area: 100};

     if (config.color) {
         newSquare.color = config.color;
     }

     if (config.width) {
         newSquare.area = config.width * config.width;
     }
     return newSquare;
 }

 let mySquare = createSquare({color: 'black'});
 console.log('mySquare:' + mySquare);


 interface SquareConfig {
    color?: string;
    width?: number;
 }

 function createSquare2(config: SquareConfig) : { color: string, area: number } {
    let newSquare = { color: 'white', area: 100 };
    
    // if (config.clor) {
    //     newSquare.color = config.clor;
    // } 
    // err

    if (config.width) {
        newSquare.area = config.width * config.width;
    }

    return newSquare;
 }

 let mySquare2 = createSquare2({color: 'black' });


 /* 只读属性 */
// 一些对象属性只能在对象刚刚创建时候修改其值。你可以在属性名前用readonly来指定只读属性

interface Point {
    readonly x: number;
    readonly y: number;
}

let p1: Point = { x: 10, y: 20 };
// p1.x = 5;  // err


// TypeScript 具有ReadonlyArray<T>类型，它与Array<T>相似，只是把所有可变方法去掉了，
// 因此可以确保数组创建后再也不能被修
let a: number[] = [1, 2, 3, 4];
let ro: ReadonlyArray<number> = a;
// ro[0] = 12; // err
// ro.push(5); // err
// ro.length = 100; // err
// a = ro; // err

// 上面，可以看到就算把整个ReadonlyArray赋值到一个普通数组也是不可以的。
// 但是你可以用类型断言重写
a = ro as number[];



/*  readonly vs const */
// 最简单判断该用readonly还是const的方法是看要把它做为变量使用还是作为一个属性。
// 作为变量使用用const, 若做为属性则使用readonly。
// error
interface SquareConfig {
    color?: string;
    width?: number;
}

function createSquare3(config: SquareConfig) : { color: string; area: number } {
    let newSquare = { color: 'white', area: 100 };
    if (config.color) {
        newSquare.color = config.color;
    }
    if (config.width) {
        newSquare.area = config.width * config.width;
    }

    return newSquare;
}
// let mySquare3 = createSquare3({ colour: 'read', width: 100 });
/*
let mySquare3 = createSquare3({ colour: 'read', widht: 100 });
这段代码TypeScript会认为可能存在bug。对象字面量会被特殊对待而且会经过额外属性检查，当将它们
赋值给变量或作为参数传递的时候。如果一个对象字面量存在任何“目标类型”不包含的属性时，你会得到
一个错误。
*/
// 绕开这些检查非常简单，最简便的方法是使用类型断言：
let mySquare3 = createSquare2({ width: 100, opacity: 0.5} as SquareConfig);

// 然而，是佳的方式是能够添加一个字符串索引签名，前提是你能够确定这个对象可能具有某些作为特殊
// 用途的额外属性。
interface SquareConfig {
    color?: string;
    width?: number;
    [propName: string]: any; // 这里表示的是SquareConfig可以有任意数量的属性，并且只要它们
    // 不是color和width， 那么就无所谓它们的类型是什么。
}


// 最后一种跳过这些检查的方式， 就是将这个对象赋值给另一个变量
let squareOptions = { colour: 'red', widht: 100 };
let mySquare4 = createSquare3(squareOptions);



/* 函数类型 */
// 为了使用接口表示函数类型，我们需要给接口定义一个调用签名。
interface SearchFunc {
    (source: string, subString: string): boolean;
}

let mySearch: SearchFunc;
mySearch = function(source: string, subString: string) {
    let result = source.search(subString);
    return result > -1;
}

// 对于函数类型的类型检查来说，函数的参数名不需要与接口里定义的名字相匹配。

let mySearch2: SearchFunc;
mySearch2 = function(src: string, sub: string): boolean {
    let result = src.search(sub);
    return result > -1;
}



/* 可索引的类型 */
// 索引类型具有一个索引签名，它描述了对象索引的类型，还有相应的索引返回值类型。

interface StringArray {
    [index: number]: string; // 索引签名 index: number, 返回string类型
}

let myArray: StringArray;
myArray = ['Bob', 'Fred'];

let myStr: string = myArray[0];

// 共支持两种索引签名：字符串和数字。可以时使用两种类型的索引，但是数字索引
// 的返回值必须是字符串索引返回值类型的子类型。这是因为当使用number来索引时，
// JavaScript会将它转换成string然后再去索引对象。也就是说用100（一个number）
// 去索引等同于使用"100"（一个string）去索引，因此两者需要保持一致。

class Animal {
    name: string;
}

class Dog extends Animal {
    breed: string;
}

// err: 使用数值型的字符串索引，有时会得到完全不同的Animal !
// interface NotOkay {
//     [x: number]: Animal;
//     [x: string]: Dog;
// }


// 字符串索引签名能够很好的描述dictionary模式，并且它们也会确保所有属性与其返回值
// 类型相匹配。

interface NumberDictionary {
    [index: string]: number;
    length: number;
    // name: string; // err, 'name'的类型与索引类型返回值的类型不匹配
}

// 可以将索引签名设置为只读
interface ReadonlyDictionary {
    readonly [index: number]: string;
}

let myArray2: ReadonlyDictionary = ['Alice', 'Bob'];
// myArray2[2] = 'Mallory'; //err



/* 类类型 */
// 实现接口
interface ClockInterface {
    currentTime: Date;
}

class Clock implements ClockInterface {
    currentTime: Date;
    constructor(h: number, m: number) {}
}

interface ClockInterface2 {
    currentTime: Date;
    setTime(d: Date);
}

class Clock2 implements ClockInterface2 {
    currentTime: Date;
    setTime(d: Date) {
        this.currentTime = d;
    }
    constructor(h: number, m: number) {}
}

let c2 = new Clock2(1, 2);
c2.setTime(new Date);
console.log('date' + c2.currentTime)



/* 类静态部分与实例部分的区别 */
// 当你操作类和接口的时候，你要知道类是具有两个类型的： 静态部分的类型和实例的类型。
// 你会注意到，当你用构造器签名去定义一个接口并试图定义一个类去实现这个接口时会得到
// 一个错误：

interface ClockConstructor {
    // new (hour: number, minute: number);
}

class Clock3 implements ClockConstructor {
    currentTime: Date;
    constructor(h: number, m: number) {} 
    
    // err,因为当一个类实现一个接口时，只对其实例部分进行类型检查。constructor
    // 存在于类的静态部分，所以不在检查的范围。
}
// 因此， 我们应该直接操作类的静态部分。！！！
let c3 = new Clock3(1, 2);

// 定义两个接口，ClockConstructor为构造函数所用和ClockInterface为实例方法所用。
interface ClockConstructor4 {
    new (hour: number, minute: number): ClockInterface4;
}

interface ClockInterface4 {
    tick();
}

class DigitalClock implements ClockInterface4 {
    constructor(h: number, m: number) {}
    tick() {
        console.log('beep beep');
    }
}

class AnalogClock implements ClockInterface4 {
    constructor(h: number, m: number) {}
    tick() {
        console.log('tick tock');
    }
}

// 定义一个构造函数createClock,它用传入的类型创建实例。(类似一个中转的功能，不一般不建议在项目中使用，会增加代码的阅读难度)
function createClock4(ctor: ClockConstructor4, hour: number, minute: number): ClockInterface4 {
    return new ctor(hour, minute);
}


let digital = createClock4(DigitalClock, 12, 17);
digital.tick();
let analog = createClock4(AnalogClock, 7, 32);
analog.tick();

/* 继承接口 */
// 和类一样，接口也可以相互继承。
interface Shape {
    color: string;
}

interface Square extends Shape {
    sideLength: number;
}

let square = <Square>{}; // 声明一个空类 并做类型断言
square.color = 'blue'; // 类的属性赋值
square.sideLength = 10;


// 一个接口可以继承多个接口，创建出多个接口的合成接口
interface Shape {
    color: string;
}

interface PenStroke {
    penWidth: number;
}

interface Square extends Shape, PenStroke {
    sideLength: number;
}

let square2 = <Square>{};
square2.color = 'blue';
square2.sideLength = 10;
square2.penWidth = 5.0;



/* 混合类型 */
// 由于JavaScript其动态灵活的特点，有时你会希望一个对象以同时具有上面提到的多种类型


//example： 一个对象可以同时做为函数和对象使用，并带胡额外的属性。
interface Counter {
    (start: number): string; // （）是指类或方法的参数，返回string类型
    interval: number;
    reset(): void;
}

function getCounter(): Counter {
    let counter = <Counter>function (start: number) {};
    counter.interval = 123;
    counter.reset = function () {};
    return counter;
}


let c = getCounter();
c(10);
c.reset();
c.interval = 5.0;



/* 接口继承类 */
// 当接口继承了一个类类型时，它会继承类的成员但不包括其实现。
// 接口同样会继承到类的private和protected成员。这意味着当你创建了一个接口继承了一个
// 拥有私有或受保护的成员的类时，这个接口类型只能被这个类或其子类所实现(implement)。


// 基类
class Control { 
    private state: any;
}

// 接口继承基类
interface SelectableControl extends Control {
    select(): void;
}

// 类继承基类，并实现接口类
class Button extends Control implements SelectableControl {
    select() { }
}

// 类继承基类
class TextBox extends Control {
    select() { }
}


// 类实现接口
// err: state在基类中是私有属性，而Image类没有
// class Image implements SelectableControl {
//     select() { }
// }

