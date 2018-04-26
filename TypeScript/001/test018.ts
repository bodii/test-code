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
function color(value: string) { // 这是一个装饰工厂
    return function (target) { // 这是装饰器

    }
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
    return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
        console.log('f(): called');
    }
}

function g() {
    console.log('g(): evaluated');
    return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
        console.log('g(): called');
    }
}

class Ca {
    @f()
    @g()
    method() {}
}


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
@sealed
class Greetera {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet () {
        return 'Hello, ' + this.greeting;
    }
}

//　定义@sealed装饰器
function sealed(constructor: Function) {
    Object.seal(constructor);
    Object.seal(constructor.prototype);
}
// 当@sealed被执行时，它将密封此类的构造函数和原型。


// 重载构造函数的例子
function classDecorator<T extends {new (...args:any[]): {}}>(constructor: T) {
    return class extends constructor {
        newProperty = 'new property';
        hello = 'override';
    }
}

@classDecorator
class Greeterb {
    property = 'property';
    hello: string;
    constructor(m: string) {
        this.hello = m;
    }
}

console.log(new Greeterb('world'));



/* 
    方法装饰器
方法装饰器声明在一个方法的声明之前（紧靠着方法声明）。它会被应用到方法的属性描述符上，可以用来监视，修改或者
替换方法定义。方法装饰器不能用在声明文件(.d.ts)，重载或者任何外部上下文（比如declare的类）中。

方法装饰器表达式会在运行时当作函数被调用，传入下列3个参数：
1. 对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
2. 成员的名字。
3. 成员的属性描述符。


!!! 如果代码输出目标版本小于ES5, 属性描述符将会是undefined。

如果方法装饰器返回一个值，它会被用作方法的属性描述符。
！ 注意 如果代码输出目标版本小于ES5返回值会被忽略。
 */

 class Greeterc {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }

    @enumerable(false)
    greet() {
        return 'Hello, ' + this.greeting;
    }
 }

 // enumerable装饰器
 function enumerable(value: boolean) {
     return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
         descriptor.enumerable = value;
     }
 }
// 这里的@enumerable（false)是一个装饰器工厂。当装饰器@enumerable(false)被调用时，它会修改属性
// 描述符的enumerable属性。


/*
    访问器装饰器

访问器装饰器声明在一个访问器的声明之前（紧靠着访问器声明）。访问器装饰器应用于访问器的属性描述符并且
可以用来监视，修改或替换一个访问器的定义。访问器装饰器不能用在声明文件中（.d.ts)，或者任何外部上下文
（比如declare的类）里。

注意： TypeScript不允许同时装饰一个成员的get和set访问器。取而代之的是，一个成员的所有装饰的必须应用
在文档顺序的第一个访问器上。这是因为，在装饰器应用于一个属性描述符时，它联合了get和set访问器，而不是
分开声明的。

访问器装饰器表达式会在运行时当作函数被调用，传入下列3个参数：
1. 对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
2. 成员的名字。
3. 成员的属性描述符。
 */

class Point2 {
    private _x: number;
    private _y: number;
    constructor(x: number, y: number) {
        this._x = x;
        this._y = y;
    }

    @configurable(false)
    get x1() { return this._x; }

    @configurable(false)
    get y1() { return this._y; }
}

// 定义@configurable装饰器：
function configurable(value: boolean) {
    return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
        descriptor.configurable = value;
    };
}



/* 
    属性装饰器

属性装饰器声明在一个属性声明之前（紧靠着属性声明）。属性装饰器不能用在声明文件中(.d.ts),
或者任何外部上下文（比如declare的类）里。

属性装饰器表达式会在运行时当作函数被调用，传入下列2个参数：
1. 对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
2. 成员的名字。

注意： 属性描述符不会做为参数传入属性装饰器，这与TypeScript是如何初始化属性装饰器的有关。
因为目前没有办法在定义一个原型对象的成员时描述一个实例属性，并且没办法监视或修改一个属性的
初始化方法。返回值也会被忽略。因此，属性描述符只能用来监视类中是否声明了某个名字的属性。
*/
/*
class Greeter {
    @format("Hello, %s")
    greeting: string;

    constructor(message: string) {
        this.greeting = message;
    }

    greet() {
        let formatString = getFormat(this, 'greeting');
        return formatString.replace("%s", this.greeting);
    }
}
*/
// 然后定义@format装饰器和getFormat函数：
/*
import "reflect-metadata";

const formatMetadataKey = Symbol("format");

function format(formatString: string) {
    return Reflect.metadata(formatMetadataKey, formatString);
}

function getFormat(target: any, propertyKey: string) {
    return Reflect.getMetadata(formatMetadataKey, target, propertyKey);
}
*/
/*
    这个@format("Hello, %%s")装饰器是个装饰器工厂。当@format("Hello, %s")被调用时，
    它添加一条这个属性的元数据，通过reflect-metadata库里Reflect.metadata函数。当getFormat
    被调用时，它读取格式的元数据。

    注意  这个例子需要使用reflect-metadata库。
*/


/*
    参数装饰器
    参数装饰器声明在个参数声明之前（紧靠着参数声明）。参数装饰器应用于类构造函数或方法声明。
    参数装饰器不能用在声明文件（.d.ts），重载或其它外部上下文（比如declare的类）里。

    参数装饰器表达式会在运行时当作函数被调用，传入下列3个参数：
    1. 对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
    2. 成员的名字。
    3. 参数在函数参数列表中的索引。

    注意： 参数装饰器只能用来监视一个方法的参数是否被传入。
*/

// 下例定义了参数装饰器（@required）并应用于Greeter类方法的一个参数：
/*
class Greeter02 {
    greeting: string;

    constructor(message: string) {
        this.greeting = message;
    }

    @validate
    greet(@required name:string) {
        return "Hello " + name + ", " + this.greeting;
    }
}
*/

// 然后我们使用下面的函数定义@required 和 @validate装饰器：
/*
import "reflect-metadata";

const requiredMetadataKey = Symbol("required");

function required(target: Object, propertyKey: string | symbol, parameterIndex: number) {
    let existingRequiredParameters: number[] = Reflect.getOwnMetadata(requiredMetadataKey, target, propertyKey) || [];
    existingRequiredParameters.push(parameterIndex);
    Reflect.defineMetadata(requiredMetadataKey, existingRequiredParameters, target, propertyKey);
}

function validate(target: any, propertyName: string, descriptor: TypedPropertyDescriptor<Function>) {
    let method = descriptor.value;
    descriptor.value = function () {
        let requiredParameters: number[] = Reflect.getOwnMetadata(requiredMetadataKey, target, propertyName);
        if (requiredParameters) {
            for (let parameterIndex of requiredParameters) {
                if (parameterIndex >= arguments.length || arguments[parameterIndex] === undefined) {
                    throw new Error("Missing required argument.");
                }
            }
        }

        return method.apply(this, arguments);
    }
}
*/

// @required装饰器添加了元数据实体把参记为必需的。@validate装饰器把greet方法包裹在一个函数里在调用
// 原先的函数前验证函数参数。
/*
import "reflect-metadata";
class Point3 {
    x: number;
    y: number;
}

class Line {
    private _p0: Point3;
    private _p1: Point3;

    @validate
    set p0(value: Point) { 
        this._p0 = value; 
    }
    get p0() {
        retrun this._p0;
    }

    @validate
    set p1(value: Point) {
        this._p1 = value;
    }
    get p1() {
        return this._p1;
    }
}

function validate<T>(target: any, propertyKey: string, descriptor: TypedPropertyDescriptor<T>) {
    let set = descriptor.set;
    descriptor.set = function (value: T) {
        let type = Reflect.getMetadata("design:type", target, propertyKey);
        if ( !(value instanceof type) ) {
            throw new TypeError('Invalid type.');
        }
        set(value);
    }
}

// TypeScript编译器可以通过@Reflect.metadata装饰器注入设计阶段的类型信息。你可以认为它相当于下面的TypeScript:
class Line {
    private _p0: Point;
    private _p1: Point;

    @validate
    @Reflect.metadata("design:type", Point)
    set p0(value: Point) { this._p0 = value; }
    get p0() { return this._p0; }

    @validate
    @Reflect.metadata("disign:type", Point)
    set p1(value:Point) { this._p1 = value; }
    get p1() { return this._p1; }
}
*/

// 注意 装饰器元数据是个实验性的特性，并且可能在以后的版本中发生破坏性的改变。