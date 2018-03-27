/**
 * 变量的声明
 */
var __assign = (this && this.__assign) || Object.assign || function(t) {
    for (var s, i = 1, n = arguments.length; i < n; i++) {
        s = arguments[i];
        for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
            t[p] = s[p];
    }
    return t;
};
var __rest = (this && this.__rest) || function (s, e) {
    var t = {};
    for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p) && e.indexOf(p) < 0)
        t[p] = s[p];
    if (s != null && typeof Object.getOwnPropertySymbols === "function")
        for (var i = 0, p = Object.getOwnPropertySymbols(s); i < p.length; i++) if (e.indexOf(p[i]) < 0)
            t[p[i]] = s[p[i]];
    return t;
};
/* let 声明 */
var hello = 'Hello!';
/* 块作用域 */
function f(input) {
    var a = 100;
    if (input) {
        var b = a + 1;
        return b;
    }
    // return b; // err
}
/*
    拥有块级作用域的变量的另一个特点是，它们不能在被声明之前读或写。
    虽然这些变量始终“存在”于它们的作用域里，但在直到声明它的代码之
    前的域都属于 暂时性死区。它只是用来说明我们不能在let语句之前访问
    它们，幸运的是TypeScript可以告诉我们这些信息
    
*/
try {
    throw 'oh no!';
}
catch (e) {
    console.log('Oh well.');
}
// console.log(e);
// function f3() {
//     return a;
// }
// let a;
// f3(); // 如果之前没有声明a, 就会报错
// 使用var声明时，它不在乎你声明多少次;你只会得到1个。
function f4(x) {
    var x;
    var x;
    if (true) {
        var x;
    }
}
var x = 10;
// let x = 20; // err, 不能在1个作用域里多次声明1个变量。
// 并不是要求两个均是块级作用域的声明TypeScript才会给出一个错误的警告
// function f5(x) {
//     let x = 100; // err
// }
function g() {
    var x = 100;
    // var x = 100; // err
}
// 并不是说块级作用域变量不能用函数作用域变量来声明。而是块级作用域变量
// 域变量需要在明显不同的块里声明
function f6(condition, x) {
    if (condition) {
        var x_1 = 100;
        return x_1;
    }
    return x;
}
console.log(f6(false, 0)); // return 0
console.log(f6(true, 0)); // reutrn 100
//　在一个嵌套作用域里引入一个新名字的行为称为屏蔽。它是一把双刃剑，它可
// 能会不小心引入新问题，同时也可能会解决一些错误。
function sumMatrix(matrix) {
    var sum = 0;
    for (var i = 0; i < matrix.length; i++) {
        var currentRow = matrix[i];
        for (var i_1 = 0; i_1 < currentRow.length; i_1++) {
            sum += currentRow[i_1];
        }
    }
    return sum;
} // 通常来讲应该避免使用屏蔽，因为我们需要写出清晰的代码。
/* 块级作用域变量的获取 */
function theCityThatAlwaysSleeps() {
    var getCity;
    if (true) {
        var city_1 = 'Seattle';
        getCity = function () {
            return city_1;
        };
    }
    return getCity();
}
var _loop_1 = function (i) {
    setTimeout(function () { console.log(i); }, 100 * i);
};
for (var i = 0; i < 10; i++) {
    _loop_1(i);
}
/* const 声明 */
// const 声明是声明变量的另一种方式。
var numLivesForCat = 9;
// 它与let声明相似，但是就像它的名字表达的，它被赋值后不能再
// 改变。
var numLivesForCat2 = 9;
var kitty = {
    name: 'Aurora',
    numLives: numLivesForCat2
};
// err
// kitty = {
//     name: 'Danielle', 
//     numLives: numLivesForCat2
// };
kitty.name = 'Rory';
kitty.name = 'Kitty';
kitty.name = 'Cat';
kitty.numLives--;
/* 解构 */
// 解构数组
var input = [1, 2];
var first = input[0], second = input[1];
console.log(first);
console.log(second);
first = input[0];
second = input[1];
_a = [second, first], first = _a[0], second = _a[1];
// 作用于函数参数：
function f7(_a) {
    var first = _a[0], second = _a[1];
    console.log('data:function f7([first, second]:array[number, number]) : frist: ' + first);
    console.log('data: function f7([first, second]: array[number, number]): second:' + second);
}
f7(input);
var _b = [1, 2, 3, 4, 5], first1 = _b[0], rest1 = _b.slice(1);
console.log(first1);
console.log(rest1);
var first2 = [1, 2, 3, 4, 5][0];
console.log('data:[1, 2, 3, 4, 5]->array:first2:' + first2);
var _c = [1, 2, 3, 4, 5], second2 = _c[1], fourth = _c[2];
console.log('data:[1, 2, 3, 4, 5]->array:second2: ' + second2, 'array:fourth: ' + fourth);
/**
解构对象
*/
var o = {
    a: 'foo',
    b: 12,
    c: 'bar'
};
// let {a, b} = o;
// console.log('data: o={a:"foo",b:12,c:"bar"}->object:a ',  'object:b: '+b);
// // 可以用没有声明的赋值（需要用括号括起来，因为javascript通常会将{起始的语句解析为一个块）
// ({a, b} = {a: 'baz', b: 101});
// console.log(`赋值： a: ${ a } b: ${ b }`);
var a = o.a, passthrough = __rest(o, ["a"]);
var total = passthrough.b + passthrough.c.length;
console.log("total " + total + " ");
// 属性重命名
var newName1 = o.a, newName2 = o.b;
// ==
// let newName1 = o.a;
// let newName2 = o.b;
// let {a, b}: {a: string, b: number} = o;
// 默认值
function keepWholeObject(wholeObject) {
    var a = wholeObject.a, _a = wholeObject.b, b = _a === void 0 ? 10001 : _a;
}
function f8(_a) {
    var a = _a.a, b = _a.b;
}
// 指定默认值
function f9(_a) {
    var _b = _a === void 0 ? { a: '', b: 0 } : _a, a = _b.a, b = _b.b;
}
f9();
function f10(_a) {
    var _b = _a === void 0 ? { a: '' } : _a, a = _b.a, _c = _b.b, b = _c === void 0 ? 0 : _c;
}
f10({ a: 'yes' });
f10();
// f10({}); // err
// 展开
// 展开操作符正与解构相反。它允许你将一个数组展开为另一个数组，或将一个对象展开为
// 另一个对象。
var first3 = [1, 2];
var second3 = [3, 4];
var bothPlus = [0].concat(first3, second3, [5]);
console.log("\u5C55\u5F00 bothPlus " + bothPlus);
// 展开对象
var defaults = { food: 'spicy', price: '$$', ambiance: 'noisy' };
var search = __assign({}, defaults, { food: 'rich' });
console.log('defaults: ' + defaults);
console.log('search={...defaults, food: "rich"}: ' + search);
// 对象的展开比数组的展开要复杂的多。像数组展开一样，它是从左至右进行处理，但
// 结果仍为对象。这就意味着出现在展开对象后面的属性会覆盖前面的属性。
var defaults2 = { food: 'spicy', price: '$$', ambiance: 'noisy' };
var search2 = __assign({ food: 'rich' }, defaults2);
var C1 = /** @class */ (function () {
    function C1() {
        this.p = 12;
    }
    C1.prototype.m = function () {
    };
    return C1;
}());
var c = new C1();
var clone = __assign({}, c);
console.log('C1 class->p param:' + clone.p);
var _a;
//clone.m(); // err
