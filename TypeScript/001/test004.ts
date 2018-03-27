/**
 * 类
 * 
 * 传统的JavaScript程序使用函数和基于原型的继承来创建可重用的组件，但对于熟悉使用面向对象
 * 方式的程序员来讲就有些棘手，因为他们用的是基于类的继承并且对象是由类构建出来的。
 * 
 * 从ECMAScript2015，也就是ECMAScript6开始，JavaScript程序员将能够使用基于类的面向对象
 * 方式。使用TypeScript,允许开发者现在就使用这些特性，并且编译成JavaScript。
 */

// example
class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return 'Hello, ' + this.greeting;
    }
}

let greeter1 = new Greeter('world');


/* 继承 */
class Animal {
    move(distanceInMeters: number = 0) {
        console.log(`Animal moved ${distanceInMeters}m.`);
    }
}

class Dog extends Animal {
    bark() {
        console.log('Woof! Woof!');
    }
}

const dog = new Dog();
dog.bark();
dog.move(10);
dog.bark();



class Animal2 {
    name: string;
    constructor(theName: string) { this.name = theName; }
    move(distanceInMeters: number = 0 ) {
        console.log(`${ this.name } moved ${ distanceInMeters }m.`);
    }
}

class Snake extends Animal2 {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 5) {
        console.log('Slithering...');
        super.move(distanceInMeters);
    }
}

class Horse extends Animal2 {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 45) {
        console.log('Galloping...');
        super.move(distanceInMeters);
    }
}

let sam = new Snake('Sammy the Python');
let tom: Animal2 = new Horse('Tommy the Palomino');

sam.move();
tom.move(34);


/* 公共， 私有与受保护的修饰符 */
// 默认为public
class Animal3 {
    public name: string;
    public constructor(theName: string) { this.name = theName; }
    public move(distanceInMeters: number) {
        console.log(`${ this.name } moved ${ distanceInMeters }m.`);
    }
}


// private 
class Animal4 {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}
// new Animal4('Cat').name; // err: 'name' 是私有的。

class Animal5 {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

class Rhino extends Animal5 {
    constructor() { super('Rhino'); }
}

class Employee {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

let animal = new Animal5('Goat');
let rhino = new Rhino();
let employee = new Employee('Bob');

animal = rhino;
// animal = employee; // err: Animal5 与 Employee 不兼容


// protected
class Person {
    protected name: string;
    constructor(name: string) { this.name = name; }
}

class Employee2 extends Person {
    private department: string;

    constructor(name: string, department: string) {
        super(name);
        this.department = department;
    }

    public getElevatorPitch() {
        return `Hello, my name is ${ this.name } and I work in ${ this.department }.`;
    }
}

let howard = new Employee2('Howard', 'Sales');
console.log(howard.getElevatorPitch());
// console.log(howard.name); // err: name is private type


// example
class Person3 {
    protected name: string;
    protected constructor(theName: string) { this.name = theName; }
}

class Employee3 extends Person3 {
    private department: string;

    constructor(name: string, department: string) {
        super(name);
        this.department = department;
    }

    public getElevatorPitch() {
        return `Hello, my name is ${ this.name } and I work in ${ this.department }.`;
    }
}

let howard3 = new Employee3('Howard', 'Sales');
// let john = new Persion3('John'); // err: Persion3的构造函数是被保护的。



/* readonly修饰符 */
// 只读属性必须在声明时或构造函数里被初始化
class Octopus {
    readonly name: string;
    readonly numberOfLegs: number = 8;

    constructor (theName: string) {
        this.name = theName;
    }
}

let dad = new Octopus('Man with the 8 strong legs');
// dad.name = 'Man with the 3-piece suit'; // err: name是只读。



// 参数属性
class Animal6 {
    constructor(private name: string) {} // 这里舍弃了theName
    // 在构造函数里使用private name: string参数来创建和初始化name成员。
    move(distanceInMeters: number) {
        console.log(`${ this.name } moved ${ distanceInMeters }m.`);
    }
}
// 参数属性通过给构造函数参数添加一个访问限定符来声明。使用private限定一个参数属性
// 会声明并初始一个私有成员; 对于public和protected来说也是一样。



/* 存取器 */
// TypeScript支持通过getters/setters来截取对对象成员的访问。

class Employee4 {
    fullName: string;
}

let employee4 = new Employee4();
employee4.fullName = 'Bob Smith';
if (employee4.fullName) {
    console.log(employee4.fullName);
}

// next example
let passcode = 'secret passcode';

class Employee5 {
    private _fullName: string;

    get fullName(): string {
        return this._fullName;
    }

    set fullName(newName: string) {
        if (passcode && passcode == 'secret passcode') {
            this._fullName = newName;
        }
        else {
            console.log('Error: Unauthorized update of employee!');
        }
    }
}

let employee5 = new Employee5();
employee5.fullName = 'Bob Smith';
if (employee5.fullName) {
    console.log(employee5.fullName);
}

/*
对于存取器有下面几点需要注意：
1. 存取器要求你将编译器设置为输出ECMAScript5或更高。不支持降级到ECMAScript3。
2. 只带有get不带有set的存取器自动被推断为readonly。这在从代码生成 .d.ts文件时是
有帮助的，因为利用这个属性的用户会看到不允许改变它的值。
*/



/**
 * 静态属性
 * 
 * 使用static定义。
 * 访问 在类内和类外部都用 [类名.属性名]，类外也用实例对象.属性名
 */

 class Grid {
     static origin: {x: number, y: number} = {x: 0, y: 0};
     
     constructor (public scale: number) { } // scale  是参数属性

     calculateDistanceFromOrigin(point: {x: number; y: number;}) {
         let xDist = (point.x - Grid.origin.x); // 只能用类名.属性名访问
         let yDist = (point.y - Grid.origin.y);
         return Math.sqrt(xDist * xDist + yDist * yDist) / this.scale;
     }
 }

 let grid1 = new Grid(1.0);
 let grid2 = new Grid(5.0);


 console.log('grid1:', grid1.calculateDistanceFromOrigin({x: 10, y: 10}));
 console.log('grid2:', grid2.calculateDistanceFromOrigin({x: 10, y: 10}));



 /**
  * 抽象类
  * 
  * 抽象类做为其它派生类的基类使用。它们一般不会直接被实例化。
  * 不同于接口，抽象类可以包含成员的实现细节。
  * [ abstract ] 关键字是用于定义抽象类和在抽象类内部定义抽象方法。
  */

abstract class Animal7 {
    abstract makeSound(): void;
    move(): void {
        console.log('roaming the earch...');
    }
}

/*
    抽象类中的抽象方法不包含具体实现并且必须在派生类中实现。抽象方法的语法与接口方法
    相似。两者都是定义方法签名但不包含方法体。然而，抽象方法必须包含abstract关键字并
    且可以包含访问修饰符。
*/

abstract class Department {

    constructor(public name: string) {}

    printName(): void {
        console.log('Department name: ' + this.name);
    }

    abstract printMeeting(): void; // 必须在派生类中实现
}

class AccountingDepartment extends Department {
    constructor () {
        super('Accounting and Auditing');
    }

    printMeeting(): void {
        console.log('The Accounting Department meets each Monday at 10am.');
    }

    generateReports(): void {
        console.log('Generating accounting reports...');
    }
}

// 创建一个抽象类型的引用
let department: Department;

// department = new Department(); // err: 不能创建一个抽象类的实例。
department = new AccountingDepartment(); // 对一个抽象子类进行实例化
department.printName();
department.printMeeting();
// department.generateReports(); // err: 方法在声明的抽象类中不存在



/* 构造函数 */
// 当你在TypeScript里声明一个类时，实际上同时声明了很多东西。首先就是类的实例的类型
class Greeter2 {
    greeting: string;

    constructor(message: string) {
        this.greeting = message;
    }

    greet() {
        return 'Hollo, ' + this.greeting;
    }
}

let greeter2: Greeter2; // Greeter类的实例的类型是Greeter。
greeter2 = new Greeter2('world');
console.log(greeter2.greet());


let Greeter3 = (function () {
    function Greeter(message) {
        this.greeting = message;
    }
    Greeter.prototype.greet = function () {
        return 'Hello, ' + this.greeting;
    }
    return Greeter;
})();

let greeter3;
greeter3 = new Greeter3('world');
console.log(greeter3.greet());

// 改写版
class Greeter4 {
    static standardGreeting = 'Hello, there';
    greeting: string;
    greet () {
        if (this.greeting) {
            return 'Hello, ' + this.greeting;
        }
        else {
            return Greeter4.standardGreeting;
        }
    }
}

let greeter4: Greeter4;
greeter4 = new Greeter4();
console.log(greeter4.greet());

//　typeof Greeter4, 意思是取Greeter4类的类型，而不是实例的类型。或者更确切的说，
// “告诉我Greeter4标识符的类型”， 也就是构造函数的类型。这个类型包含了类的所有静
// 态成员和构造函数。之后，就和前面一样，我们在greeterMaker上使用new,创建实例。
let greeterMaker: typeof Greeter4 = Greeter4;
greeterMaker.standardGreeting = 'Hey there!';

let greeter5: Greeter4 = new greeterMaker();
// 等价于：
// let greeter5: Greeter4;
// greeter5 = new greeterMaker();
console.log(greeter5.greet());


/* 把类当做接口使用 */
class Point {
    x: number;
    y: number;
}

interface Point3d extends Point {
    z: number;
}

let point3d: Point3d = {x: 1, y: 2, z: 3};

/* 接口可以继承接口 */
interface Point2 {
    x: number;
    y: number;
}

interface Point3d2 extends Point2 {
    z: number;
}

let point3d2 : Point3d2 = {x: 1, y: 2, z: 3};