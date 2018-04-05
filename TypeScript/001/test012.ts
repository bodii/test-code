/**
 * 迭代器和生成器
 */

/* 可迭代性
    当一个对象实现了Symbol.iterator属性时，我闪认为它是可迭代的。一些内置的类型如
    Arry, Map, Set, String, Int32Array, Uint32Array等都已经实现了各自的Symbol.iterator。
    对象Symbol.iterator函数负责返回供迭代的值。
*/


/**
 * for..of语句
 * for..of会遍历可迭代的对象，调用对象上的Symbol.iterator方法。
 */
let someArray = [1, 'string', false];
for(let entry of someArray) {
    console.log(entry); // 1, 'string', false
}

/* for..of vs for..in语句


1.for..of 和for..in均可迭代一个列表；但是用于迭代的值却不同，for..in迭代的是对象的键的列表，for..
of迭代的对象的键对应的值。

2. for..in可以操作任何对象;它提供了查看对象属性的一种方法。但是for..of关注于迭代对象的值。内置对象
Map和Set已经实现了Symbol.iterator方法， 让我们可以访问它们保存的值。
*/

let list = [4, 5, 6];
for (let i in list) {
    console.log(i); // 0, 1, 2 
}

for(let i in list) {
    console.log(i); // 4, 5, 6
}

let pets = new Set(['Cat', 'Dog', 'Hamster']);
pets['species'] = 'mammals';

for( let pet in pets ) {
    console.log(pet); // 'species'
}

for(let pet of pets) {
    console.log(pet); // 'Cat', 'Dog', 'Hamster'
}

/*
当生成目标为ES5或ES3，迭代器只允许在Array类型上使用。在非数组值上使用for..of语句会得到一个错误，
就算这些非数组值已经实现了Symbol.iterator属性。
*/

let numbers = [1, 2, 3];
for (let num of numbers) {
    console.log(num);
}

/*
编译后生成的JavaScript代码
var numbers = [1, 2, 3]
for (var _i = 0; _i < numbers.length; _i++) {
    var num = numbers[_i];
    console.log(num);
}

*/