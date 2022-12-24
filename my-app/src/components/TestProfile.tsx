import { useAuth0 } from '@auth0/auth0-react'
import React from 'react'

const TestProfile = () => {

    const { user, isAuthenticated, isLoading } = useAuth0();

    if (isLoading) {
        return <div>Loading...</div>
    } else if (!isAuthenticated) {
        return <div>Not AUthenticated yet</div>
    } else if (!user) {
        return <div>No user</div>
    }

    console.log(user)

    return (
        isAuthenticated && user && (
            <div>
                <img src={user.picture} alt={user.name} />
                <h2>{user.name}</h2>
                <p>{user.email}</p>
            </div>
        )
    );
}

export default TestProfile