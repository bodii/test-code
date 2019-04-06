class Circle {
    constructor(radius) {
        this.radius = radius;
    }

    draw() {
        console.log('draw');
    }

    // static method
    static parse(str) {
        const radius =  JSON.parse(str).radius;
        return new Circle(radius);
    }

}

// instance a object
const c = new Circle(1);
// c.parse('b'); // error，静态方法不能通过实例对象调用

// call class static method
const circle = Circle.parse('{ "radius": 1}');
