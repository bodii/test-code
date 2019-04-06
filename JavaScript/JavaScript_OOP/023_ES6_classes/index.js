/**
 * 类
 * javascript中的类，只是对原有对象的一种封装或者说外衣
 */

// function Circle(radius) {
//     this.radius = radius;
    
//     this.draw = function() {
//         console.log('draw');
//     }
// }

class Circle {
    constructor(radius) {
        this.radius = radius;

        // this.draw = function() {
        //     console.log('draw');
        // }  //也可以这样，但不推荐
    }

    draw() {
        console.log('draw');
    }
}

const c = new Circle(1);