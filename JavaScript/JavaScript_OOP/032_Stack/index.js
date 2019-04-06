const _items = new WeakMap();  // WeakMap 弱映射

class Stack {

    constructor() {
        _items.set(this, []);
    }

    peek() {
        const items = _items.get(this);
        if (items.length <= 0) throw new Error('Stack is empty.');

        return items[this.count -1];
    }

    push(item) {
        _items.get(this).push(item);
    }

    pop() {
        const items = _items.get(this);
        if (items.length <= 0) throw new Error('Stack is empty.');

        return items.pop();
    }

    get count() {
        return _items.get(this).length;
    }
}


const stack = new Stack();