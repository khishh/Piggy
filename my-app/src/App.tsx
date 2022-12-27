import React, { useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import LoginButton from './components/LoginButton';
import LogoutButton from './components/LogoutButton';
import TestProfile from './components/TestProfile';
import { useAuth0 } from '@auth0/auth0-react';

function App() {

  const { user, isAuthenticated } = useAuth0();

  useEffect(() => {
    if (user) {
      console.log(`user logs in ${user}`);  

    }
  }, [isAuthenticated])
  

  return (
    <div className="App">
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}
      <LoginButton/>
      <LogoutButton/>
      <TestProfile/>
    </div>
  );
}

export default App;
