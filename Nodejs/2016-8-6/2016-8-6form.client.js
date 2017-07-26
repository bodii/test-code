var http = require('http')
	, qs = require('querystring');

function send (theName) {
	http.request({
		host: '127.0.0.1'
		, port: 3000
		, url: '/'
		, method: 'POST'
	}, function (res) {
		res.setEncoding('utf8');
		res.on('data', function () {
			console.log('\n   \033[90m  request complete!\033[39m');
			process.stdout.write('\n  your name: ');
		});
	}).end(qs.stringify({name: theName}));
	// querystring模块的stringify方法，可以将一个对象
	// 转化为url编码过的数据
}

process.stdout.write('\n  your name: ');
process.stdin.resume();
process.stdin.setEncoding('utf-8');
process.stdin.on('data', function (name) {
	send(name.replace('\n', ''));
});