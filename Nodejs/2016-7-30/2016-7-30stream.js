// example -1
console.log('Hello world');
console.log('this is test line');

// example -2 process 
process.stdout.write('Hello world');
console.log('this is test line');

/*
	process 全局对象中包含了三个流对象， 分别对应三个unix 标准流;
	stdin : 标准输入
	stdout : 标准输出
	stderr : 标准错误
*/
