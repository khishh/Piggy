
export type User = {
    id: Number,
    first_name: string,
    last_name: string,
    picture: string,
    email: string,
    sub: string,
    access_token: string
}

export type UserInput = {
    first_name: string,
    last_name: string,
    picture: string,
    email: string,
    sub: string
}