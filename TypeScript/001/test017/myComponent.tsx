class MyComponent {
    render() {}
}

// 使用构造签名
var myComponent = new MyComponent();

// 元素类的类型 =》 MyComponent
// 元素实例的类型 => { render: () => void }
function MyFactoryFunction() {
    return {
        render: () => {}
    }
}

// 使用调用签名
var myComponent = MyFactoryFunction();

// 元素类的类型 => MyFactoryFunction
// 元素实例的类型 => { render: () => void }

