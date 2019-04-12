// 高阶函数withLogin,用来判断当前用户状态

(function() {
	// 用随机数的方式来模拟一个获取用户信息的方法
	var getLogin = function() {
		var a = parseInt(Math.random() * 10).toFixed(0));
		if (a % 2 == 0) {
			return {
				login: false
			};
		}

		return {
			login: true,
			userInfo: {
				nickname: 'jack',
				vip: 11,
				userid: '666666',
			}
		}
	}

	var withLogin = function(basicFn) {
		var loginInfo = getLogin();

		return basicFn.bind(null, loginInfo);
	}

	window.withLogin = withLogin;
})();


// 通过renderIndex的方法来渲染
// 首页
(function() {
	var withLogin = window.withLogin;

	var renderIndex = function(loginInfo) {
		if (loginInfo.login) {
			// 处理已经登录之后的逻辑
		} else {
			// 处理未登录的逻辑
		}
	}

	window.renderIndex = withLogin(renderIndex);
})();


// 个人中心页
(function() {
	var withLogin = window.withLogin;

	var renderPersonal = function(loginInfo) {
		if (loginInfo.login) {
			// do something
		} else {
			// do other something
		}
	}

	window.renderPersonal = withLogin(renderPersonal);
})();


// 渲染
(function() {
	window.renderIndex();
})();
