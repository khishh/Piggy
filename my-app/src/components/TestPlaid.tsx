
import { useAuth0 } from '@auth0/auth0-react'
import React, { useContext, useEffect } from 'react'
import plaidContext from '../context/PlaidContext';
import Link from './Link';

const TestPlaid = () => {

    const { user, isAuthenticated } = useAuth0();
    const { linkToken, dispatch } = useContext(plaidContext)

    // console.log("TestPlaid Component render " + linkToken);

    const postData = {
        SubId: user?.sub
    }

    // useeffect hook to request link token from the backend
    useEffect(() => {
        if (isAuthenticated) {
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

    useEffect(() => {
        if (linkToken) {
            console.log(`Link Token has been updated! ${linkToken}`);
        }


    }, [linkToken])

    return (
        <>
            <div>TestPlaid</div>
            <p>{linkToken}</p>
        </>
    )
}

export default TestPlaid