function Circle(radius) {
    this.radius = radius;
    this.draw = function() {
        console.log('draw');
    }
}

const circle = new Circle(10);

// setter 
circle.location = { x: 1 };

const propertyName = 'location2';
circle[propertyName] = { x: 1 };

// remove
delete circle.location;