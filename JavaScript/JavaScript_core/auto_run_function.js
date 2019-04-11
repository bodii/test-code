// 自执行创建模块
(function() {
	var states = {}; // 私有变量，用来存储状态与数据

	// 判断数据类型
	function type(elem) {
		if (elem == null) {
			return elem + '';
		}

		return toString.call(elem).replace(/[\[\]]/g, '').split(' ')[1].toLowerCase();
	}

	/**
	 * @param name 属性名
	 * @description 通过属性名获取保存在states中的值
	 */
	function get(name) {
		return states[name] ? states[name] : '';
	}

	function getStates() {
		return states;
	}

	/**
	 * @param options {object} 键值对
	 * @param target {object} 属性值为对象的属性，只在函数实现，递归调用时传入
	 * @desc 通过传入键值对的方式修改state树，使用方式与小程序的data或react中的
	 * setStates类似
	 */
	function set(options, target) {
		var keys = Object.keys(options);
		var o = target ? target : states;

		keys.map(function(item) {
			if (typeof o[item] == 'undefine') {
				o[item] = options[item];
			}
			else {
				type(o[item]) == 'object' ? 
					set(options[item], o[item]) : 
					o[item] = options[item];
			}

			return item;
					
		});
	}

	// 对外提供接口
	window.get = get;
	window.set = set;
	window.getStates = getStates;
})();


// 具体使用方式如下
set({ a: 20 }); // 保存属性a
set({ b: 100 });
set({ c: 10 });

// 保存属性o,它的值为一个对象
set({
	o: {
		m: 10,
		n: 20
	}
});

// 修改对象o的m值
set({
	o: {
		m: 1000
	}
});

// 给对象o中增加一个C属性
set({
	o: {
		c: 100
	}
});

console.log(getStates());
