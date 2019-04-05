// 除了__proto__对象本身，所有对都继承自__proto__
let person = { name: 'Mosh' };

// Object.getPrototypeOf 获取这个对象出自哪里
// let objectBase = Object.getPrototypeOf(person);

// Ojbect.getOwnPropertyDescriptor 获取属性描述器
// let descriptor = Object.getOwnPropertyDescriptor(objectBase, 'toString');
// console.log(descriptor);

// Object.defineProperty 在一个对象上定义一个属性，或修改现有的属性，并返回这个对象
Object.defineProperty(person, 'name', {
    writable: false,
    enumerable: true, // 所有属性可写，可枚举，可配置
    configurable: false, // 如果设置成false, 这个属性不能被删除
});

person.name = 'John'; // 如果writable  为false 这个设置将失效
// delete person.name; //如果configurable 为false 这个将失效
console.log(person);