const _radius = new WeakMap();
const _move = new WeakMap();
const privateProps = new WeakMap();

class Circle {
    constructor(radius) {
        // _radius.set(this, radius);

        // // 箭头函数内的this，指向包含它的对象
        // _move.set(this, () => {
        //     console.log('move', this);
        // });
        // _move.set(this, function() {
        //     console.log('move', this); // 这里获取不到this对象
        // });

        // 等价于
        privateProps.set(this, {
            radius: radius,
            move: () => {
                console.log('move', this);
            }
        });  // 推荐这种，还是单个写

    }

    draw() {
        // console.log(_radius.get(this));
        // _move.get(this)();


        // privateProps.get(this).move();
        //or
        privateProps.get(this)['move']();

        console.log('draw');
    }
}

const c = new Circle(1);
c.draw(); 