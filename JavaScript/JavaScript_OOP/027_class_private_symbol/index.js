/**
 * 通过Symbol()创建私有属性和方法
 */

const _radius = Symbol();  // Symbol是一个方法，每次调用会新成新值
// console.log(Symbol() === Symbol());  // false
const _draw = Symbol();

class Circle {
    constructor(radius) {
        // this.radius = radius;
        // ==等价
        // this['radius'] = radius;

        this[_radius] = radius;
    }

    [_draw]() {
        console.log(this);
    }
}

const c = new Circle(1);
// const key = Object.getOwnPropertyNames(c);
const radius_key = Object.getOwnPropertySymbols(c)[0];
console.log(c[radius_key]);
