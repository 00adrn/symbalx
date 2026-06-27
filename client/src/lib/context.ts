import { createContext } from 'svelte'

interface UserContext {
    username: string;
    email: string;
    access_token: string;
}

export const [getUserContext, setUserContext] = createContext<UserContext | null>();