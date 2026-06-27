import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PRIVATE_BACKEND_URL } from '$env/static/private'

export const GET: RequestHandler = async ({ cookies, fetch }) => {
    const sessionToken = cookies.get("smblx-session");
    if (!sessionToken)
        return error(400, "No session token found");

    //console.log("Attempting to fetch user data...");

    const response = await fetch(`${PRIVATE_BACKEND_URL}/user/profile`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${sessionToken}`
        },
    });

    if (!response.ok)
        return error(response.status, "Error fetching user data.");

    const data = await response.json();

    //console.log(`Successfully read data: ${JSON.stringify(data)}`);

    return json(data)
}
