function Shape() {

}

Shape.prototype.duplicate = function() {
    console.log('duplicate');
}

function extend(Child, Parent) {
    Child.prototype = Object.create(Parent.prototype);
    Child.prototype.constructor = Child;
}

function Circle() {

}

extend(Circle, Shape);

Circle.prototype.draw = function() {
    console.log('draw');
}

Circle.prototype.duplicate = function() {
    console.log('duplicate circle');
}

function Square() {

}

extend(Square, Shape);
Square.prototype.duplicate = function() {
    console.log('duplicate square');
}

const shapes = [
    new Circle(),
    new Square(),
];

for (let shape of shapes) {
    shape.duplicate()
}