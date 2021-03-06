/**
 * 枚举
 * 
 * 使用枚举可以定义一些带名字的常量。
 * 使用枚举可以清晰地表达意图或创建一组有区别的用例。
 * TypeScript支持数字和基于字符串的枚举。
 */

 enum Direction {
     Up = 1,
     Down,
     Left,
     Right,
 }
 // Up使用初始化为1.其余的成员会从1开始增。也就是说
 // Direction.Up的值为1， Down为2， Left为3， Right为4.


 // 完全不使用初始化器：
 enum Direction2 {
     Up,
     Down,
     Left,
     Right,
 }
 // 现在Up的值为0, Down的值为1。

 // 使用枚举很简单: 通过枚举的属性来访问检举成员，和枚举的名字来访问枚举类型：
 enum Response1 {
     No = 0, 
     Yes = 1,
 }

 function respond(recipient: string, message: Response1): void {
     // ...
 }
 respond('Princess Caroline', Response1.Yes);


// 数字枚举可以被混入到计算过的和常量成员。简短地说，不带初始化器的枚举或者被放在
// 第一的位置，或者被放在使用了数字常量或其它常量初始化了的枚举后面
enum E {
    // A = getSomeValue(),
    B, // err: 'A' is not constant-initialized, so 'B' needs an initializer
}



/* 字符串枚举 */
// 在一个字符串枚举里，每个成员都必须用字符串字面量，或另外一个字符串枚举成员进行
// 初始化
enum Direction3 {
    Up = 'UP',
    Down = 'DOWN',
    Left = 'LEFT',
    Right = 'RIGHT',
}
// 由于字符串枚举没有自增长的行为，字符串枚举可以很好的序列化。换句话说，如果你正在
// 调试并且必须要读一个数字枚举的运行时的值，这个值通常是很难读的-它并不能表达有用的
// 信息（尽管反向映射可以），字符串枚举允许你提供一个运行时有意义的并且可读的值，独立
// 于枚举成员的名字。

// 不建议混合字符串和数字成员
enum BooleanLikeHeterogeneousEnum {
    No = 0,
    Yes = 'YES',
}



/* 计算的和常量成员 */
// 每个枚举成员都带有一个值，它可以是常量或计算出来的。当满足如下条件时，枚举成员被
// 当作是常量：
// 1. 它是枚举的第一个成员且没有初始化器，这种情况下它被赋予值0：
enum E1 { X }  // E.X = 0

// 2. 它不带有初始化器且它之前的枚举成员是一个数字常量。这种情况下，当前枚举成员的值
// 为它上一个枚举成员的值加1

enum E2 { X, Y, Z }

enum E3 {
    A = 1,
    B,
    C,
}

// 3. 枚举成员使用常量枚举表达式初始化。常数枚举表达式是TypeScript表达式的子集，它
// 可以在编译阶段求值。当一个表达式满足下面条件之一时，它就是一个常量枚举表达式：
/*
    3.1 一个枚举表达式字面量（主要是字符串字面量或数字字面量）
    3.2 一个对之前定义的常量枚举成员的引用（可以是在不同的枚举类型中定义的）
    3.3 带括号的常量枚举表达式
    3.4 一元运算符 +，-，~其中之一应用在了常量枚举表达式
    3.5 常量枚举表达式做为二元运算 +，- ，*，/, %, <<, >>, >>>, &, |, ^的操作对象。若
    常数枚举表达式求值后为NaN或Infinity, 则会在编译阶段报错。
*/

enum FileAccess {
    // constant members
    None,
    Read  = 1 << 1,
    Write = 1 << 2,
    ReadWrite = Read | Write,
    // Computed member
    G = '123'.length,
}



/* 联合枚举与枚举成员的类型 */
/*
任何字符串字面量(例如：'foo', 'bar', 'baz')
任何数字字面量（例如：1, 100）
应用于一元 -符号的数字字面量（例如：-1，-100）
*/

enum ShapeKind {
    Circle,
    Square,
}

interface Circle {
    kind: ShapeKind.Circle;
    radius: number;
}

interface Square {
    kind: ShapeKind.Square;
    sideLength: number;
}
let c: Circle = {
    // kind: ShapeKind.Square, // err kind: ShapeKind.Circle,
    kind: ShapeKind.Circle,
    radius: 100,
}

enum E4 {
    Foo,
    Bar,
}

function f(x: E4) {
    // if (x !== E4.Foo ||　x !== E4.Bar) {
    if (x !== E4.Foo) {
        // Err !==  cannot be applied to types 'E4.Foo' and 'E4.Bar'
    }
}



/* 运行时的枚举 */
enum E5 {
    X,
    Y,
    Z,
}

function f2(obj: { X: number }) {
    return obj.X;
}
f2(E5);


/* 反向映射 */

enum Enum {
    A,
}
let a = Enum.A;
let nameOfA = Enum[a]; // A

/*
这面TypeScript编译为JavaScript:
var Enum;
(function(Enum) {
    Enum[Enum["A"] = 0] = "A";
})(Enum || (Enum = {}));
var a = Enum.A;
var nameOfA = Enum[a]; // A

生成的代码中，枚举类型被编译成一个对象，它包含了正向映射（name->value）和反向
映射（value -> name）。引用枚举成员总会成为对属性访问并且永远也不会内联代码。
*/



/* const 枚举 */
const enum Enum2 {
    A = 1,
    B = A * 2,
}
// 常量枚举只能使用常量检举表达式，并且不同于常规的枚举，它们在编译阶段会被删除。
// 常量枚举成员在使用的地方会被内联进来。之所以可以这么做是因为，常量枚举不允许包含
// 计算成员。

const enum Directions {
    Up,
    Down,
    Left,
    Right,
}
let directions = [Directions.Up, Directions.Down, Directions.Left, Directions.Right];
/*
编译后：
var directions = [0 \/* Up *\/, 1 \/* Down *\/, 2 \/* Left *\/, 3 /* Right *\/];
*/


/* 外部枚举 */
declare enum Enum3 {
    A = 1,
    B, 
    C = 2,
}
// 外部枚举和非外部枚举之间有一个重要的区别，在正常的枚举里，没有初始化方法的成员被当成常
// 数成员。对于非常数的外部枚举而言，没有初始化方法时被当做需要经过计算的。