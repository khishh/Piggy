import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { Auth0Provider } from '@auth0/auth0-react';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

if (!process.env.REACT_APP_AUTH0_DOMAIN || !process.env.REACT_APP_AUTH0_CLIENT_ID) {
  root.render(
    <h1>Need to set up Auth0 setup in .env</h1>
  )
} else {
  root.render(
    <Auth0Provider
      domain={process.env.REACT_APP_AUTH0_DOMAIN}
      clientId="T2V3aqcr8xETmNpIAeWm8sEbklDTby3Z"
      redirectUri={window.location.origin}
    >
      <App />
    </Auth0Provider>
  );
}



// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();