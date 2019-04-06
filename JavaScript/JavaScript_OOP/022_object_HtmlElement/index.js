function HtmlElement() {
    this.render = function() {
        
    }
}

function HtmlSelectElement(items = []) {
    this.items = items;

    let hasItemEmpty = function(item) {
        if (!item)
            throw new Error('item can\'t empty');
    }

    this.addItem = function(item) {
        hasItemEmpty(item);

        this.items.push(item);
    }

    this.removeItem = function(item) {
        hasItemEmpty(item);

        this.items.splice(this.items.indexOf(item), 1);
    }

    this.render = function() {
        // let htmls = '<select>\n';
        // for (let element of this.items) {
        //     htmls += '    <option>' + element + '</option>\n';
        // }
        // htmls += '</select>';

        // return htmls;

        // 等价：
        return `
<select>
    ${this.items.map(item => ` <option>${item}</option>`).join('')}
 </select>`;
    }
}

function HtmlImageElement(src) {
    this.src = src;

    this.render = function() {
        if (!this.src)
            throw new Error('src can\'t empty.');
        
        // return '<img src="' + this.src + '" />'; // 等价
        return `<img src="${this.src}" />`;
    }
}

function extend(Child, Parent) {
    Child.prototype = Object.create(Parent.prototype);
    Child.prototype.constructor = Child;
}

extend(HtmlSelectElement, HtmlElement);

extend(HtmlImageElement, HtmlElement);