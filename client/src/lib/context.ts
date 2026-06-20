import { createContext } from 'svelte'

interface User {
    username: string;
    email: string;
}

export const [getUserContext, setUserContext] = createContext<User | null>();