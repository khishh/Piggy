import { gql } from "@apollo/client";

export const CREATE_USER_ON_SIGN_IN = gql`
    mutation CreateUserOnSignIn($input: UserInput!) {
        CreateUserOnSignIn(input: $input) {
            id,
            first_name,
            last_name,
            picture,
            email,
            sub
        }
    }
`