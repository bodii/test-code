/**
 * 继承方法重写
 */

function Shape() {

}

Shape.prototype.duplicate = function() {
    console.log('duplicate')
}

function Circle(radius) {
    this.radius = radius
}

Circle.prototype.draw = function() {
    console.log('draw');
}

function extend(Child, Parent) {
    Child.prototype = Object.create(Parent.prototype);
    Child.prototype.constructor = Child;
}

extend(Circle, Shape);

// 对父方法的重写
Circle.prototype.duplicate = function() {
    Shape.prototype.duplicate.call(this); // 子方法中调用父方法
    console.log('duplicate circle');
}


function Square() {

}

extend(Square, Shape);

Square.prototype.duplicate = function() {
    console.log('duplicate square');
}

const c = new Circle(1);