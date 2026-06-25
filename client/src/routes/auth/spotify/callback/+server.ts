import { PRIVATE_SPOTIFY_CLIENT_ID, PRIVATE_SPOTIFY_CLIENT_SECRET, PRIVATE_BACKEND_URL} from "$env/static/private"
import { PUBLIC_REDIRECT_URL } from "$env/static/public"
import { redirect } from "@sveltejs/kit"
import type { RequestHandler } from "./$types"

export const GET: RequestHandler = async ({ url, cookies, fetch }) => {
    const state = cookies.get("state");
    const sessionToken = cookies.get("smblx-session");
    const code = url.searchParams.get("code");
    const returnedState = url.searchParams.get("state");

    if (!state || !code || !returnedState || state !== returnedState) {
        console.log("Error on callback")
        redirect(302, "/auth/spotify/login"); 
    }

    const params = new URLSearchParams();
    params.append("grant_type", "authorization_code");
    params.append("code", code);
    params.append("redirect_uri", PUBLIC_REDIRECT_URL);

    const response = await fetch("https://accounts.spotify.com/api/token", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "Authorization": "Basic " + btoa(PRIVATE_SPOTIFY_CLIENT_ID + ":" + PRIVATE_SPOTIFY_CLIENT_SECRET)
        },
        body: params
    });

    if (!response.ok) 
        redirect(302, "/spotify/login");

    const data = await response.json();
    if (data.error || !data.access_token) { redirect(302, "/spotify/login"); }
    
    const backendResponse = await fetch(`${PRIVATE_BACKEND_URL}/user/spotify/update`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${sessionToken}`
        },
        body: JSON.stringify({
            access_token: data.access_token,
            refresh_token: data.refresh_token,
        }),
    });

    if (!backendResponse.ok) {
        console.error("Error updating backend spotify values.");
        const errorText = await backendResponse.text();
        console.error(`Response: ${errorText}`);
    }

    cookies.set("spotify_access_token", data.access_token, {
        path: "/",
        httpOnly: true,
        sameSite: "lax",
        maxAge: 60 * 60,
    });

    console.log(`Spotify auth succeeded with token ${data.access_token}!`);
    redirect(302, "/home");
}