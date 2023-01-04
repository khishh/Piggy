import React, { useContext, useEffect } from 'react'
import Link from './Link'
import { useAuth0 } from '@auth0/auth0-react';
import plaidContext from '../context/PlaidContext';

const PlaidLink = () => {

    const { user, isAuthenticated } = useAuth0();
    const { linkToken, accessToken, dispatch } = useContext(plaidContext)

    // useeffect hook to request link token from the backend
    useEffect(() => {
        console.log(`===== ${accessToken} ====`);

        const postData = {
            SubId: user?.sub
        }

        if (isAuthenticated && accessToken === "") {
            const response = fetch("http://localhost:80/api/create_link_token", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(postData)
            }).then(response => {
                if (!response.ok) {
                    return;
                }
                return response.json();
            }).then(json => {
                console.log(json);
                dispatch({
                    type: "SET_STATE",
                    state: {
                        linkToken: json.link_token
                    }
                });
            })
        }
    }, [isAuthenticated])


    return (
        <>
            <h1>Connect your financial instituition via Plaid API!</h1>
            <Link />
        </>
    )
}

export default PlaidLink