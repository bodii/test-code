declare namespace JSX {
    interface IntrinsicElements {
        foo: { bar?: boolean }
    }
}

// foo的元素属性类型为{bar?: boolean}
<foo bar />;