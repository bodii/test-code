declare namespace JSX {
    interface IntrinsicElements {
        foo: { requiredProp: string; optionalProp?: number }
    }
}


<foo requiredProp='bar' />; // 正确
<foo requiredProp='bar' optionalProp={0} />; // 正确
// <foo />; // 错误，缺少requiredProp
// <foo requiredProp={0} />; // err, requiredProp应该是字符串
// <foo requiredProp="bar" unknownProp />; // err, unknowProp不存在
<foo requiredProp='bar' some-unknown-prop />; // 正确，some-unknow-prop不是个合法的标识符
