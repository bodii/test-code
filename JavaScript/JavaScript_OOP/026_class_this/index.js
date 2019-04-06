/**
 * 类的this对象
 */

//  'use strict';  // 使用严格模式

const Circle = function() {
    this.draw = function() {
        console.log(this);
    }
}

const c = new Circle();
// class method call
c.draw();

// const draw = c.draw;
// // window(or Global) method call
// draw();   // 如果是在严格模式下，返回undefined, 因为严格模式下window(or Global) 是undefined


class Square {
    draw() {
        console.log(this);
    }
}

const s = new Square();
const draw = s.draw;
draw(); // undefined，类声明下默认启用严格模式