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
 * 泛型
 *
 * 在像C#和Java这新的语言中，可以使用泛型来创建可重用的组件，一个组件可以支持
 * 多种类型的数据。
 */
function identity(arg) {
    return arg;
}
// 或者可以使用any类型来定义函数：
function identity2(arg) {
    return arg;
}
// 使返回值的类型与传入参数的类型相同,使用了类型变量，它是一种特殊的变量，
// 只用于表示类型而不是值。 // 泛型，因为它可以适用于多个类型。它不会丢失信息，像
// 第一个例子那样准确性，传入数值类型并返回数值类型。
function identity3(arg) {
    return arg;
}
var output = identity3('myString');
// 这里明确的指定T是string类型，并做为一个参数传给函数，使用了<>括起来而不是()。
// 第二种方法更普通。利用了类型推论--即编译器会根据传入的参数自动地帮助我们确定T的类型
var output2 = identity3('myString');
// 编译器可以查看myString的值，然后把T设置为它的类型。类型推论帮助我们保持代码精简和
// 高可读性。如果编译器不能够自动地推断出类型的话，只能像上面那样明确的传T的类型，在一
// 些复杂的情况下，这是可能出现的。
/* 使用泛型变量 */
function loggingIdentity(arg) {
    // console.log(arg.length); // err // 如T是数字时，没有length属性的
    return arg;
}
function loggingIdentity2(arg) {
    console.log(arg.length);
    return arg;
}
// Array 自定义泛型
function loggingIdentity3(arg) {
    console.log(arg.length);
    return arg;
}
/* 泛型类型 */
// 泛型函数的类型与非泛型函数的类型没有什么不同，只是有一个类型参数在最前面，像函数声明一样：
function identity4(arg) {
    return arg;
}
var myIdentity = identity4;
// 也可以使用不同的泛型参数名，只要在数量上和使用方式上能对应上就可以
function identity5(arg) {
    return arg;
}
var myIdentity2 = identity5;
// 可以使用带有调用签名的对象字面量来定义泛型函数
function identity6(arg) {
    return arg;
}
var myIdentity3 = identity6;
function identity7(arg) {
    return arg;
}
var myIdentity4 = identity7;
myIdentity4(5);
function identity8(arg) {
    return arg;
}
var myIdentity5 = identity8;
/* 泛型类 */
// 泛型类看上去与泛型接口差不多。泛型类使用(<>)括起泛型类型，跟在类名后面。
var GenericNumber = /** @class */ (function () {
    function GenericNumber() {
    }
    return GenericNumber;
}());
var myGenericNumer = new GenericNumber();
myGenericNumer.zeroValue = 0;
myGenericNumer.add = function (x, y) { return x + y; };
var myGenericString = new GenericNumber();
myGenericString.zeroValue = '';
myGenericString.add = function (x, y) { return x + y; };
console.log(myGenericString.add(myGenericString.zeroValue, 'test'));
function loggingIdentity4(arg) {
    console.log(arg.length);
    return arg;
}
// loggingIdentity4(3); // err
loggingIdentity4('a');
loggingIdentity4({ length: 10, value: 3 }); // 对象的话，必须包含length属性
loggingIdentity4([1, 2, 3]);
/* 在泛型约束中使用类型参数 */
// 你可似声明一个类型参数，且它被另一个类型参数所约束。
/*
err:
function getProperty(obj: T, key: K) {
    return obj[key];
}
let x = { a: 1, b: 2, c: 3, d: 4 };
getProperty(x, 'a');
*/
/* 在泛型里使用类类型 */
// 在TypeScript使用泛型创建工厂函数时，需要引用构造函数的类类型
function create(c) {
    return new c();
}
// 使用原型属性推断并约束构造函数与类实例的关系
var BeeKeeper = /** @class */ (function () {
    function BeeKeeper() {
    }
    return BeeKeeper;
}());
var ZooKeeper = /** @class */ (function () {
    function ZooKeeper() {
    }
    return ZooKeeper;
}());
var Animal = /** @class */ (function () {
    function Animal() {
    }
    return Animal;
}());
var Bee = /** @class */ (function (_super) {
    __extends(Bee, _super);
    function Bee() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return Bee;
}(Animal));
var Lion = /** @class */ (function (_super) {
    __extends(Lion, _super);
    function Lion() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return Lion;
}(Animal));
function createInstance(c) {
    return new c();
}
createInstance(Lion).keeper.nametag;
createInstance(Bee).keeper.hasMask;
