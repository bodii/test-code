class Shape {
    constructor(color) {
        this.color = color;
    }

    move() {
        console.log('move');
    }
}

class Circle extends Shape {
    constructor(color, radius) {
        super(color); // 父类的构造函数

        this.radius = radius;
    }
    draw() {
        console.log('draw');
    }
}

const c = new Circle('red', 1);
