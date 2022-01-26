import React from "react";

const Home = (props: { name: string }) => {


    return (
        <div>
            {props.name ? 'Hi ' + props.name : 'You are not login'}
        </div>
    );
}

export default Home;