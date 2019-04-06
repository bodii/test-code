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

function Person() {

}

Object.assign(Person.prototype, canEat, canWalk);

const person = new Person();