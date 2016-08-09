var connect = require('connect');

/**
 * 创建服务器
 */

var server = connect.createServer();

/**
 * 中间件
 */
server.use(function (req, res, next) {
	//记录日志
	console.error(' %s %s', req.method, req.url);
	next();
});

server.use(function (req, res, next) {
	if ('GET' == req.method && '/images' == req.url.substr(0, 7)){
		//托管图片
	} else {
		// 交给其他的中间件去处理
		next();
	}
});

server.use(function (req, res, next) {
	if ('GET' == req.method && '/' == req.url){
		// 响应index文件
	} else {
		// 交给其他中间件去处理
		next();
	}
});

server.use(function (req, res, next) {
	//最后一个中间件，如果到了这里， 就意味着无能为力，只能返回404了
	res.writeHead(404);
	res.end('Not found');
})

/**
 * 处理静态文件
 */

server.use(connect.static(__dirname + './'));

/**
 * 监听
 */
server.listen(3000);