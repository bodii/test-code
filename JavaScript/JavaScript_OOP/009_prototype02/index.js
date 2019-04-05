// 定义原型属性和方法
function Circle(radius) {
    // Instance members
    this.radius = radius;
}

// 定义原型方法
// prototype members
Circle.prototype.draw = function() {
    console.log('draw');
}

const c1 = new Circle(1);
const c2 = new Circle(2);

Circle.prototype.toString = function() {
    return 'Circle with radius ' + this.radius;
}
Circle.prototype.memberAttr = 1;