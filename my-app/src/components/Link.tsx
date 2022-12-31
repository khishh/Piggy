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

type CreateAccessTokenResponse = {
    item_id: string,
    access_token: string
}

const Link = () => {

    const { linkToken, dispatch } = useContext(plaidContext);

    const onSuccess = useCallback((public_token: string, metadata: PlaidLinkOnSuccessMetadata) => {
        console.log(public_token);

        const exchangePublicTokenForAccessToken = async() => {
            const postData = {
                PublicToken: public_token
            }
            
            const response = await fetch("http://localhost:80/api/create_access_token", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(postData)
            });

            if(!response.ok) {
                dispatch({
                    type: "SET_STATE",
                    state: {
                        itemId: "",
                        accessToken: "",
                    }
                });
                return;
            }  

            const data: CreateAccessTokenResponse = await response.json();
            dispatch({
                type: "SET_STATE",
                state: {
                    itemId: data.item_id,
                    accessToken: data.access_token
                }
            });
        };

        exchangePublicTokenForAccessToken();
        
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