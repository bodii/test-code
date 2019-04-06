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
        console.log('swiming');
    },
};

function Person() {

}

Object.assign(Person.prototype, canEat, canWalk);


function Goldfish() {

}
Object.assign(Goldfish.prototype, canEat, canSwim);

const g = new Goldfish();