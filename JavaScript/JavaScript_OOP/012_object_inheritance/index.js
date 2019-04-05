/**
 * 原型继承
 */

// base Shape object
function Shape()  {

}

// 定义Shape Objec 原型方法
Shape.prototype.duplicate = function() {
    console.log('duplicate');
}


// 创建Circle object
function Circle(radius) {
    this.radius = radius;
}

// Circle.prototype = Object.create(Object.prototype); // 默认继承根对象(Object)的原型
// 将Shape object prototype base object  赋值给 Circle 的 base object
Circle.prototype = Object.create(Shape.prototype);
Circle.prototype.constructor = Circle;
/**
 * -----------------------------------------------------------------------
 * 当你重设原型属性，作为最佳实践应确保也重设了构造器
 * -----------------------------------------------------------------------
 */

Circle.prototype.draw = function() {
    console.log('draw');
}


const s = new Shape();
const c = new Circle(1);
