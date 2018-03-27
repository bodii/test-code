/**
 * 变量的声明
 */


/* let 声明 */

let hello = 'Hello!';

/* 块作用域 */
function f(input: boolean) {
    let a = 100;

    if(input) {
        let b = a + 1;
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



try{
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


let x = 10;
// let x = 20; // err, 不能在1个作用域里多次声明1个变量。

// 并不是要求两个均是块级作用域的声明TypeScript才会给出一个错误的警告
// function f5(x) {
//     let x = 100; // err
// }

function g() {
    let x = 100;
    // var x = 100; // err
}


// 并不是说块级作用域变量不能用函数作用域变量来声明。而是块级作用域变量
// 域变量需要在明显不同的块里声明
function f6(condition: boolean, x: number) {
    if (condition) {
        let x = 100;
        return x;
    }

    return x;
}

console.log(f6(false, 0)); // return 0
console.log(f6(true, 0)); // reutrn 100



//　在一个嵌套作用域里引入一个新名字的行为称为屏蔽。它是一把双刃剑，它可
// 能会不小心引入新问题，同时也可能会解决一些错误。
function sumMatrix(matrix: number[][]) {
    let sum = 0;
    for (let i = 0; i < matrix.length; i++) {
        var currentRow = matrix[i];
        for (let i = 0; i < currentRow.length; i++) {
            sum += currentRow[i];
        }
    }

    return sum;
} // 通常来讲应该避免使用屏蔽，因为我们需要写出清晰的代码。


/* 块级作用域变量的获取 */
function theCityThatAlwaysSleeps() {
    let getCity;

    if (true) {
        let city = 'Seattle';
        getCity = function() {
            return city;
        }
    }

    return getCity();
}


for (let i = 0; i < 10; i ++) {
    setTimeout(function() {console.log(i);}, 100 * i);
}




/* const 声明 */

// const 声明是声明变量的另一种方式。
const numLivesForCat = 9;

// 它与let声明相似，但是就像它的名字表达的，它被赋值后不能再
// 改变。

const numLivesForCat2 = 9;
const kitty = {
    name: 'Aurora',
    numLives: numLivesForCat2,
};

// err
// kitty = {
//     name: 'Danielle', 
//     numLives: numLivesForCat2
// };

kitty.name = 'Rory';
kitty.name = 'Kitty';
kitty.name = 'Cat';
kitty.numLives --;



/* 解构 */

// 解构数组
let input:[number, number] = [1, 2];
let [first, second] = input;
console.log(first);
console.log(second);

first = input[0];
second = input[1];

[first, second] = [second, first];


// 作用于函数参数：
function f7([first, second]: [number, number]) {
    console.log('data:function f7([first, second]:array[number, number]) : frist: ' +first);
    console.log('data: function f7([first, second]: array[number, number]): second:' +second);
}

f7(input);


let [first1, ...rest1] = [1, 2, 3, 4, 5];
console.log(first1);
console.log(rest1);

let [first2] = [1, 2, 3, 4, 5]
console.log('data:[1, 2, 3, 4, 5]->array:first2:'+first2)

let [, second2, fourth] = [1, 2, 3, 4, 5]
console.log('data:[1, 2, 3, 4, 5]->array:second2: '+second2, 'array:fourth: '+fourth)



/** 
解构对象 
*/

let o = {
    a: 'foo',
    b: 12, 
    c: 'bar'
};
// let {a, b} = o;
// console.log('data: o={a:"foo",b:12,c:"bar"}->object:a ',  'object:b: '+b);

// // 可以用没有声明的赋值（需要用括号括起来，因为javascript通常会将{起始的语句解析为一个块）
// ({a, b} = {a: 'baz', b: 101});
// console.log(`赋值： a: ${ a } b: ${ b }`);

let { a, ...passthrough } = o;
let total = passthrough.b + passthrough.c.length;
console.log(`total ${ total } `);



// 属性重命名
let { a: newName1, b: newName2 }  = o;
// ==
// let newName1 = o.a;
// let newName2 = o.b;

// let {a, b}: {a: string, b: number} = o;



// 默认值
function keepWholeObject(wholeObject: { a: string, b?: number }) {
    let { a, b = 10001 } = wholeObject;
}
// 现在，即使b为undefined, keepWholeObject函数的变量wholeOjbect的属性a 和b都会有值。


// 函数声明
type C = { a: string, b?: number};
function f8({ a, b }: C): void {

}

// 指定默认值
function f9({ a, b } = { a: '', b: 0 }): void {

} 
f9();


function f10({ a, b = 0 } = { a: '' }): void {

}
f10({ a: 'yes' });
f10();
// f10({}); // err


// 展开
// 展开操作符正与解构相反。它允许你将一个数组展开为另一个数组，或将一个对象展开为
// 另一个对象。
let first3 = [1, 2];
let second3 = [3, 4];
let bothPlus = [0, ...first3, ...second3, 5];
console.log(`展开 bothPlus ${ bothPlus }`);

// 展开对象
let defaults = { food: 'spicy', price: '$$', ambiance: 'noisy' };
let search = { ...defaults, food: 'rich' };
console.log('defaults: ' + defaults);
console.log('search={...defaults, food: "rich"}: ' + search);
// 对象的展开比数组的展开要复杂的多。像数组展开一样，它是从左至右进行处理，但
// 结果仍为对象。这就意味着出现在展开对象后面的属性会覆盖前面的属性。
let defaults2 = { food: 'spicy', price: '$$', ambiance: 'noisy' };
let search2 = { food: 'rich', ...defaults2 };


class C1 {
    p = 12;
    m() {

    }
}
let c = new C1();
let clone = { ...c };
console.log('C1 class->p param:' + clone.p);
//clone.m(); // err


// 其次， TypeScript编译器不允许展开泛型函数上的类型参数。（在未来的版本中可能添加）