// 对象的成员的获取和遍历的区别
function Circle(radius) {
    // Instance members
    this.radius = radius;

    this.move = function() {
        console.log('move');
    }
}

// 实例对象
const c1 = new Circle(1);
 
// 定义原型方法
// prototype members
Circle.prototype.draw = function() {
    console.log('draw');
}

//Object.keys() 返回的是实例的成员
console.log('keys: ', Object.keys(c1));

console.log('loop: ');
// for 返回的是原型的成员
for (let key in c1) console.log(key);

console.log(c1.hasOwnProperty('radius'));
console.log(c1.hasOwnProperty('draw'));