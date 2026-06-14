import { PRIVATE_BACKEND_KEY, PRIVATE_BACKEND_URL } from "$env/static/private";
import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'

export const POST: RequestHandler = async ({ request, fetch }) => {
    console.log("attempting user login")

    const accountData = await request.json();

    const response = await fetch(`${PRIVATE_BACKEND_URL}/auth/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `${PRIVATE_BACKEND_KEY}`
        },
        body: JSON.stringify({ accountData })
    })

    if (!response.ok)
        return error(401, "Error creating user account");

    console.log(response.json());
    return json({ success: true })
}