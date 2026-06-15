export const handleSignUp = async (username: string, email: string, password: string) => {
    const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, email, password })
    });

    return response.ok ? true : false;
}

export const handleLogin = async (email: string, password: string) => {

    const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    });

    return response.ok ? true : false;
}