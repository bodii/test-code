/**
 * JSX
 * 
 * JSX是一种嵌入式的类似XML的语法。它可以被转换成合法的JavaScript，尽管转换的语义是依据
 * 不同的实现而定的。JSX因React框架而流行，但是也被其它应用所使用。TypeScript支持内嵌，
 * 类型检查和将JSX直接编译为JavaScript。
 */

 /*
     基本用法
    想使用JSX必须做两件事：
    1. 给文件一个.tsx扩展名
    2. 启用jsx选项

    TypeScript具有三种JSX模式：preserve, react和react-native.这些模式只在代码生成阶段起作用-
    类型检查并不受影响。在preserve模式下生成代码中会保留JSX以供后续的转换操作使用（例如：Babel）。
    另外，输出文件会带有.jsx扩展名。react模式会生成React.createElement,在使用前不需要再进行转换操作
    了，输出文件的扩展名为.js。react-native相当于preserve,它也保留了所有的JSX,但是输出文件的扩展名
    是.js。

    可以通过在命令行里使用 --jsx标记或tsconfig.json里的选项来指定模式。
    注意：React标识符是写死的硬编码，所以你必须保证React（大写的R)是可用的。
 */


 // as 操作符

 // 回想一下类型断言：
//  var foo = <foo>bar;
// TypeScript在.tsx文件里禁用了使用尖括号的类型断言。
// 为了弥补.tsx里的这个功能，新加入了一个类型断言符号： as。 上面的例子使用as操作符改写：
// var foo = bar as foo;
// as 操作符在.ts 和 .tsx里都可胳膊，并且与其它类型断言行为是等价的。

/*
固有元素

固有元素使用特殊的接口JSX.IntrinsicElements来查找。默认地。如果这个接口没有指定，会全部通过，
不对固有元素进行类型检查。然而，如果这个接口存在，那么固有元素的名字需要在JSX.IntrinsicElements
接口的属性里查找。
*/
declare namespace JSX {
    interface IntrinsicElements {
        foo: any
    }
}

// <foo />; // 正确
// <bar />; // 错误

// 上例中，<foo />没有问题，但是<bar />会报错，因为它没在JSX.IntrinsicElements里指定。

// !!! 你可以在JSX.IntrinsicElements上指定一个用来捕获所有字符串索引：
declare namespace JSX {
    interface IntrinsicElements {
        [elemName: string]: any;
    }
}


// 基于值的元素
// 基于值的元素会简单的在它所在的作用域里按标识符查找
// import MyComponent from './myComponent';

// <MyComponent />; // 正确
// <SomeOtherComponent />; // 错误
// 有两种方式可以定义基于值的元素
// 1. 无状态函数组件（SFC） 2. 类组件


// 无状态函数组件
interface FooProp {
    name: string;
    X: number;
    Y: number;
}

// declare function AnotherComponent(prop: {name: string});
// function ComponentFoo(prop: FooProp) {
//     return <AnotherComponent name=prop.name />;
// }
// const Button = (prop: {value: string}, context: { color: string }) => <button>


// 由于无状态函数组件是简单的JavaScript函数，所以可以利用函数重载
// interface ClickableProps {
//     children: JSX.Element[] |　JSX.Element
// }

// interface HomeProps extends ClickableProps {
//     home: JSX.Element;
// }

// interface SideProps extends ClickableProps {
//     side: JSX.Element |　string;
// }
// function MainButton(prop: HomeProps): JSX.Element;
// function MainButton(prop: SideProps): JSX.Element { }



/*
类组件

元素类的类型和元素实例的类型
现在

*/

/*
子孙类型检查
children是元素属性(attribute)类型的一个属性(property)。与使用JSX.ElementAttributesProperty
来决定props名类似，我们可以利用JSX.ElementChildrenAttribute来决定children名。
JSX.ElementChildrenAttribute应该被声明在单一的属性(property)里。
*/
declare namespace JSX {
    interface ElementChildrenAttribute {
        children: {};
    }
}


// JSX结果类型
/*
默认地JSX表达式结果的类型为any。你可以自定义这个类型，通过指定JSX.Element接口。
然而，不能够从接口里检索元素，属性或JSX的子元素的类型信息。它是一个黑盒。
*/

// 嵌入的表达式
// JSX允许你使用{}标签来内嵌表达式。
// var a = <div>
