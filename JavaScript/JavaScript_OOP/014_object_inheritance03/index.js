/**
 * 对象的多继承
 * 
 */

function Shape(color) {
    this.color = color;
}

Shape.prototype.duplicate = function() {
    console.log('duplicate');
}

function Circle(radius, color) {
    Shape.call(this, 'color');

    this.radius = radius;
}

Circle.prototype.draw = function() {
    console.log('draw');
}

Circle.prototype  = Object.create(Shape.prototype);
Circle.prototype.constructor = Circle;


function Square(size, color, radius) {

    Circle.call(this, color, radius);

    this.size = size;
}

Square.prototype = Object.create(Circle.prototype);
Square.prototype.constructor = Square;

const s = new Shape('red');
const c = new Circle(1, 'pink');
const sq = new Square(20, 'black', 3);
