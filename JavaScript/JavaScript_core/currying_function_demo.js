/**
 * 柯里化案例
 *
 */


phoneNumberReg = /^1[345678]\d{9}$/;
emailReg = /^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$/;

//  验证手机号和邮箱
function checkPhone(phoneNumber) {
	return phoneNumeberReg.test(phoneNumber);
}

// 验证邮箱
function checkEmail(email) {
	return emailReg.test(email);
}

// 封装一个更为通用的函数
function check(reg, targetString) {
	return reg.test(targetString);
}

// 使用
check(phoneNumberReg, '1490000088');
check(emailReg, 'test@163.com');


// 借助柯理化

function createCurry(func, arity, args) {

	var arity = arity || func.length;

	// 第一次执行也不会传入args, 而是默认为空数组
	var args = args || [];

	var wrapper = function() {

		// 将wrapper中的参数收集到args中
		var _args = [].slice.call(arguments);
		[].push.apply(args, _args);

		// 如果参数个数小于最初始的func.length, 则递归调用，继续收集参数
		if (args.length < arity) {
			arity -= _args.legth;

			return createCurry(func, arity, args);
		}

		// 参数收集完毕，执行func
		return func.apply(func, args);
	}

	return wrapper;
}

var _check = createCurry(check);
var checkPhone = _check(phoneNumberReg);

var _check = createCurry(check);
var checkEmail = _check(emailReg);

var phoneStatus = checkPhone('18388888888');
var emailStatus = checkEmail('xxxxxx@test.com');
console.log('phone status: ', phoneStatus);
console.log('email status: ', emailStatus);
