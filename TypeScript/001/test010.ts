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



/*
    可以为null的类型
    TypeScript具有两种特殊的类型，null和undefined, 它们分别具有值null和undefined。
    默认情况下，类型检查器认为null与undefined可以赋值给任何类型。null与undefined是
    所有其它类型的一个有效值。
 */
// --strictNullChecks标记可以解决此错误：当你声明一个变量时，它不会自动地包含null和
// undefined。

let s = 'foo';
// s = null; // err, null不能赋值给string
let sn: string | null = 'bar';
sn = null; // ok
// sn = undefined; // err;

// !!! TypeScript会把null和undefined区别对待。 string | null, string | undefined 和
// string | null | undefined是不同的类型。
// 使用了 --strictNullChecks 可选参数会被自动地加上 | undefined

function f(x: number, y?: number) {
    return x + (y || 0);
}

f(1, 2);
f(1);
f(1, undefined);
// f(1, null); // error



/*
    类型保护和类型断言
    由于可以为null的类型是通过联合类型实现，需要使用类型保护去除null。
*/
function f2(sn: string | null): string {
    if (sn == null) {
        return 'default';
    }
    else {
        return sn;
    }
}

// 或
function f3(sn: string | null): string {
    return sn || 'default';
}


// 如果编译器不能去除null和undefined，可以使用类型断言手动去除。语法是添加!后缀：
// identifier!从identifier的类型里去除null和nudefined
// function broken(name: string | null): string {
//     function postfix(epithet: string) {
//         return name.charAt(0) + '. the' + epithet; // err name is possibly null
//     }
//     name = name || 'Bob';
//     return postfix('great');
// }

function fixed(name: string | null): string {
    function postfix(epithet: string) {
        return name!.charAt(0) + ' .the '+ epithet; // ok
    }
    name = name || 'Bob';
    return postfix('great');
}


/* 类型别名
类型别名会给一个类型起个新名字，类型别名有时和接口很像，但是可以作用于原始值，联合类型
*/
type Name = string;
type NameResolver = () => string;
type NameOrResolver = Name | NameResolver;
function getName(n: NameOrResolver): Name {
    if (typeof n === 'string') {
        return n;
    }
    else {
        return n();
    }
}

// 同接口一样，类型别名也可以是泛型
type Container<T> = { value: T };
// 可以使用类型别名来属性里引用自己
type Tree<T> = {
    value: T;
    left: Tree<T>;
    right: Tree<T>;
}
// 与交叉类型一起使用，创建出一些十分稀奇古怪的类型
type LinkedList<T> = T & { next: LinkedList<T> };
interface Person {
    name: string;
}
var people: LinkedList<Person>;
var s0 = people.name;
var s1 = people.next.name;
var s2 = people.next.next.next.next.name;


// 然而， 类型别名不能出现在声明右侧的任何地方。
// type Yikes = Array<Yikes>; // err



/* 接口 vs 类型别名
类型别名和接口一样，但仍一些细微差别：
    1. 接口创建了一个新的名字，可以在其它任何地方使用。类型别名并不创建新名字 
    2. 类型别名不能被extends 和implements（自已也不能extends和implements其它类型）
    3. 如果无法通过接口来描述一个类型，并且需要使用联合类型或元组类型，这时通常会使用类型别名
 */
type Alias = { name: number }
interface Interface {
    num: number;
}
declare function aliased(arg: Alias): Alias;
declare function interface(arg: Interface): Interface;



/*
    字符串字面量类型
    字符串字面量类型允许你指定字符串必须的固定值。在实际应用中，字符串字面量类型可以与
    联合类型、类型保护和类型别名很好的配合。通过结合使用这些特性，你可以实现类似枚举类型
    的字符串。
*/
type Easing = 'ease-in' | 'ease-out' | 'ease-in-out';
class UIElement {
    animate(dx: number, dy: number, easing: Easing) {
        if (easing === 'ease-in') {
            // ...
        }
        else if (easing === 'ease-out') {

        }
        else if (easing === 'ease-in-out') {

        }
        else {
            // erro!
        }
    }
}

let button = new UIElement();
button.animate(0, 0 , 'ease-in'); // ok
// button.animate(0, 0, 'uneasy'); // err: 'uneasy' is not allowed here

// 字符串字面量类型还可以用于区分函数重载：
// function createElement(tagName: 'img'): HTMLImageElement;
// function createElement(tagName: 'input'): HTMLInputElement;
// function createElement(tagName: string): Element {}


// 数字字面量类型
function rollDie(): 1 | 2 | 3 | 4 | 5 | 6 {
    return 1;
}
// 很少这样使用，但它们可以用在缩小范围调试bug的时候：
// function foo(x: number) {
//     if ( x !== 1 || x !== 2 ) { // 非法的比较检查
        
//     }
// }

// 枚举成员类型
// 当每个枚举成员都是用字面量初始化的时候枚举成员是具有类型的。



/*
    可辨识联合

    可以合并单例类型，联合类型，类型保护和类型别名来创建一个叫做“可辨识联合的高级模式”
    也称做“标签联合” 或 “代数数据类型”。
    可辨识联合在函数式编程很有用处。
    它具有3个要素：
        1. 具有普通的单例类型属性 -- 可辨识的特征
        2. 一个类型别名包含了那些类型的联合 -- 联合
        3. 此属性上的类型保护
 */


// 声明将要联合的接口
interface Square {
    kind: 'square';
    size: number;
}
interface Rectangle {
    kind: 'rectangle';
    width: number;
    height: number;
}
interface Circle {
    kind: 'circle';
    radius: number;
}
// 每个接口都有kind属性但有不同的字符串字面量类型。
// kind属性称做 “可辨识的特征” 或 “标签”。其它的属性则特定于各个接口
type Shape2 = Square | Rectangle | Circle;
//使用 可辨识联合
function area(s: Shape2) {
    switch (s.kind) {
        case 'square': return s.size * s.size;
        case 'rectangle': return s.height * s.width;
        case 'circle': return Math.PI * s.radius ** 2;
    }
}

// 完整性检查
/*
    function area(s: Shape) : number { // 这里s可能不存在，所以默认返回可能为undefined
        switch(s) {
            case '':....
        }
    }
*/
// 当switch没有包涵所有情况，所以TypeScript认为这个函数有时会返回undefined


function assertNever(x: never): never { // never 的函数必须有无法被执行到的终止点
    throw new Error('Unexpected object: ' + x);
}

function area2(s: Shape2) {
    switch (s.kind) {
        case 'square': return s.size * s.size;
        case 'rectangle': return s.height * s.width;
        case 'circle': return Math.PI * s.radius ** 2;
        default: return assertNever(s); 
    }
}
// 这里，assertNever检查s是否为never类型--即为除去所有可能情况后剩下的类型


/*
    多态的this类型
    多态的this类型表示的是某个包含类或接口的 子类型。这被称为 F-bounded多态。它能很容
    易的表现连贯接口间的继承。
*/

class BasicCalculator {
    public constructor(protected value: number = 0) {}
    public currentValue(): number {
        return this.value;
    }
    public add(operand: number): this {
        this.value += operand;
        return this;
    }
    public multiply(operand: number): this {
        this.value *= operand;
        return this;
    }
}

let v = new BasicCalculator(2)
            .multiply(5)
            .add(1)
            .currentValue();

console.log(v);


// 由于这个类使用了this类型，你可以继承它，新的类可以直接使用之前的方法，不需要做任何改变
class ScientificCalculator extends BasicCalculator {
    public constructor(value = 0) {
        super(value);
    }
    public sin(): this {
        this.value = Math.sin(this.value);
        return this;
    }
}

let v2 = new ScientificCalculator(2)
            .multiply(5)
            .sin()
            .add(1)
            .currentValue();

console.log(v2);



/* 索引类型

JavaScript模式是从对象中选取属性的子集
function pluck(o, names) {
    return names.map(n => o[n]);
}

TypeScript，通过索引类型查询和索引访问操作符
*/
function pluck<T, K extends keyof T>(o: T, names: K[]): T[K][] {
    return names.map(n => o[n]);
}

interface Person {
    name: string;
    age: number;
}

let person: Person = {
    name: 'Jarid',
    age: 35
};
let strings: string[] = pluck(person, ['name']);
console.log(strings);
// keyof T, 索引类型查询操作符。对于任何类型T, keyof T的结果为T上已知的公共属性名的联合。
let personProps: keyof Person; // 'name' | 'age'
// keyof Person是完全可以与'name' | 'age'互相替换的。不同的是如果你添加了其它的属性到Person,
// 例如：address: string, 那么keyof Person会自动变为'name' | 'age' | 'address'。
// 你可以像pluck函数这类上下文里使用keyof，因为在使用之前你并不清楚可能出现的属性名。

function getProperty<T, K extends keyof T>(o: T, name: K) : T[K] {
    return o[name];
}

let name1: string = getProperty(person, 'name');
let age: number = getProperty(person, 'age');
// let unknown = getProperty(person, 'unknown'); // error 'unknown' is not in keys


/*
索引类型和字符串索引签名
keyof 和 T[K]与字符串索引签名进行交互。如果有一个字符串索引签名的类型，那么keyof T会是string.
并且T[string]为索引签名的类型。
 */
interface Maps<T> {
    [key: string]: T;
}
let keys: keyof Maps<number>; // string
let value: Maps<number>['foo']; // number


/*
    映射类型
    将一个已知的类型每个属性都变为可选
*/
interface PersonPartial {
    name?: string;
    age?: number; 
}
// 或者一个只读版本
interface PersonReadonly {
    readonly name: string;
    readonly age: number;
}

// TypeScript提供了从旧类型中创建新类型的一种方法 -- 映射类型
// 新类型以相同的形式去转换旧类型里每个属性。
type Readonly1<T> = {
    readonly [p in keyof T]: T[p];
}
type Partial1<T> = {
    [p in keyof T]?: T[p];
}

type PersonPartial2 = Partial1<Person>;
type ReadonlyPerson = Readonly1<Person>;

type Keys = 'option1' | 'option2';
type Flags = { [k in Keys]: boolean };

/*
它的语法与索引签名的方法类型，内部使用了for...in.
1. 类型变量k, 它会依次绑定每个属性
2. 字符串字面量联合的Keys，它包含了要迭代的属性名的集合。
3. 属性的结果类型
 */

// 等价于
type Flags1 = {
    option1: boolean;
    option2: boolean;
}

type NullablePerson = { [P in keyof Person]: Person[P] | null }
type PartialPerson2 = { [p in keyof Person]?: Person[p] }

type Nullable<T> = { [P in keyof T]: T[P]  | null}
type Partial3<T> = { [p in keyof T]?: T[p] }


type Proxy<T> = {
    get(): T;
    set(value: T): void;
}
type Proxify<T> = {
    [p in keyof T]: Proxify<T[p]>;
}
// function proxify2<T>(o: T): Proxify<T> {
//     return;
// }
// let proxyProps = proxify2(props);


/**
 Readonly<T> 和 Partial<T>用处非常大，因此它们与Pick和Record一同被包含进了TypeScript的
 标准库：
 */
type Pick2<T, K extends keyof T> = {
    [p in K]: T[p];
}
type Record2<K extends string, T> = {
    [p in K]: T;
}

// Readonly, Partial和Pick是同态的，但Record不是。因为Record并不需要输入类型来拷贝属性
// 非同态类型本质上会创建新的属性，因此它们不会从它处拷贝属性修饰符。
type ThreeStringProps = Record<'prop1' | 'prop2' | 'prop3', string>

/*
由映射类型进行推断
*/
function unproxify<T>(t: Proxy<T>): T {
    let result = {} as T;
    for (const k in t) {
        result[k] = t[k].get();
    }
    return result;
}

// let originalProps = unproxify(proxyProps);