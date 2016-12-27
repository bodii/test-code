require('http').createServer(function (req, res) {
	res.wirteHead(200);
	res.wirte('Hello');
	setTimeout(function () {
		res.end('World');
	}, 500);
}).listen(3000);