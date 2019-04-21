(function(ROOT) {

	// 构造函数
	var jQuery = function(selector) {
		// 在该方法中直接返回new创建的实例
		// 因此这里的init才是真正的构造函数
		return new jQuery.fn.init(selector);
	}

	jQuery.fn = jQuery.prototype = {
		constructor: jQuery,
		version: '1.0.0',
		init: function(selector) {
			var elem, selector;
			elem = document.querySelector(selector);
			this[0] = elem;

			// 在jQuery中返回的是一个由所有原型属性方法组成的数组，
			// 这里做了简化，直接反回this即可
			return this;
		},

		toArray: function() {},
		get: function() {},
		each: function() {},
		ready: function() {},
		first: function() {},
		slice: function() {}
	}

	// init的原型指向jQuery.prototype
	jQuery.fn.init.prototype = jQuery.fn;

	// 实现jQuery的两种扩展方法
	jQuery.extend = jQuery.fn.extend = function(options) {
		var target = this;
		var copy;
		for (name in options) {
			copy = options[name];
			target[name] = copy;
		}
		return target;
	}

	// 添加静态扩展方法，即工具方法
	// 静态方法调用： $.ajax()
	jQuery.extend({
		isFunction: function() (),
		type: function() {},
		perseHTML: function() {},
		perseJSON: function() {},
		ajax: function() {}
	});


	// 添加原型方法
	// 原型方法调用$('#test').css()
	jQuery.fn.extend({
	
		query: function() {},
		promise: function() {},
		attr: function() {},
		prop: function() {},
		addClass: function() {},
		removeClass: function() {},
		val: function() {},
		css: function() {}
	});

	// 对外暴露的接口
	ROOT.jQuery = ROOT.$ = jQuery;
})(window);
