let isDone: boolean = false;

let decLiteral: number = 6;
let hexLiteral: number= 0xf00d;
let binaryLiteral: number = 0b1010;
let octalLiteral: number = 0o744;

let names: string = "bob";
// name = "smith";

// let name: string = `Gene`;
let age: number = 37;
let sentence: string = `Hello, my name is ${ names }'

I'll be ${ age + 1 } years old next month.`;

// 数组
// let list: number[] = [1, 2, 3];
// 第二种方式是使用数组泛型， Array<元素类型>
let list: Array<number> = [1, 2, 3];


// 元组 Tuple
// 元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同。
// 比如，你可以定义一对值分别为string和number类型的元组
let x: [string, number];
x = ['hello', 10];

//x = [10, 'hello']; // err

console.log(x[0].substr(1))
// console.log(x[1].substr(1)) //err

x[3] = 'world';
console.log(x[5].toString());
// x[6] = true // err, 布尔不是（string | number）类型


// 枚举
// enum Color {Red, Green, Blue}
// let c: Color = Color.Green;

// 手动赋值
enum Color {Red = 1, Green = 2, Blue = 4}
let c: Color = Color.Green;

let colorName: string = Color[2];
alert(colorName);



// Any
let notSure: any = 4;
notSure = 'maybe a string instead';
notSure = false;

notSure.ifItExists();
notSure.toFixed();

let prettySure: Object = 4;
prettySure.toString();

let list1: any[] = [1, true, 'free']
list1[1] = 100;


// void
// 当一个函数没有返回值时，你通常会见到其返回类型是void
function warnUser(): void {
    alert('This is my warning message');
}
let unusable: void = undefined;



// null 和 undefined
let u: undefined = undefined;
let n: null = null;



// never
// never类型表示的是那些不存在的值的类型。

// 返回never的函数必须存在无法达到的终点
function error(message: string): never {
    throw new Error(message)
}

// 推断的返回值类型为never
function fail() {
    return error('Something failed')
}

// 返回never的函数必须存在无法达到的终点
function infiniteLoop(): never {
    while (true) {

    }
}



// 类型断言
// 类型断言有两种形式，其一是“尖括号”语法：
let someValue: any = 'this is a string';
let strLength: number = (<string>someValue).length;
// 另一个为as语法
let someValue2: any = "this is a string";
let strLength2: number = (someValue2 as string).length;



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
    setTimeout(function() {console.log(i);}, 100 * i);
}