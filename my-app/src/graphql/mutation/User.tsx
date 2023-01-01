import { gql } from "@apollo/client";

export const CREATE_USER_ON_SIGN_IN = gql`
    mutation CreateUserOnSignIn($input: UserInput!) {
        CreateUserOnSignIn(input: $input) {
            id,
            first_name,
            last_name,
            picture,
            email,
            sub,
            access_token
        }
    }
`

export const UPDATE_USER_WITH_ACCESS_TOKEN = gql`
    mutation UpdateUserWithAccessToken($auth0_sub_id: String!, $plaid_access_token: String!) {
        UpdateUserWithAccessToken(
            id: $auth0_sub_id,
            access_token: $plaid_access_token
        ) {
            id,
            email,
            first_name,
            last_name,
            sub,
            access_token
        }
    }
`