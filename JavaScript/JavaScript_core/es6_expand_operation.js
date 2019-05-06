/*
 展开运算

 在ES6中，使用...来表示展开运算符，它可以展开数组/对象
 */

// 声明一个数组
const arr1 = [1, 2, 3];

// 其次声明另一个数组，新数组中包含数组arr1的所有元素，
// 因此可以利用展开运算符
const arr2 = [...arr1, 4, 5, 6];
console.log(arr2);


// 展开对象
const object1 = {
	a: 1,
	b: 2,
	c: 3
}

const object2 = {
	...object1,
	d: 4,
	e: 5,
	f: 6
}

console.log(object2);


// 在解析结构中，也常常使用展开运算符
const tom = {
	name: 'TOM',
	age: 20,
	gender: 1,
	job: 'student'
}

const { name, ...others } = tom;
console.log(others);


const add = (a, b, ...more) => {
	return more.reduce((m, n) => m + n) + a + b;
}

console.log(add(1, 23, 1, 2, 3, 4, 5));
