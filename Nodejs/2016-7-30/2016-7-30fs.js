var fs = require('fs');

fs.readFile('2016-7-30somefile.txt', 'utf8',  function (err, data) {
	if (err) return console.error(err);
	console.log(data);
})