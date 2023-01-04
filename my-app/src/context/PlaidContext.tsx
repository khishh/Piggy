import React, { Dispatch, ReactNode, createContext, useReducer } from "react";

interface PlaidState {
    linkToken: string,
    accessToken: string,
    itemId: string
};

type PlaidAction = {
    type: string,
    state: Partial<PlaidState>
};

const initialPlaidState: PlaidState = {
    linkToken: "",
    accessToken: "",
    itemId: ""
};

interface PlaidContext extends PlaidState {
    dispatch: Dispatch<PlaidAction>
}

const plaidContext = createContext<PlaidContext>(initialPlaidState as PlaidContext);

export const PlaidProvider: React.FC<{ children: ReactNode }> = (
    props
) => {
    const reducer = (
        state: PlaidState,
        action: PlaidAction
    ): PlaidState => {
        console.log(action);
        switch (action.type) {
            case "SET_STATE":
                return { ...state, ...action.state }
            default:
                return { ...state }
        }
    };

    const [state, dispatch] = useReducer(reducer, initialPlaidState)

    return <plaidContext.Provider value={{ ...state, dispatch }}>{props.children}</plaidContext.Provider>
}

export default plaidContext;
