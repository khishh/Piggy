import { useQuery } from '@apollo/client';
import { useAuth0 } from '@auth0/auth0-react'
import React from 'react'
import { Book } from '../graphql/model/Book';
import { GET_ALL_BOOKS } from '../graphql/query/Books';

const TestProfile = () => {

    const { user, isAuthenticated, isLoading } = useAuth0();

    const { loading, error, data } = useQuery<{
        GetAllBooks: Book[]
    }>(GET_ALL_BOOKS);

    if (isLoading) {
        return <div>Loading...</div>
    } else if (!isAuthenticated) {
        return <div>Not AUthenticated yet</div>
    } else if (!user) {
        return <div>No user</div>
    }

    console.log(user)
    console.log(data)

    return (
        <>
            <h1>This screen should be only visible only if user is signed in.</h1>

            {user && (
                <div>
                    <img src={user.picture} alt={user.name} />
                    <h2>{user.name}</h2>
                    <p>{user.email}</p>
                </div>
            )}

            {data && (
                <div>
                    {data!.GetAllBooks.map((book) =>
                        <div key={book.id.toString()}>
                            <p>{book.id.toString()}</p>
                            <p>{book.title}</p>
                            <p>{book.author}</p>
                            <p>{book.publisher}</p>
                        </div>
                    )}
                </div>
            )}
        </>

    );
}

export default TestProfile