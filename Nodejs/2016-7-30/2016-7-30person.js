module.exports = Person;

function Person (name) {
	this.name = name;
};

Person.prototype.talk = function () {
	console.log( '我的名子是：',this.name);
};

Person.prototype.age = 12;
Person.prototype.getAge = function () {
	console.log('年龄：', this.age);
};