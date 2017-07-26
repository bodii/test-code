/**
 * 模块依赖
 */
var http = require('http')
	, fs = require('fs');

/**
 * 创建服务器
 */
var server = http.createServer(function (req, res) {
	if ('GET' == req.method && '/2016-8-7images' == req.url.substr(0, 7) && '.jpg' == req.url.substr(-4)) {
		fs.stat(__dirname + req.url, function (err, stat) {
			if (err || !stat.isFile()) {
				res.writeHead(404);
				res.end('Not Found');
				return;
			}
			serve(__dirname + req.url, 'application/jpg');
		});
	} else if ('GET' == req.method && '/' == req.url) {
		serve(__dirname + '/2016-8-7index.html', 'text/html');
	} else {
		res.writeHead(404);
		res.end('Not found');
	}
	
	function serve (path, type) {
		res.writeHead(200, {'Content-type' : type});
		fs.createReadStream(path).pipe(res); // pipe 是将文件系统流接到HTTP响应流中
	}
});

/**
 * 监听
 */
server.listen(3000);
