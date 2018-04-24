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
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
/**
 * 装饰器
 *
 * 装饰器是一种特殊的声明，它能够被附加到类声明，方法，访问符，属性或参数上。
 * 装饰器使用@expression这种形式，expression求值后必须为一个函数，它会在运行时
 * 被调用，被装饰的声明信息做为参数传入。
 */
// !!!这个文件的生成方法： tsc --target ES5 --experimentalDecorators  --outFile build\test018.js test018.ts
// example  @sealed
//  function sealed(target) {  }
// 装饰器工厂
// 装饰器工厂就是一个简单的函数，它返回一个表达式，以供装饰器在运行时调用。
function color(value) {
    return function (target) {
    };
}
// 装饰器组合
// 写在同一行上：
// @f @g x
// 写在多行上：
// @f
// @g
// x
/* 当多个装饰器应用于一个声明上，它们求值方式与复合函数相似。在这个模型下，当复合f和g时，
复合的结果(f g)(x)等同于f(g(x))

同样的，在TypeScript里，当多个装饰器应用在一个声明上时会进行如下步骤的操作：
1. 由上至下依次对装饰器表达式求值。
2. 求值的结果会被当作函数，由下至上依次调用。
 */
// 装饰器工厂的例子
function f() {
    console.log('f(): evaluated');
    return function (target, propertyKey, descriptor) {
        console.log('f(): called');
    };
}
function g() {
    console.log('g(): evaluated');
    return function (target, propertyKey, descriptor) {
        console.log('g(): called');
    };
}
var Ca = /** @class */ (function () {
    function Ca() {
    }
    Ca.prototype.method = function () { };
    __decorate([
        f(),
        g()
    ], Ca.prototype, "method", null);
    return Ca;
}());
/*
    装饰器求值

类中不同声明上的装饰器将按以下规定的顺序应用：
1. 参数装饰器，然后依次是方法装饰器，访问符装饰器，或属性装饰器应用到每个实例成员。
2. 参数装饰器，然后依次是方法装饰器，访问符装饰器，或属性装饰器应用到每个静态成员。
3. 参数装饰器应用到构造函数。
4. 类装饰器应用到类。
*/
/*
    类装饰器
类装饰器在类声明之前被声明（紧靠着类声明）。类装饰器应用于类构造函数，可以用来监视，修改或
替换类定义。类装饰器不能用在声明文件中（.d.ts),也不能用在任何外部上下文中（比如declare的类）。

类装饰器表达式会在运行时当作函数被调用，类的构造函数作为其唯一的参数。

如果类装饰器返回一个值，它会使用提供的构造函数来替换类的声明。

!! 注意，如果要返回一个新的构造函数，你必须注意处理好原来的原型链。在运行时的装饰器调用逻辑中
不会为此做这些。
*/
// 类装饰器(@sealed)的例子，应用在Greeter类
var Greetera = /** @class */ (function () {
    function Greetera(message) {
        this.greeting = message;
    }
    Greetera.prototype.greet = function () {
        return 'Hello, ' + this.greeting;
    };
    Greetera = __decorate([
        sealed
    ], Greetera);
    return Greetera;
}());
//　定义@sealed装饰器
function sealed(constructor) {
    Object.seal(constructor);
    Object.seal(constructor.prototype);
}
// 当@sealed被执行时，它将密封此类的构造函数和原型。
// 重载构造函数的例子
function classDecorator(constructor) {
    return /** @class */ (function (_super) {
        __extends(class_1, _super);
        function class_1() {
            var _this = _super !== null && _super.apply(this, arguments) || this;
            _this.newProperty = 'new property';
            _this.hello = 'override';
            return _this;
        }
        return class_1;
    }(constructor));
}
var Greeterb = /** @class */ (function () {
    function Greeterb(m) {
        this.property = 'property';
        this.hello = m;
    }
    Greeterb = __decorate([
        classDecorator
    ], Greeterb);
    return Greeterb;
}());
console.log(new Greeterb('world'));
