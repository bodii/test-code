var __extends = (this && this.__extends) || (function () {
    var extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
/**
 * 接口
 */
// TypeScript的核心原则之一是对值所具有的结进行类型检查。它有时被称为“鸭式辨型法”
// 或 “结构性子类型化”。在TypeScript里，接口的作用就是为这些类型命名和为你的代码
// 或第三方代码定义契约。
function printLabel(labelledObj) {
    console.log(labelledObj.label);
}
var myObj = { size: 10, label: 'Size 10 object' };
printLabel(myObj);
function printLabel2(labelledObj) {
    console.log(labelledObj.label);
}
var myObj2 = { size: 10, label: 'Size 10 Object' };
printLabel2(myObj2);
function createSquare(config) {
    var newSquare = { color: 'white', area: 100 };
    if (config.color) {
        newSquare.color = config.color;
    }
    if (config.width) {
        newSquare.area = config.width * config.width;
    }
    return newSquare;
}
var mySquare = createSquare({ color: 'black' });
console.log('mySquare:' + mySquare);
function createSquare2(config) {
    var newSquare = { color: 'white', area: 100 };
    // if (config.clor) {
    //     newSquare.color = config.clor;
    // } 
    // err
    if (config.width) {
        newSquare.area = config.width * config.width;
    }
    return newSquare;
}
var mySquare2 = createSquare2({ color: 'black' });
var p1 = { x: 10, y: 20 };
// p1.x = 5;  // err
// TypeScript 具有ReadonlyArray<T>类型，它与Array<T>相似，只是把所有可变方法去掉了，
// 因此可以确保数组创建后再也不能被修
var a = [1, 2, 3, 4];
var ro = a;
// ro[0] = 12; // err
// ro.push(5); // err
// ro.length = 100; // err
// a = ro; // err
// 上面，可以看到就算把整个ReadonlyArray赋值到一个普通数组也是不可以的。
// 但是你可以用类型断言重写
a = ro;
function createSquare3(config) {
    var newSquare = { color: 'white', area: 100 };
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
var mySquare3 = createSquare2({ width: 100, opacity: 0.5 });
// 最后一种跳过这些检查的方式， 就是将这个对象赋值给另一个变量
var squareOptions = { colour: 'red', widht: 100 };
var mySquare4 = createSquare3(squareOptions);
var mySearch;
mySearch = function (source, subString) {
    var result = source.search(subString);
    return result > -1;
};
// 对于函数类型的类型检查来说，函数的参数名不需要与接口里定义的名字相匹配。
var mySearch2;
mySearch2 = function (src, sub) {
    var result = src.search(sub);
    return result > -1;
};
var myArray;
myArray = ['Bob', 'Fred'];
var myStr = myArray[0];
// 共支持两种索引签名：字符串和数字。可以时使用两种类型的索引，但是数字索引
// 的返回值必须是字符串索引返回值类型的子类型。这是因为当使用number来索引时，
// JavaScript会将它转换成string然后再去索引对象。也就是说用100（一个number）
// 去索引等同于使用"100"（一个string）去索引，因此两者需要保持一致。
var Animal = /** @class */ (function () {
    function Animal() {
    }
    return Animal;
}());
var Dog = /** @class */ (function (_super) {
    __extends(Dog, _super);
    function Dog() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return Dog;
}(Animal));
var myArray2 = ['Alice', 'Bob'];
var Clock = /** @class */ (function () {
    function Clock(h, m) {
    }
    return Clock;
}());
var Clock2 = /** @class */ (function () {
    function Clock2(h, m) {
    }
    Clock2.prototype.setTime = function (d) {
        this.currentTime = d;
    };
    return Clock2;
}());
var c2 = new Clock2(1, 2);
c2.setTime(new Date);
console.log('date' + c2.currentTime);
var Clock3 = /** @class */ (function () {
    function Clock3(h, m) {
    }
    return Clock3;
}());
// 因此， 我们应该直接操作类的静态部分。！！！
var c3 = new Clock3(1, 2);
