require('http').createServer(function (req, res) {
	if ('/' == req.url) {
		res.writeHead(200, {'Content-Type' : 'text/html'});
		res.end([
				'<form method="post" action="/url">'
				,	'<h1>My form</h1>'
				,	'<fieldset>'
				,	'<label>Personal informaction</label>'
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
			res.end(['<p>Content-Type : ' + req.headers['content-type'] + '</p>'
				,'<p>Data:</p>'
				,'<pre>' + body + '</pre>'
				].join(''));
		})
	}
}).listen(3000);