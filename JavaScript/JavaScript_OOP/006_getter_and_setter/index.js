function Circle(radius) {
    this.radius = radius;

    let defaultLocation = { x: 0, y: 0 };

    this.getDefaultLocalt = function() {
        return defaultLocation;
    };

    Object.defineProperty(this, 'defaultLocation', {
        get: function() {
            return defaultLocation;
        },
        set: function(value) {
            if (!value.x || !value.y)
                throw new Error('Invalid location.');
            defaultLocation = value;
        }
    });

    this.draw = function() {
        console.log('draw');
    };
}

const circle = new Circle(10);
circle.getDefaultLocalt();
circle.draw();