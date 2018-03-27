/**
 *  类型兼容性
 * 
 * TypeScript里的类型兼容性是基于结构子类型的。
 * 结构类型是一种只使用其成员来描述类型的方式。它正好与名义类型形成对比
 * 基于名义类型的类型系统中，数据类型的兼容性或等价性是通过明确的声明和/或
 * 类型的名称来决定的。这与结构性类型系统不同，它是基于类型的组成结构，且不要求
 * 明确地声明。
 */

 // example
 interface Named {
     name: string;
 }

 class Person {
     name: string;
 }

 let p: Named;
 p = new Person();

 // 在使用基于名义类型的语言，比如C#或Java中，这段代码会报错，因为Person类没有明确
 // 说明其实现了Named接口。
 // TypeScript的结构子类型是根据JavaScript代码的典型写法来设计的。因为JavaScript里
 // 广泛地使用匿名对象，例如函数表达式和对象字面量，所以使用结构类型系统来描述这些类型
 // 系统更好。


 // TypeScript结构化类型系统的基本规则是，如果x要兼容y,那么y至少具有与x相同的属性
 interface Named2{
     name: string;
 }
 let x: Named2;
 let y = { name: 'Alice', location:'Seattle' };
 x = y;

 /*
 这里要检查y是否能赋值给x，编译器检查x中的每个属性，看是否能在y中也找到对应属性。
 在这个例子中，y必须包含名字是name的string类型成员。y满足条件，因此赋值正确。
 */
function greet(n: Named2) {
    console.log('Hello, ' + n.name);
}
greet(y);
/*
注意，y有个额外的location属性，但这个不会引发错误。只有目标类型（这里Named2）
的成员会被一一检查是否兼容。
这个比较过程是递归进行的，检查每个成员及子成员。
*/


/* 比较两个函数 */
let x1 = (a: number) => 0;
let y1 = (b: number, s: string) => 0;

console.log(y1 = x1);
// console.log(x1 = y1); // err: x1能赋值给y1，但y1不能赋值给x
/*
x的每个参数必须在y里找到对应类型的参数。注意的是参数的名字相同与否无所谓，只看它们的类型。
这里，x的每个参数在y中都能找到对应的参数，所以允许赋值。
第二个赋值错误，因为y有个必需的第二个参数，但是x并没有，所以不允许赋值。
*/

let items = [1, 2, 3];
items.forEach((item, index, array) => console.log(item));
items.forEach((item) => console.log(item));

let x2 = () => ({name: 'Alice'});
let y2 = () => ({name: 'Alice', location: 'Seattle'});

x2 = y2;
// y2 = x2;  // err



/* 函数参数双向协变 */
/*
当比较函数参数类型时，只有当源函数参数能够赋值给目标函数或者反过来时才能赋值成功。
这是不稳定的，因为调用者可能传入了一个具有更精确类型信息的函数，但是调用这个传入的函数时
却使用了不是那么精确的类型信息。实际上，这极少会发生错误，并且能够实现很多JavaScript里
的常见模式。
 */

 enum EventType {
     Mouse,
     Keyboard,
}

interface Event { timestamp: number; }
interface MouseEvent extends Event { x3: number; y3: number; }
interface KeyEvent extends Event { keyCode: number; }

function listentEvent(eventType: EventType, handler: (n: Event) => void) {
    /* ... */
}

listentEvent(EventType.Mouse, (e: MouseEvent) => console.log(e.x3 + ',' + e.y3));

listentEvent(
    EventType.Mouse, 
    (e: Event) => console.log(
        (<MouseEvent>e).x3 + ',' + (<MouseEvent>e).y3
    )
);

listentEvent(
    EventType.Mouse,
    <(e: Event) => void>(
        (e: MouseEvent) => console.log(
            e.x3 + ',' + e.y3
        )
    )
);

// listentEvent(EventType.Mouse, (e: number) => console.log(e)); // err



/* 可选参数及剩余参数 */
// 比较函数兼容性的时候，可选参数与必须参数是可互换的。源类型上有额外的可选参数
// 不是错误，目标类型的可选参数在源类型里没有对应的参数也不是错误。
// 当一个函数有剩余参数时，它被当做无限个可选参数。

function invokeLater(args: any[], callback: (...args: any[]) => void) {
    /* ... Invoke callback with args ... */
}

invokeLater([1, 2], (x, y) => console.log(x + ', ' + y));

invokeLater([1, 2], (x?, y?) => console.log(x + ', ' + y));


// 函数重载
// 对于有重载的函数，源函数的每个重载都要在目标函数上找到对应的函数签名。这确保了目标函数
// 可以在所有源函数可调用的地方调用


/* 枚举 
不同枚举类型之间是不兼容的。
*/ 
enum Status { Ready, Waiting };
enum Color { Red, Blue, Green };

let status1 = Status.Ready;
// status1 = Color.Green; // err



/* 类
类与对象字面量和接口差不多，但有一点不同：类有静态部分和实例部分的类型。
比较两个类类型的对象时，只有实例的成员会被比较。静态成员和构造函数不在比较的范围内。
*/
class Animal {
    public feet: number;
    public constructor(name: string, numFeet: number) { }
}

class Size {
    public feet: number;
    public constructor(numFeet: number) { }
}

let a: Animal;
let s: Size;
a = s; // OK
s = a; // OK


/* 类的私有成员
私有成员会影响兼容性判断。当类的实例用来检查兼容时，如果目标类型包含一个私有成员，那么源类型
必须包含来自同一个类的这个私有成员。这允许子类赋值给父类，但是不能赋值给其它有同样类型的类
*/


/* 泛型
因为TypeScript是结构性的类型系统，类型参数只影响使用其做为类型一部分的结果类型。
*/
interface Empty<T> { }
let x3: Empty<number>;
let y3: Empty<string>;

x3 = y3; // ok
// x3 和 y3是兼容的，因为它们的结构使用类型参数时并没有不同。
// 对于没指定泛型类型的泛型参数时，会把所有泛型参数当成any比较。
let identity = function<T>(x: T): void {
    // ...
}

let reverse = function<T>(y: T): void {
    // ...
}

identity = reverse; // ok

/* 高级主题 */
// 子类型与赋值
/*
 在TypeScript里，有两种类型的兼容性：子类型与赋值。它们的不同点在于，赋值扩展了子类型兼容，
 允许给any赋值或从any取值和允许数字赋值给枚举类型，或枚举类型赋值给数字
 */

enum Example {
    N,
}

let N1: number = 1;
N1 = Example.N;

/*
    语言里的不同地方分别使用了它们之中的机制。实际上，类型兼容性是由赋值兼容性来控制的，
    即使在implements和exentds语句也不例外。
*/