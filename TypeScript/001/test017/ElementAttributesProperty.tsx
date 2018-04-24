declare namespace JSX {
    interface ElementAttributesProperty {
        props; // 指定用来使用的属性名
    }
}

class MyComponent {
    // 在元素实例类型上指定属性
    props: {
        foo?: string;
    }
}

// MyComponent的元素属性类型为{foo?: string}
<MyComponent foo='bar' />