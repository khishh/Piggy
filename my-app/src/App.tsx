import React, { useContext, useEffect, useState } from 'react';
import './App.css';
import LoginButton from './components/LoginButton';
import LogoutButton from './components/LogoutButton';
import TestProfile from './components/TestProfile';
import { useAuth0 } from '@auth0/auth0-react';
import { useMutation } from '@apollo/client';
import { User } from './graphql/model/User';
import { CREATE_USER_ON_SIGN_IN } from './graphql/mutation/User';
import TestPlaid from './components/TestPlaid';
import Link from './components/Link';
import plaidContext from './context/PlaidContext';
import Onboard from './components/Onboard';
import PlaidLink from './components/PlaidLink';
import SideBar from './components/SideBar';
import TopBar from './components/TopBar';
import { AppPageName } from './enums/AppPageName';
import Balance from './components/Balance';
import Transactions from './components/Transactions';
import Dashboard from './components/Dashboard';
import Transfer from './components/Transfer';



function App() {

  const { user, isAuthenticated } = useAuth0();
  const [createUserOnSignIn, { loading, error, data }] = useMutation<{
    CreateUserOnSignIn: User
  }>(CREATE_USER_ON_SIGN_IN);

  const [currentPageName, setCurrentPageName] = useState<AppPageName>(AppPageName.DASHBOARD);

  useEffect(() => {
    console.log(currentPageName);
    
  }, [currentPageName])
  

  const { accessToken, dispatch } = useContext(plaidContext);

  const showCurrentPage = (currentPageName: AppPageName) => {
    switch (currentPageName) {
      case AppPageName.BALANCE:
        return <Balance />
      case AppPageName.TRANSACTIONS:
        return <Transactions />
      case AppPageName.TRANSFER:
        return <Transfer />
      default:
        return <Dashboard />
    }
  }

  useEffect(() => {
    if (isAuthenticated) {
      console.log(`user logs in ${user}`);
      // Only create user if this user has not sign in from this application yet.
      createUserOnSignIn({
        variables: {
          input: {
            first_name: user?.given_name,
            last_name: user?.family_name,
            picture: user?.picture,
            email: user?.email,
            sub: user?.sub
          }
        }
      }).then((data) => {
        console.log("createUserOnSignIn completed");
        console.log(data)
        // case when the access token for this user is available in the DB
        if (data.data?.CreateUserOnSignIn.access_token) {
          dispatch({
            type: "SET_STATE",
            state: {
              accessToken: data.data?.CreateUserOnSignIn.access_token
            }
          });
        }
      })
    }
  }, [isAuthenticated])


  return (
    <div className="app">


      <div className='sidebar-wrapper'>
        <SideBar setCurrentPageName={setCurrentPageName} currentPageName={currentPageName} />
      </div>

      <div className='main-wrapper'>
        <div className='topbar-wrapper'>
          <TopBar />
        </div>
        <div className='app-wrapper'>
          {/* {!isAuthenticated && <Onboard />}

          {isAuthenticated && <LogoutButton />}

          {isAuthenticated && accessToken === "" && <PlaidLink />}

          {isAuthenticated && accessToken !== "" && <TestPlaid />} */}

          {
            showCurrentPage(currentPageName)
          }
        </div>
      </div>
      


    </div>
  );
}

export default App;
