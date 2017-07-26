console.log('当前执行文件所在的目录是：', process.cwd());

var dir = process.argv;
console.log('Node的安装路径是：', dir[0]);
console.log('当前文件的完整路径是：', dir[1]);
console.log('当前执行文件所在的目录是：', __dirname);

//process.chdir方法，允许灵活地更改工作目录
process.chdir('/');
console.log('更改后的工作目录是:', process.cwd());

//process.env变量访问shell环境下的变
console.log(process.env);

//获取控制程序是运行在开发模式下还是产品模式下
console.log(process.env.NODE_ENV);
console.log(process.env.SHELL);

//退出  
	//要让一个应用退出， 可以调用process.exit并提出一个退出代码。
	//退出后，退出程序以后的代码不会执行、

console.error('An error occurred');
//process.exit(1);
console.log('dffdfdf');


//Node程序是通过process 对象上以事件分发的形式来发送信号的
//process.on('SIGKILL', function () {
	//信号已收到
//});


// ANSI 转义码
	//要在文本终端下控制格式、颜色以及其他输出选项，可以使用ANSI转义码
	//在文本周围添加的明显不用于输出的字符，称为非打印字符
	//console.log('\033[90m' + data.replace(/(.*)/g, '    $1') + '\033[39m');

	// \033表示转义序列的开始
	// [表示开始颜色设置
	// 90表示前景色为亮灰色
	// m表示颜色设置结束

	// 最后用的是 \033[39m 是将颜色再设置回去
