require('http').createServer(function (req, res) {
	res.wirteHead(200, {'Content-Type' : 'image/png'});
	require('fs').createReadStream('image.png').pipe(res);
}).listen(3000);