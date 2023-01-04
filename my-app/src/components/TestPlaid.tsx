
import { useAuth0 } from '@auth0/auth0-react'
import React, { useContext, useEffect } from 'react'
import plaidContext from '../context/PlaidContext';
import Link from './Link';
import { useMutation } from '@apollo/client';
import { UPDATE_USER_WITH_ACCESS_TOKEN } from '../graphql/mutation/User';
import { User } from '../graphql/model/User';

const TestPlaid = () => {

    const { user, isAuthenticated } = useAuth0();
    const { linkToken, accessToken, dispatch } = useContext(plaidContext)

    const [ updateUserWithAccessToken, { loading, error, data} ] = useMutation<{
        UPDATE_USER_WITH_ACCESS_TOKEN: User
    }>(UPDATE_USER_WITH_ACCESS_TOKEN);

    // const postData = {
    //     SubId: user?.sub
    // }

    // // useeffect hook to request link token from the backend
    // useEffect(() => {
    //     console.log(`===== ${accessToken} ====`);
        
    //     if (isAuthenticated && accessToken === "") {
    //         const response = fetch("http://localhost:80/api/create_link_token", {
    //             method: "POST",
    //             headers: {
    //                 "Content-Type": "application/json"
    //             },
    //             body: JSON.stringify(postData)
    //         }).then(response => {
    //             if (!response.ok) {
    //                 return;
    //             }
    //             return response.json();
    //         }).then(json => {
    //             console.log(json);
    //             dispatch({
    //                 type: "SET_STATE",
    //                 state: {
    //                     linkToken: json.link_token
    //                 }
    //             });
    //         })
    //     }
    // }, [isAuthenticated])



    useEffect(() => {
        if (accessToken !== "") {
            console.log(`Access Token has been updated! ${accessToken}`);
            // save accessToken to database
            const updateUserAccessToken = async () => {
                console.log("==== updateUserAccessToken called ====");
                console.log(user?.sub);
                console.log(accessToken);
                
                const response = await updateUserWithAccessToken(
                    {
                        variables: {
                            auth0_sub_id: user!.sub,
                            plaid_access_token: accessToken
                        }
                    }
                );
                if(!error && data) {
                    console.log(response.data);
                }
            }
            updateUserAccessToken();
        }
    }, [user, accessToken])
    

    return (
        <>
            <div>TestPlaid</div>
            <p>{linkToken}</p>
            <p>{accessToken}</p>
        </>
    )
}

export default TestPlaid