class Shape {
    move() {
        console.log('move');
    }
}


class Circle extends Shape {
    move() {
        super.move(); // 当需要调用父类的方法时，使用super关键字
        console.log('circle move');
    }
}


const c = new Circle();