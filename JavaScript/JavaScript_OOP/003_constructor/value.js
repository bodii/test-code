let number = 10;

function increase(number) {
    number++;
}

increase(number);
console.log('value ', number);


let object = { value: 10 }
function objectValueIncrease(obj) {
    obj.value++;
}

objectValueIncrease(object);
console.log('object is value:', object.value)