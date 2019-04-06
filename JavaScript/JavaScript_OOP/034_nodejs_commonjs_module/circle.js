// Implementation Detail 
const _radius = new WeakMap();

// public Interface
class Circle {
    constructor(radius) {
        _radius.set(this, radius);
    }

    draw() {
        console.log('Circle with radius ' + _radius.get(this));
    }
}

module.exports = Circle;