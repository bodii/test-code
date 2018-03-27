/**
 * 泛型
 * 
 * 在像C#和Java这新的语言中，可以使用泛型来创建可重用的组件，一个组件可以支持
 * 多种类型的数据。
 */

 function identity(arg: number): number {
     return arg;
 }

 // 或者可以使用any类型来定义函数：
 function identity2(arg: any): any {
     return arg;
 }

// 使返回值的类型与传入参数的类型相同,使用了类型变量，它是一种特殊的变量，
// 只用于表示类型而不是值。 // 泛型，因为它可以适用于多个类型。它不会丢失信息，像
// 第一个例子那样准确性，传入数值类型并返回数值类型。
function identity3<T>(arg: T): T {
    return arg;
}

let output = identity3<string>('myString');
// 这里明确的指定T是string类型，并做为一个参数传给函数，使用了<>括起来而不是()。

// 第二种方法更普通。利用了类型推论--即编译器会根据传入的参数自动地帮助我们确定T的类型
let output2 = identity3('myString');
// 编译器可以查看myString的值，然后把T设置为它的类型。类型推论帮助我们保持代码精简和
// 高可读性。如果编译器不能够自动地推断出类型的话，只能像上面那样明确的传T的类型，在一
// 些复杂的情况下，这是可能出现的。



/* 使用泛型变量 */
function loggingIdentity<T>(arg: T): T {
    // console.log(arg.length); // err // 如T是数字时，没有length属性的
    return arg;
}

function loggingIdentity2<T>(arg: T[]): T[] {
    console.log(arg.length);
    return arg;
}

// Array 自定义泛型
function loggingIdentity3<T>(arg: Array<T>): Array<T> {
    console.log(arg.length);
    return arg;
}



/* 泛型类型 */
// 泛型函数的类型与非泛型函数的类型没有什么不同，只是有一个类型参数在最前面，像函数声明一样：
function identity4<T> (arg: T): T {
    return arg;
}
let myIdentity: <T>(arg: T) => T = identity4;

// 也可以使用不同的泛型参数名，只要在数量上和使用方式上能对应上就可以
function identity5<T> (arg: T) : T {
    return arg;
}
let myIdentity2: <U>(arg: U) => U = identity5;

// 可以使用带有调用签名的对象字面量来定义泛型函数
function identity6<T> (arg: T): T {
    return arg;
}
let myIdentity3: {<T>(arg: T): T} = identity6;


/* 泛型接口 */
interface GenericIdentityFn {
    <T>(arg: T): T;
}

function identity7<T> (arg: T) : T {
    return arg;
}
let myIdentity4: GenericIdentityFn = identity7;
myIdentity4(5);

// 或这样
interface GenericIdentityFn2<T> {
    (arg: T) : T;
}
function identity8<T> (arg: T) : T {
    return arg;
}
let myIdentity5: GenericIdentityFn2<number> = identity8;



/* 泛型类 */
// 泛型类看上去与泛型接口差不多。泛型类使用(<>)括起泛型类型，跟在类名后面。
class GenericNumber<T> {
    zeroValue: T;
    add: (x: T, y: T) => T;
}

let myGenericNumer = new GenericNumber<number>();
myGenericNumer.zeroValue = 0;
myGenericNumer.add = function (x, y) { return x + y; };

let myGenericString = new GenericNumber<string>();
myGenericString.zeroValue = '';
myGenericString.add = function (x, y) { return x + y; };

console.log(myGenericString.add(myGenericString.zeroValue, 'test'));



/* 泛型约束 */
// 定义一个接口来描述约束条件。创建一个包含.length属性的接口，使用这个接口和extends关键字
// 来实现约束
interface Lengthwise {
    length: number;
}

function loggingIdentity4<T extends Lengthwise>(arg: T): T {
    console.log(arg.length);
    return arg;
}
// loggingIdentity4(3); // err
loggingIdentity4('a');
loggingIdentity4({length: 10, value: 3}); // 对象的话，必须包含length属性
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
function create<T>(c: {new(): T;}): T {
    return new c();
}

// 使用原型属性推断并约束构造函数与类实例的关系
class BeeKeeper {
    hasMask: boolean;
}
class ZooKeeper {
    nametag: string;
}
class Animal {
    numLegs: number;
}

class Bee extends Animal {
    keeper: BeeKeeper;
}
class Lion extends Animal {
    keeper: ZooKeeper;
}

function createInstance<A extends Animal>(c: new() => A): A {
    return new c();
}

createInstance(Lion).keeper.nametag;
createInstance(Bee).keeper.hasMask;
