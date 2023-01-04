import React from 'react'
import '../styles/TopBar.css'
import { useAuth0 } from '@auth0/auth0-react'
import LoginButton from './LoginButton';
import LogoutButton from './LogoutButton';

const TopBar = () => {

    const { user, isAuthenticated } = useAuth0();

    return (
        <div className='topbar'>

            {
                isAuthenticated &&
                <div className='topbar-logout-btn'>
                    <LogoutButton/>
                </div>

            }
            {
                !isAuthenticated &&
                <div className='topbar-login-btn'>
                    <LoginButton />
                </div>
            }

        </div>
    )
}

export default TopBar