const canEat = {
    eat: function() {
        this.hunger--;
        console.log('eating');
    },
};

const canWalk = {
    walk: function() {
        console.log('walking');
    },
};

const canSwim = {
    swim: function() {
        console.log('swim');
    },
};

// 封装对象拷贝
function mixin(target, ...sources) {
    Object.assign(target, ...sources);
}


function Person() {

}

mixin(Person.prototype, canEat, canWalk);
const person = new Person();

function Goldfish() {

}

mixin(Goldfish.prototype, canEat, canSwim);
const goldfish = new Goldfish();