import { gql } from "@apollo/client";


export const GET_ONE_USER = gql`
    query GetOneUser($auth0_sub_id: String!) {
        GetOneUser(id: $auth0_sub_id) {
            id,
            first_name,
            last_name,
            picture,
            email,
            sub
        }
    }
`