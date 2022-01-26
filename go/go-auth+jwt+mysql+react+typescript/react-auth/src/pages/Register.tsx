import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        console.log({
            name,
            email,
            password
        })


        const response = await fetch('http://127.0.0.1:8000/api/register', {
            method: "POST",
            headers: { 'Content-Type': "application/json" },
            body: JSON.stringify({
                name,
                email,
                password
            })
        })

        if (!response.ok) {
            return false
        }


        setRedirect(true)

        const content = await response.json();
        console.log(content)
    };

    if (redirect) {
        return (<Navigate to="/login" />);
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please register</h1>
            <div className="form-floating">
                <input type="name" className="form-control" id="floatingInput" placeholder="name" required
                    onChange={e => setName(e.target.value)}
                />
                <label htmlFor="floatingInput">Name</label>
            </div>
            <div className="form-floating">
                <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com" required
                    onChange={e => setEmail(e.target.value)}
                />
                <label htmlFor="floatingInput">Email address</label>
            </div>
            <div className="form-floating">
                <input type="password" className="form-control" id="floatingPassword" placeholder="Password" required
                    onChange={e => setPassword(e.target.value)}
                />
                <label htmlFor="floatingPassword">Password</label>
            </div>

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    );
}

export default Register;