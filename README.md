[WIP]Personal finance app using Plaid Api


# Learning goals in this project

- Get familiar with GoLang through building REST APIs
- Build GraphQL supported backend server from the scratch
- Learn and implement system design of the modern full-stack application
- Get deeper exposures to DevOps tools such as Docker, nginx, Kubernetes, Workflow, and more.
...
- More learnings as this project progresses

# [WIP] System Design Overview

![diagram (1)](https://user-images.githubusercontent.com/60908906/210921954-5feb88c1-84e0-4ff8-8db1-c06ddff148eb.png)


# 1. Clone the repository

Using https:

```
git clone https://github.com/khishh/personal-finance-app.git
```

Using ssh:
```
git@github.com:khishh/personal-finance-app.git
```


# 2. Set up environment variables

This application uses Auth0 for the user authentication and Plaid API to authenticate and communicate with financial institutions of users. 
Therefore, you need to create an account for Auth0 and Plaid API. There should not be any incurred costs for trying this app out.

First copy `.env.example` to a new file called `.env` and you will fill out the environment variables for Plaid and Auth0 below. 

```
cp .env.example .env
```

- Obtain Auth0 domain and client ID from your application dashboard and place them within your `.env` (`REACT_APP_AUTH0_DOMAIN` and `REACT_APP_AUTH0_CLIENT_ID`). More info from https://auth0.com/docs/quickstart/spa/react/interactive

- Obtain `PLAID_CLIENT_ID` and `PLAID_SECRET` from https://dashboard.plaid.com/account/keys and place them within your `.env` file after you create your own Plaid account.

# 3. Run this application locally

## Run with Docker

### Pre-requisites

- Docker installed on your machine
- You have created `.env` file and fill out variables for Plaid and Auth0

### How to run

```
docker compose up --build --detach
```

Once all containers spin ups, you can do
- open the web application at `localhost:3000`
- open the graphql playground at `localhost:8080`


## Links
ER diagram I created for this project: https://dbdiagram.io/d/63a78b3a7d39e42284e7630f
