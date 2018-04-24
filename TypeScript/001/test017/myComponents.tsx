declare namespace JSX {
    interface ElementClass {
        render: any;
    }
}

class MyComponents {
    render() {}
}

function MyFactoryFunctions() {
    return { render: () => {} }
}

<MyComponents />; // 正确
<MyFactoryFunctions />; // 正确

class NotAValidComponent {}
function NotAValidFactoryFunction() {
    return {};
}

// <NotAValidComponent />; // 错误 ，缺少render属性
// <NotAValidFactoryFunction />; // 错误， 