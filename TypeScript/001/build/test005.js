/**
 * 函数
 *
 * 函数是JavaScript应用程序的基础。它帮助你实现抽象层，模拟类，信息隐藏和模块。
 * 在TypeScript里，虽然已经支持类，命名空间和模块，但函数仍然是主要定义行为的地方。
 * TypeScript为JavaScript函数添加了额外的功能。
 * TypeScript函数可以创建有名字的函数和匿名函数。
 */
function add(x, y) {
    return x + y;
}
var myAdd = function (x, y) { return x + y; };
var z = 100;
function addToZ(x, y) {
    return x + y + z;
}
// 函数类型
function add2(x, y) {
    return x + y;
}
var myAdd2 = function (x, y) { return x + y; };
// 书写完整函数类型
var myAdd3 = function (x, y) { return x + y; };
/*
    函数类型包含两部分： 参数类型和返回值类型。当写出完整函数类型的时候，这两部分都是需要的。
*/
//　也可以这么写
var myAdd4 = function (x, y) { return x + y; };
// 只要参数类型是匹配的，那么就认定是有效的函数类型，而不在乎参数名是否正确
/*
    推断类型

如果你在赋值语句的一边指定 了类型但是另一边没有类型的话，TypeScript编译器会自动识别出类型:
*/
var myAdd5 = function (x, y) { return x + y; };
var myAdd6 = function (x, y) { return x + y; };
// 这叫做“按上下文归类”， 是类型推论的一种。它帮助我们更好地为程序指定类型。
/*
    可选参数和默认参数

TypeScript里的每个函数参数都是必须的。这不是指不能传递null或undefined作为参数，而是说编译器
检查用户是否为每个参数都传入了值。编译器还会假设只有这些参数会被传递进函数。简短地说，传递一个
函数的参数个数必须与函数期望的参数个数一致。
*/
function buildName(fristName, lastName) {
    return fristName + ' ' + lastName;
}
// let result1 = buildName('Bob'); //err
// let result1 = buildName('Bob', 'Adams', 'Sr.'); // err
var result1 = buildName('Bob', 'Adams');
// JavaScript里，每个参数都是可选的， 可传可不传。没传参数的时候，它的值是undefined。在
// TypeScript里可以在参数名旁使用？实现可选参数的功能。
function buildName2(firstName, lastName) {
    if (lastName)
        return firstName + ' ' + lastName;
    else
        return firstName;
}
var result2 = buildName2('Bob');
// let result3 = buildName2('Bob', 'Adams', 'Sr.'); // err
var result3 = buildName2('Bob', 'Adams');
/*
    !!!
可选参数必须跟在必须参数后面。如果上面firstName是可选的，那么就必须调整它们的位置，
把firstName放在后面。

在TypeScript里，可以为参数提供一个默认值，当用户没有传递这个参数或传递的值是undefined时。
它们叫做有默认初始化值的参数。
*/
function buildName3(firstName, lastName) {
    if (lastName === void 0) { lastName = 'Smith'; }
    return firstName + ' ' + lastName;
}
var result4 = buildName3('Bob');
console.log(result4);
var result5 = buildName3('Bob', undefined);
console.log(result5);
// let result6 = buildName3('Bob', 'Adams', 'Sr.'); // err
var result6 = buildName3('Bob', 'Adams');
console.log(result6);
/*
在所有必须参数后面的带默认初始化的参数都是可选的。
与可选参数一样，在调用函数的时候可以省略。也就是说可选参数与末尾的默认参数共享参数类型

function buildName(firstName: string, lastName?: string) { }
和
function buildName(fristName: string, lastName = 'Smith') { }

共享同样的类型(firstName: string, lastName?: string) => string。 默认参数的默认值
消失了，只保留了它是一个可选参数的信息。

与普通可选参数不同的是，带默认值的参数不需要放在必须参数的后面。
如果带默认值的参数出现在必须参数的前面，用户必须明确的传入undefined值来获得默认值。

*/
function buildName4(firstName, lastName) {
    if (firstName === void 0) { firstName = 'Will'; }
    return firstName + ' ' + lastName;
}
// let result7 = buildName4('Bob'); // err 
// let result7 = buildName('Bob', 'Adams', 'Sr.'); // err
var result7 = buildName('Bob', 'Adams');
var result8 = buildName(undefined, 'Adams');
// let result9 = buildName(1, 'Bob'); // err : 默认值的类型 默认成为参数的类型
// 剩余参数
/*
    必要参数，默认参数和可选参数有个共同点： 它们表示某一个参数。有时，你想同时操作多个参数，
    或者你并不知道会有多少参数传递进来。在JavaScript里，你可以使用arguments来访问所有传入
    的参数。
 */
function buildName5(firstName) {
    var restOfName = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        restOfName[_i - 1] = arguments[_i];
    }
    return firstName + ' ' + restOfName.join(' ');
}
var employeeName = buildName5('Joseph', 'Samuel', 'Lucas', 'Mackinzie');
// 剩余参数会被当做个数不限的可选参数。可以不一个都没有，同样也可以有任意个
function buildName6(firstName) {
    var restOfName = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        restOfName[_i - 1] = arguments[_i];
    }
    return firstName + ' ' + restOfName.join(' ');
}
var buildNameFun = buildName6;
/**
 * this 和 箭头函数
 *
 * JavaScript里，this的值在函数被调用的时候才会指定。这是个既强大又灵活的特点，但是你需要
 * 花点时间弄清楚函数调用的上下文是什么
 */
// emapmle
var deck = {
    suits: ['hearts', 'spades', 'clubs', 'diamonds'],
    cards: Array(52),
    createCardPicker: function () {
        return function () {
            var pickedCard = Math.floor(Math.random() * 52);
            var pickedSuit = Math.floor(pickedCard / 13);
            return { suit: this.suit[pickedSuit], card: pickedCard % 13 };
        };
    }
};
var cardPicker = deck.createCardPicker();
var pickedCard = cardPicker();
console.log('card: ' + pickedCard.card + ' of ' + pickedCard.suit);
// createCardPicker返回的函数里的this被设置成了window而不是deck对象。
// 因为只是独立使用了cardPicker().顶级的非方法式调用会将this视为window。（
// 注意：在严格模式下，this为undefined而不是window）。
// 可以在函数返回时就绑定好正确的this。需要改变函数表达式来使用ECMAScript6箭头语法。
// 箭头函数能保存函数创建时的this值，而不是调用时
var deck2 = {
    suits: ['hearts', 'spades', 'clubs', 'diamonds'],
    cards: Array(52),
    createCardPicker: function () {
        var _this = this;
        return function () {
            var pickedCard = Math.floor(Math.random() * 52);
            var pickedSuit = Math.floor(pickedCard / 13);
            return { suit: _this.suits[pickedSuit], card: pickedCard % 13 };
        };
    }
};
var cardPicker2 = deck2.createCardPicker();
var pickedCard2 = cardPicker2();
console.log('card: ' + pickedCard2.card + ' of ' + pickedCard2.suit);
// this参数
// this.suits[pickedSuit]类类型为any。 这是因为this来自对象字面量的函数表达式
// 修改的方法是， 提供一个显式的this参数。this参数是个假的参数，它出现在参数列表的
// 最前面。
function f() {
    // make sure 'this' is unusable in this standalone function
}
var deck3 = {
    suits: ['hearts', 'spades', 'clubs', 'diamonds'],
    cards: Array(52),
    createCardPicker: function () {
        var _this = this;
        return function () {
            var pickedCard = Math.floor(Math.random() * 52);
            var pickedSuit = Math.floor(pickedCard / 13);
            return { suit: _this.suits[pickedSuit], card: pickedCard % 3 };
        };
    }
};
var carPicker = deck3.createCardPicker();
var pickedCard3 = cardPicker();
console.log('card: ' + pickedCard.card + ' of ' + pickedCard.suit);
var Handler = /** @class */ (function () {
    function Handler() {
    }
    Handler.prototype.onClickBad = function (e) {
        // this.info = e.message; // err
    };
    return Handler;
}());
var h = new Handler();
// uiElement.addClickListener(h.onClickBad); // err
var Handler2 = /** @class */ (function () {
    function Handler2() {
    }
    Handler2.prototype.onClickGood = function (e) {
        console.log('clicked!');
    };
    return Handler2;
}());
var h2 = new Handler2();
// UIElement.addClickListener(h2.onClickGood);
// 重载
// JavaScript本身是个动态语言。JavaScript里函数根据传入不同的参数而返回不同类型的数据
var suits = ['hearts', 'spades', 'clubs', 'diamonds'];
function pickCard(x) {
    if (typeof x == 'object') {
        var pickedCard_1 = Math.floor(Math.random() * x.length);
        return pickedCard_1;
    }
    else if (typeof x == 'number') {
        var pickedSuit = Math.floor(x / 13);
        return { suit: suits[pickedSuit], card: x % 13 };
    }
}
var myDeck = [
    { suit: 'diamonds', card: 2 },
    { suit: 'spades', card: 10 },
    { suit: 'hearts', card: 4 },
];
var pickedCard4 = myDeck[pickCard(myDeck)];
console.log('card: ' + pickedCard4.card + ' of ' + pickedCard4.suit);
var pickedCard5 = pickCard(15);
console.log('card: ' + pickedCard5.card + ' of ' + pickedCard5.suit);
function pickCard3(x) {
    if (typeof x == 'object') {
        var pickedCard_2 = Math.floor(Math.random() * x.length);
        return pickedCard_2;
    }
    else if (typeof x == 'number') {
        var pickedSuit = Math.floor(x / 13);
        return { suit: suits[pickedSuit], card: x % 13 };
    }
}
var pickedCard6 = myDeck[pickCard3(myDeck)];
console.log('card: ' + pickedCard6.card + ' of ' + pickedCard6.suit);
var pickedCard7 = pickCard3(15);
console.log('card: ' + pickedCard2.card + ' of ' + pickedCard2.suit);
function vv(x) { }
;
var v = function (x) { };
var v1;
v1(4);
