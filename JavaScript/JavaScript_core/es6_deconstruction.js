/*
	解析结构是一种从对象或者数组中取得值扔一种全新的写法 
 */

var tom = {
	name: 'TOM',
	age : 20, 
	gender: 1,
	job: 'studend'
}

var name = tom.name;
var age = tom.age;
var pender = tom.gender;
var job = tom.job;


// 解析结构
const { name, age, gender, job } = tom;
// 还可以这样写
// const {name: anme, age: age, gender: gender, job: job}


const peoples = {
	consts: 100,
	detail : {
		tom: {
			name: 'Tom',
			age : 20,
			gender: 1,
			job: 'student'
		}
	}
}



const { detail: { tom } } = peoples;

const { detail: {tom: {age, gender}}} = peoples;
