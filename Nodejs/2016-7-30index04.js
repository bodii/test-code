//# 2016-7-30 index04.js
/**
* Modeule dependencies.
*/

var fs = require('fs');

// 获取当前目录的文件列表 （两种方法） 同步
console.log(fs.readdirSync(__dirname));
console.log(fs.readdirSync('.'));

// 异步的版本
function async (err, files) {
	console.log(files);
};
fs.readdir('.', async);