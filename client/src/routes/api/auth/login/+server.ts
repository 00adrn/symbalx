import { PRIVATE_BACKEND_KEY, PRIVATE_BACKEND_URL } from "$env/static/private";
import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'

class loginData {
    email: string = '';
    password: string = '';
}

export const POST: RequestHandler = async ({ request, fetch }) => {
    console.log("attempting user login")

    const accountData: loginData = await request.json();
    console.log(JSON.stringify(accountData));


    const response = await fetch(`${PRIVATE_BACKEND_URL}/auth/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `${PRIVATE_BACKEND_KEY}`
        },
        body: JSON.stringify(accountData)
    })

    if (!response.ok) {
        console.log("error logging in bruhhhh" + response.status);
        return error(401, "Error creating user account");
    }

    console.log(await response.json());
    return json({ success: true })
}