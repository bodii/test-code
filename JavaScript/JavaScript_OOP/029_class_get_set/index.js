const _radius = new WeakMap();

class Circle {
    constructor(radius) {
        _radius.set(this, radius);

    //     Object.defineProperty(this, 'radius', {
    //         get: function() {
    //             return _radius.get(this);
    //         },
    //         set: function(value) {
    //            _radius.set(this, value)
    //         },
    //     });
    // }

    // getRadius() {
    //     return _radius.get(this);
    // }
    }

    // 等价于
    get radius() {
        return _radius.get(this);
    }

    set radius(value) {
        if(value <= 0) throw new Error('invalid radius');
        
        _radius.set(this, value);
    }
}

const c = new Circle(1);