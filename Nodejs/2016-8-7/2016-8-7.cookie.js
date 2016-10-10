var connect = require('connect');

var server = connect.createServer();
//使用cookieParser中间件
server.use(connect.cookieParser());
//然后就可以通过req.cookies对象访问cookie数据了
server.use(function (req, res, next) {
	req.cookies.secret1 = "value1";
	req.cookies.secret2 = "value2";
})

server.listen(3000);