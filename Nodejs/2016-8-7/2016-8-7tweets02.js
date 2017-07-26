var qs = require('querystring')
	, http = require('http');

var search = process.argv.slice(2).join(' ').trim();

if (!search.length) {
	retrun console.log('\n Usage: node tweets <search term>\n');
}

http.get({
	host: 'search.twitter.com'
	, path: '/search.json?' + qs.stringify({ q: 'search'})
}, function (res) {
	var body = '';
	res.setEncoding('utf8');
	res.on('data', function (chunk) {
		body += chunk;
	});
	res.on('end', function () {
		var obj = JSON.parse(body);
		obj.results.forEach(function (tweet) {
			console.log('	\033[90m' + tweet.text + '\033[39m\n');
			console.log('	\033[94m' + tweet.from_user + '\033[39m\n');
			console.log('--');
		});
	});
})