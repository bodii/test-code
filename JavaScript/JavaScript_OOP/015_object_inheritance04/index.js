function Shape(color) {
    this.color = color;
}

Shape.prototype.duplicate = function() {
    console.log('duplicate');
}

function Circle(radius, color) {

    Shape.call(this, color);

    this.radius = radius;
}

Circle.prototype.draw = function() {
    console.log('draw');
}

// 自定义继承方法
function extend(Child, Parent) {
    Child.prototype = Object.create(Parent.prototype);
    Child.prototype.constructor = Child;
}

extend(Circle, Shape);

const s = new Shape('black');
const c = new Circle(1, 'pink');


function Square(size, color) {
   
    Shape.call(this, 'color');
    
    this.size = size;
}

extend(Square, Shape);

const sq = new Square(5, 'green');