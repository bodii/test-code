var isDone = false;
var decLiteral = 6;
var hexLiteral = 0xf00d;
var binaryLiteral = 10;
var octalLiteral = 484;
var names = "bob";
// name = "smith";
// let name: string = `Gene`;
var age = 37;
var sentence = "Hello, my name is " + names + "'\n\nI'll be " + (age + 1) + " years old next month.";
// 数组
// let list: number[] = [1, 2, 3];
// 第二种方式是使用数组泛型， Array<元素类型>
var list = [1, 2, 3];
// 元组 Tuple
// 元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同。
// 比如，你可以定义一对值分别为string和number类型的元组
var x;
x = ['hello', 10];
//x = [10, 'hello']; // err
console.log(x[0].substr(1));
// console.log(x[1].substr(1)) //err
x[3] = 'world';
console.log(x[5].toString());
// x[6] = true // err, 布尔不是（string | number）类型
// 枚举
// enum Color {Red, Green, Blue}
// let c: Color = Color.Green;
// 手动赋值
var Color;
(function (Color) {
    Color[Color["Red"] = 1] = "Red";
    Color[Color["Green"] = 2] = "Green";
    Color[Color["Blue"] = 4] = "Blue";
})(Color || (Color = {}));
var c = Color.Green;
var colorName = Color[2];
alert(colorName);
// Any
var notSure = 4;
notSure = 'maybe a string instead';
notSure = false;
notSure.ifItExists();
notSure.toFixed();
var prettySure = 4;
prettySure.toString();
var list1 = [1, true, 'free'];
list1[1] = 100;
// void
// 当一个函数没有返回值时，你通常会见到其返回类型是void
function warnUser() {
    alert('This is my warning message');
}
var unusable = undefined;
// null 和 undefined
var u = undefined;
var n = null;
// never
// never类型表示的是那些不存在的值的类型。
// 返回never的函数必须存在无法达到的终点
function error(message) {
    throw new Error(message);
}
// 推断的返回值类型为never
function fail() {
    return error('Something failed');
}
// 返回never的函数必须存在无法达到的终点
function infiniteLoop() {
    while (true) {
    }
}
// 类型断言
// 类型断言有两种形式，其一是“尖括号”语法：
var someValue = 'this is a string';
var strLength = someValue.length;
// 另一个为as语法
var someValue2 = "this is a string";
var strLength2 = someValue2.length;
// 作用域
// function sumMatrix(matrix: number[][]) {
//     var sum = 0;
//     for (var i = 0; i < matrix.length; i++) {
//         var currnetRow = matrix[i];
//         for (var i = 0; i < currnetRow.length; i++) {
//             sum += currnetRow[i];
//         }
//     }
//     return sum;
// }
// 里层的for循环会覆盖变量i
for (var i = 0; i < 10; i++) {
    setTimeout(function () { console.log(i); }, 100 * i);
}
