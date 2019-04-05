function Shape(color) {
    this.color = color;
}

Shape.prototype.duplicate = function() {
    console.log('duplicate');
}

function Circle(radius, color) {

    Shape.call(this, color); // 将Shape的构造方法作用于本对象

    this.radius = radius;
}

Circle.prototype = Object.create(Shape.prototype);
Circle.prototype.constructor = Circle;

Circle.prototype.draw = function() {
    console.log('draw');
}

const c = new Circle(1, 'pink');