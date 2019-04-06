/**
 * 组合胜过继承，过多的层级继承会让程序更难维护，建议不要超过一层
 */

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

const person = Object.assign({}, canEat, canWalk);
console.log(person);