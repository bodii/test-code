import React, { useState, useEffect } from 'react';
import './App.css';
import Login from './pages/Login';
import Home from './pages/Home';
import Register from './pages/Register';
import Nav from './components/Nav';
import { BrowserRouter, Route, Routes } from "react-router-dom"

function App() {

  const [name, setName] = useState('');

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://127.0.0.1:8000/api/user', {
          headers: { "Content-Type": "application/json" },
          credentials: "include",
        });

        if (!response.ok) {
          return false
        }

        const content = await response.json()
        console.log(content)

        setName(content.name)
      }
    )();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />
        <main className="form-signin">
          <Routes>
            <Route path="/" element={<Home name={name} />} />
            <Route path="/login" element={<Login setName={setName} />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
