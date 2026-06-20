import { PRIVATE_BACKEND_KEY, PRIVATE_BACKEND_URL } from "$env/static/private";
import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'


export const POST: RequestHandler = async ({ request, cookies, fetch }) => {
    console.log("attempting user login")

    const accountData = await request.json();

    const response = await fetch(`${PRIVATE_BACKEND_URL}/auth/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${PRIVATE_BACKEND_KEY}`
        },
        body: JSON.stringify(accountData)
    })

    if (!response.ok)
        return error(401, "Error logging in");

    const data = await response.json();
    cookies.set("smblx-session", data.token, {
        path: "/",
        httpOnly: true,
        sameSite: "strict",
    });

    return json("Success", { status: 200 });
}