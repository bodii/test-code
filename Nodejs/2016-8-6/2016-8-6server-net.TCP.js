
/**
 * 追踪连接数
 */

var count = 0
	, users = {}
	, net = require('net');

var server = net.createServer(function (conn) {
	//设置输出编辑码
	conn.setEncoding('utf8');
	conn.write(
		'\n > welcome to \033[92mnode-chat\033[39m!'
		+ '\n > ' + count + ' other people are connected at this time.'
		+ '\n > please write your name and press enter:'
		);

	count ++;

	// 当前连接的昵称
	var nickname;
	conn.on('data', function (data) {
		//删除回车符
		data = data.replace('\r\n', '');
		if (!nickname) {
			if (users[data]) {
				conn.write('\033[93m> nickname already in use. try agein:\033[39m  ');
				return;
			} else {
				nickname = data;
				users[nickname] = conn;
				for (var i in users) {
					if (i != nickname) {
						users[i].write('\033[90m > ' + nickname + ':\033[39m ' + data +'\n');
						broadcast('\033[90m > ' + nickname + ':\033[39m ' + data +'\n');
					} else {
						users[i].write('\033[90m > ' + nickname + ' joined the room\033[39m\n');
						broadcast('\033[90m > ' + nickname + ' joined the room\033[39m\n');
					}
				}
			}
		}
	});

	//当客户端请求关闭时，计数器递减
	conn.on('close', function () {
		count --;
		delete users[nickname];
		broadcast('\033[90m > ' + nickname + ' left the room\033[39m\n');
	});
});

server.listen(3000, function () {
	console.log('\033[96m   server listening on *:3000\033[39m');
});


// 用户断开时通知其他用户
function broadcast (msg, exceptMyself) {
	for (var i in users) {
		if (!exceptMyself || i != nickname) {
			users[i].write(msg);
		}
	}
}
