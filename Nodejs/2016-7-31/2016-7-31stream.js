// fs.createReadStream 方法允许为一个文件创建一个可读的stream对象
var fs = require('fs');
// fs.readFile('somefile.txt', function (err, contents) {
// 	// 对文件进行处理
// });
// 回调函数必须要等到整个文件读取完毕、载入到RAM、可用的情况下才会触发。

// var stream = fs.createReadStream('somefile.txt');
// stream.on('data', function (chunk) {
// 	// 处理文件部分内容
// });
// stream.on('end', function (chunk) {
// 	//文件读取完毕
// });


// 监视 
	// Node允许监视文件或目录是否发生变化。监视意味着当文件系统中的
	// 文件(或者目录)发生变化时，会分发一个事件，然后触发指定的回调函数


	// example
	//查找工作目录下所有的CSS文件,然后监视其是否发生改变。一旦文件发生更改，
	//就将该文件名输出到控制台

	//var stream = fs.createReadStream('somefile.txt');
	//获取工作目录下所有的文件
	var files = fs.readdirSync(process.cwd());
	files.forEach(function (file) {
		//监听".css"后缀的文件
		if (/\.css/.test(file)) {
			fs.watchFile(process.cwd() + '/' + file, function () {
				console.log(' - ' + file + ' changed!');
			});
		}
	});


// fs.watchFile 监视某个文件， fs.watch监视整个目录