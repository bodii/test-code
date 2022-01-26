import React, { SyntheticEvent, useState } from 'react'
import { Navigate } from 'react-router-dom';

const Login = (props: { setName: (name: string) => void }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        console.log({
            email,
            password
        })

        const response = await fetch('http://127.0.0.1:8000/api/login', {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify({
                email,
                password
            })
        })

        if (!response.ok) {
            return false
        }

        const content = await response.json()
        console.log(content)

        props.setName(content.name)

        setRedirect(true)
    }

    if (redirect) {
        return (<Navigate to="/" />);
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please login</h1>
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

            <button className="w-100 btn btn-lg btn-primary" type="submit">submit</button>
        </form>
    );
}

export default Login;