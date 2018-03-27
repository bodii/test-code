/**
 * 类型推论
 *
 * 即，类型是在哪里如何被推断的。
 */
var x = 3;
/*
变量x的类型被推断为数字。这种推断发生在初始化变量和成员，设置默认参数值和决定函数返回
值时。大多数情况下，类型推论是直截了当地。
 */
/* 最佳通用类型 */
// 当需要从几个表达式中推断类型时，会使用这些表达式的类型来推断出一个最合适的通用类型。
var x1 = [0, 1, null];
/*
为了推断x的类型，我们必须考虑所有元素的类型。这里有两种选择:number 和 null。计算通过类型
算法会考虑所有的候选类型，并给出一个兼容所有候选类型的类型。
由于最终的通用类型取自候选类型，有些时候候选类型共享相同的通用类型，但是却没有一个类型能做
为所有候选类型的类型。
let zoo = [new Rhino(), new Elephant(), new Snake()];

这里，我们想让zoo被推断为Animal[]类型，但是这个数组没有对象是Animal类型的，因此不能推断
出这个结果。为了更正，当候选类型不能使用的时候我们需要明确的指出类型：
let zoo: Animal[] = [new Rhino(), new Elephant(), new Snake()];

如果没有找到最佳通用类型的话，类型推断的结果为联合数组类型，(Rhino | Elephant | Snake)[]。
*/
/* 上下文类型 */
// TypeScript类型推论也可能按照相反的方向进行。这就叫做"按上下文归类"。
// 按上下文归类会发生在表达式的类型与所处的位置相关。
window.onmousedown = function (mouseEvent) {
    console.log(mouseEvent.button); // err
};
// 这个函数表达式有明确的参数类型注解，上下文类型被忽略。这样的话就不报错，因为这不会使用
// 上下文类型。
window.onmousedown = function (mouseEvent) {
    console.log(mouseEvent.button);
};
// 上下文归类会在很多情况下使用到。通常包含函数的参数，赋值表达式的右边，类型断言，对象成
// 员和数组字面量和返回值语句。上下文类型也会做为最佳通用类型的候选类型
/*function createZoo(): Animal[] {
    return [new Rhino(), new Elephant(), new Snake()];
}*/ 
