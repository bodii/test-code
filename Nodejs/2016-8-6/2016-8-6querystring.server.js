var qs = require('querystring');
require('http').createServer(function (req, res) {
	if ('/' == req.url) {
		res.writeHead(200, {'Content-Type' : 'text/html'});
		res.end([
			'<form method="post" action="/url">'
			,	'<h1>My form</h1>'
			,	'<fieldset>'
			,	'<lable>Personal informaction</lable>'
			,	'<p>what is your name?</p>'
			,	'<input type="text" name="name">'
			,	'<p><button>Submit</button></p>'
			,'</form>'
			].join(''));
	} else if ('/url' == req.url && 'POST' == req.method) {
		var body = '';
		req.on('data', function (chunk) {
			body += chunk;
		});
		req.on('end', function () {
			res.writeHead(200, {'Content-Type' : 'text/html'});
			res.end([
				'<p>Your name is <b>' + qs.parse(body).name + '</b></p>'
				].join(''));
		})
	}
}).listen(3000);