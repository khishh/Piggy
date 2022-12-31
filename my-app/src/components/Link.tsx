import React, { useCallback, useContext, useEffect } from 'react'

import {
    usePlaidLink,
    PlaidLinkError,
    PlaidLinkOptions,
    PlaidLinkOnSuccess,
    PlaidLinkOnSuccessMetadata,
    PlaidLinkOnExitMetadata,
    PlaidLinkOnEventMetadata
} from 'react-plaid-link';
import plaidContext from '../context/PlaidContext';
import { response } from 'express';
// import { Button } from '@mui/material';

const Link = () => {

    const { linkToken, dispatch } = useContext(plaidContext);

    // console.log("Link Component render " + linkToken);
    

    const onSuccess = useCallback((public_token: string, metadata: PlaidLinkOnSuccessMetadata) => {
        console.log(public_token);

        const postData = {
            PublicToken: public_token
        }
        
        fetch("http://localhost:80/api/create_access_token", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(postData)
        }).then(response => {
            if(!response.ok) {
                return;
            }
            return response.json();
        }).then(json => {
            console.log(json);
        });
        
    }, [dispatch]);

    const onExit = useCallback((error: PlaidLinkError | null, metadata: PlaidLinkOnExitMetadata) => {

    }, [dispatch]);

    const onEvent = useCallback((eventName: string, metadata: PlaidLinkOnEventMetadata) => {

    }, [dispatch]);

    const config: PlaidLinkOptions = {
        onSuccess: onSuccess,
        onExit: onExit,
        onEvent: onEvent,
        token: linkToken
    };

    const { open, exit, ready } = usePlaidLink(config);

    return (
        <>
            <div>Link</div>
            <button  onClick={() => open()} disabled={!ready}>
                Launch Link
            </button>
        </>
    )
}

export default Link