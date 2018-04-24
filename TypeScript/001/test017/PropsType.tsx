declare namespace JSX {
    interface Element {
        [elemeName: string]: any;
    }
} 
interface PropsType {
    children: JSX.Element
    name: string
}

class Component extends React.Component<PropsType, {}> {
    render() {
        return (
            <h2>
                this.props.children
            </h2>
        )
    }
}

// <Component>
//     <h1>Hello World</h1>
// </Component>

// <Component>
//     <h1>Hello World</h1>
//     <h2>Hello World</h2>
// </Component>

