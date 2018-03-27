/**
 *  类型兼容性
 *
 * TypeScript里的类型兼容性是基于结构子类型的。
 * 结构类型是一种只使用其成员来描述类型的方式。它正好与名义类型形成对比
 * 基于名义类型的类型系统中，数据类型的兼容性或等价性是通过明确的声明和/或
 * 类型的名称来决定的。这与结构性类型系统不同，它是基于类型的组成结构，且不要求
 * 明确地声明。
 */
var Person = /** @class */ (function () {
    function Person() {
    }
    return Person;
}());
var p;
p = new Person();
var x;
var y = { name: 'Alice', location: 'Seattle' };
x = y;
/*
这里要检查y是否能赋值给x，编译器检查x中的每个属性，看是否能在y中也找到对应属性。
在这个例子中，y必须包含名字是name的string类型成员。y满足条件，因此赋值正确。
*/
function greet(n) {
    console.log('Hello, ' + n.name);
}
greet(y);
/*
注意，y有个额外的location属性，但这个不会引发错误。只有目标类型（这里Named2）
的成员会被一一检查是否兼容。
这个比较过程是递归进行的，检查每个成员及子成员。
*/
/* 比较两个函数 */
var x1 = function (a) { return 0; };
var y1 = function (b, s) { return 0; };
console.log(y1 = x1);
// console.log(x1 = y1); // err: x1能赋值给y1，但y1不能赋值给x
/*
x的每个参数必须在y里找到对应类型的参数。注意的是参数的名字相同与否无所谓，只看它们的类型。
这里，x的每个参数在y中都能找到对应的参数，所以允许赋值。
第二个赋值错误，因为y有个必需的第二个参数，但是x并没有，所以不允许赋值。
*/
var items = [1, 2, 3];
items.forEach(function (item, index, array) { return console.log(item); });
items.forEach(function (item) { return console.log(item); });
var x2 = function () { return ({ name: 'Alice' }); };
var y2 = function () { return ({ name: 'Alice', location: 'Seattle' }); };
x2 = y2;
// y2 = x2;  // err
/* 函数参数双向协变 */
/*
当比较函数参数类型时，只有当源函数参数能够赋值给目标函数或者反过来时才能赋值成功。
这是不稳定的，因为调用者可能传入了一个具有更精确类型信息的函数，但是调用这个传入的函数时
却使用了不是那么精确的类型信息。实际上，这极少会发生错误，并且能够实现很多JavaScript里
的常见模式。
 */
var EventType;
(function (EventType) {
    EventType[EventType["Mouse"] = 0] = "Mouse";
    EventType[EventType["Keyboard"] = 1] = "Keyboard";
})(EventType || (EventType = {}));
function listentEvent(eventType, handler) {
    /* ... */
}
listentEvent(EventType.Mouse, function (e) { return console.log(e.x3 + ',' + e.y3); });
listentEvent(EventType.Mouse, function (e) { return console.log(e.x3 + ',' + e.y3); });
listentEvent(EventType.Mouse, (function (e) { return console.log(e.x3 + ',' + e.y3); }));
// listentEvent(EventType.Mouse, (e: number) => console.log(e)); // err
/* 可选参数及剩余参数 */
// 比较函数兼容性的时候，可选参数与必须参数是可互换的。源类型上有额外的可选参数
// 不是错误，目标类型的可选参数在源类型里没有对应的参数也不是错误。
// 当一个函数有剩余参数时，它被当做无限个可选参数。
function invokeLater(args, callback) {
    /* ... Invoke callback with args ... */
}
invokeLater([1, 2], function (x, y) { return console.log(x + ', ' + y); });
invokeLater([1, 2], function (x, y) { return console.log(x + ', ' + y); });
// 函数重载
// 对于有重载的函数，源函数的每个重载都要在目标函数上找到对应的函数签名。这确保了目标函数
// 可以在所有源函数可调用的地方调用
/* 枚举
不同枚举类型之间是不兼容的。
*/
var Status;
(function (Status) {
    Status[Status["Ready"] = 0] = "Ready";
    Status[Status["Waiting"] = 1] = "Waiting";
})(Status || (Status = {}));
;
var Color;
(function (Color) {
    Color[Color["Red"] = 0] = "Red";
    Color[Color["Blue"] = 1] = "Blue";
    Color[Color["Green"] = 2] = "Green";
})(Color || (Color = {}));
;
var status1 = Status.Ready;
// status1 = Color.Green; // err
/* 类
类与对象字面量和接口差不多，但有一点不同：类有静态部分和实例部分的类型。
比较两个类类型的对象时，只有实例的成员会被比较。静态成员和构造函数不在比较的范围内。
*/
var Animal = /** @class */ (function () {
    function Animal(name, numFeet) {
    }
    return Animal;
}());
var Size = /** @class */ (function () {
    function Size(numFeet) {
    }
    return Size;
}());
var a;
var s;
a = s; // OK
s = a; // OK
var x3;
var y3;
x3 = y3; // ok
// x3 和 y3是兼容的，因为它们的结构使用类型参数时并没有不同。
// 对于没指定泛型类型的泛型参数时，会把所有泛型参数当成any比较。
var identity = function (x) {
    // ...
};
var reverse = function (y) {
    // ...
};
identity = reverse; // ok
/* 高级主题 */
// 子类型与赋值
/*
 在TypeScript里，有两种类型的兼容性：子类型与赋值。它们的不同点在于，赋值扩展了子类型兼容，
 允许给any赋值或从any取值和允许数字赋值给枚举类型，或枚举类型赋值给数字
 */
var Example;
(function (Example) {
    Example[Example["N"] = 0] = "N";
})(Example || (Example = {}));
var N1 = 1;
N1 = Example.N;
/*
    语言里的不同地方分别使用了它们之中的机制。实际上，类型兼容性是由赋值兼容性来控制的，
    即使在implements和exentds语句也不例外。
*/ 
