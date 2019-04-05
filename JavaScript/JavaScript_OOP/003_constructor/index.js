// Constructor Function
function Circle(radius) {
    console.log('this', this);

    this.radius = radius;
    this.draw = function() {
        console.log('draw');
    }
}

const another = new Circle(1);