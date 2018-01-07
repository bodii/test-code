/**
 * ------------------------------
 *   纯JavaScript实现可移动式悬浮窗
 * ------------------------------
 * 支持移动端和pc端
 * 
 *  example:
 *  实例化一个拖动悬浮窗
 * 	:parameter bDarg string element id
 * 	:parameter opacity integer(0-100): 拖动时的透明度
 * 	:parameter flag boolean: 标记
 *  Drag.init({bDarg:"float_feedbox",opacity: 50, flag: true});
 */
var Drag = (function(window){
	return {
		// 初始化拖动元素并绑定鼠标或手指拖动事件
		init: function(options) {
			var defaults=new Object();
			for(elements in this.defaults){
				defaults[elements]=this.defaults[elements]
			}
			defaults.bDarg=document.getElementById(options.bDarg);
			defaults.opacity=options.opacity; //拖动时的透明度
			defaults.flag = options.flag;
			defaults.isMobile = Drag.isMobile(); // 是否是移动端
			Drag.setCss.call(defaults.bDarg,{"position":"absolute","cursor":"move"});
			var lib = this;
			var mousedown = "mousedown", mousemove = "mousemove", mouseup = "mouseup",
				onmousedown = "onmousedown", onmousemove = "onmousemove", onmouseup = "onmouseup";
			if(defaults.isMobile) {
				mousedown = "touchstart";
				mousemove = "touchmove";
				mouseup = "touchend";
				onmousedown = "ontouchstart";
				onmousemove = "ontouchmove";
				onmouseup = "ontouchend";
			}
			if(!!defaults.bDarg.addEventListener){
				defaults.bDarg.addEventListener(mousedown,function(e){
					lib.beforeDrag.call(defaults.bDarg,e,defaults);
					window.addEventListener(mousemove,function(e){
						lib.onDrag.call(defaults.bDarg,e,defaults)
					},false);
					window.addEventListener(mouseup,function(e){
						lib.endDrag.call(defaults.bDarg,e,defaults)
					},false);
				},false);
			}else{
				defaults.bDarg.attachEvent(onmousedown,function(e){
					lib.beforeDrag.call(defaults.bDarg,e,defaults);
					document.attachEvent(onmousemove,function(e){
						lib.onDrag.call(defaults.bDarg,e,defaults)
					});
					document.attachEvent(onmouseup,function(e){
						lib.endDrag.call(defaults.bDarg,e,defaults)
					});
				});
			}
		},
		// 初始化拖动元素参数
		defaults:{
			bDarg:null,
			opacity: 100,
			flag: true,
			elem:[],
			dowOrUp:false
		},
		// 设置拖动元素的样式
		setCss: function(cObj){
			for(var c in cObj){
				this.style[c]=cObj[c];
			}
		},
		// 设置拖动元素的透明度
		setOpacity: function(opacity){
			if(window.addEventListener){
				this.style.opacity = opacity / 100;
			}else{
				this.filters.alpha.opacity = opacity;
			}
		},
		// 鼠标或手指按下时
		beforeDrag: function(e,objdefaule){
			var e = e || window.event;
			if(objdefaule.elem.length>1) return ;
			if( objdefaule.isMobile ||
				(window.addEventListener && e.button==0) ||
				(!window.addEventListener && e.button==1)
			){
				if (e.preventDefault) {
					e.preventDefault();
					e.stopPropagation();
				}
				objdefaule.dowOrUp=true;
				var clone=this.cloneNode(true);
				clone.style.zIndex=Drag.getZIndex();
				this.style.zIndex=Drag.getZIndex()+1;
				clone.removeAttribute("id");
				this.parentNode.appendChild(clone);
				Drag.setOpacity.call(this,objdefaule.opacity);
				objdefaule.elem.push({
					lib:this,
					obj:clone,
					startX:e.pageX || e.touches[0].pageX || e.clientX,
					startY:e.pageY || e.touches[0].pageY || e.clientY,
					startLeft:this.offsetLeft,
					startTop:this.offsetTop
				});
			}
		},
		// 拖动时方法
		onDrag: function(e,objdefaule){
			var e = e || window.event;
			if(!objdefaule.isMobile && !window.addEventListener){
				e.returnValue = false;
			}
			if(objdefaule.dowOrUp){
				for(var i=0; i<objdefaule.elem.length; i++){
					var pagex = null;
					var pagey = null;
					if(e.touches){
						pagex = e.touches[0].pageX;
						pagey = e.touches[0].pageY
					} else {
						pagex = e.pageX || e.clientX;
						pagey = e.pageY || e.clientY;
					}
					objdefaule.elem[i].lib.style.left=pagex-objdefaule.elem[i].startX+objdefaule.elem[i].startLeft+ "px";
					objdefaule.elem[i].lib.style.top =pagey-objdefaule.elem[i].startY+objdefaule.elem[i].startTop +"px";
				}
			}
		},
		// 拖动结束方法
		endDrag: function(e,objdefaule){
			if(objdefaule.dowOrUp){
				var i=0;
				while(objdefaule.elem.length>0){
					objdefaule.elem[i].obj.style.left=objdefaule.elem[i].lib.style.left;
					objdefaule.elem[i].obj.style.top =objdefaule.elem[i].lib.style.top;
					Drag.setOpacity.call(Drag.defaults.elem[i].lib,100)
					objdefaule.elem[i].obj.parentNode.removeChild(objdefaule.elem[i].obj)
					objdefaule.elem.shift();
				}
				Drag.defaults.dowOrUp=false;
			}
		},
		// 获取最上层的图层数
		getZIndex: function(){
			var zindex = 1;
			var div = document.getElementsByTagName("div");
			for(var i = 0 ;i < div.length; i++){
				if(parseInt(zindex)<parseInt(div[i].style.zIndex)){
					zindex=parseInt(div[i].style.zIndex);
				}
			}
			return parseInt(zindex);
		},
		// 判断是否是移动端
		isMobile: function(){
			if (/Mobile/ig.test(navigator.userAgent)) {
				return true;
			}
			return false;
		},
	}
})(window);

// 实例化拖动对象悬浮窗
window.onload=function(){
	Drag.init({bDarg:"float_feedbox",opacity: 50, flag: true});
};