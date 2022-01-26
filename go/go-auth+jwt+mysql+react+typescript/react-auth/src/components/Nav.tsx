import React from "react";
import { Link, Navigate } from "react-router-dom";

const Nav = (props: { name: string, setName: (name: string) => void }) => {
    // const [redirect, setRedirect] = useState(false);

    const logout = async () => {
        const response = await fetch('http://127.0.0.1:8000/api/logout', {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        });

        // if (!response.ok) {
        //     return false
        // }

        const content = await response.json();
        console.log(content);

        props.setName('');

        // setRedirect(true)

        return (<Navigate to="/" />);

        // setName(content.name)
    };


    let menu;
    if (props.name === '') {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                    <Link className="nav-link active" aria-current="page" to="/login">Login</Link>
                </li>
                <li className="nav-item">
                    <Link className="nav-link active" aria-current="page" to="/register">Register</Link>
                </li>
            </ul>
        );
    } else {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                    <Link className="nav-link active" aria-current="page" to="/logout"
                        onClick={logout}
                    >Logout</Link>
                </li>
            </ul>
        );
    }

    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link className="navbar-brand" to="/">Home</Link>

                <div className="collapse navbar-collapse" id="navbarCollapse">
                    {menu}
                </div>
            </div>
        </nav>
    );
};

export default Nav;