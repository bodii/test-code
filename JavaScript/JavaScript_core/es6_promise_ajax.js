// 封装一个ajax的get请求方法
function getJSON(url) {
	return new Promise(function(resolve, reject) {
		// 利用Ajax发送一个请求
		var XHR = new XMLHttpRequest();
		XHR.open('GET', url, true);
		XHR.send();

		// 等待结果
		XHR.onreadystatechange = function() {
			if (XHR.readyState == 4) {
				if (XHR.status == 200) {
					try {
						var response = JSON.parse(XHR.responseText);
						// 得到正确的结果修改状态并将数据传递出去
						resolve(response);
					} catch (e) {
						reject(e);
					}
				} else {
					reject(new Error(XHR.statusText));
				}
			}
		}
	});
}


// 封装好之后，使用就很简单了
getJSON(url).then(function(resp) {
	console.log(resp);
});

// 现在几乎所有的库都利用Promise对Ajax请求进行了封装，当然也包括常用的jQuery,
// 因此在使用jQuery等库中的Ajax请求时，可以利用Promise来让代码更加的简洁和优雅
/*
 $.get(url).then(function(resp) {
	// 处理success的结果
 })
 .catch(function(err) {
 });
 */
