import { redirect } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PUBLIC_REDIRECT_URL } from '$env/static/public'
import { PRIVATE_SPOTIFY_CLIENT_ID } from '$env/static/private'

export const GET: RequestHandler = async ({ cookies }) => {
    const state = createRandomString(64);
    const scope = "user-read-currently-playing playlist-read-private playlist-read-collaborative user-follow-read user-read-private";

    const params = new URLSearchParams();
    params.append("response_type", "code");
    params.append("client_id", PRIVATE_SPOTIFY_CLIENT_ID);
    params.append("scope", scope);
    params.append("redirect_uri", PUBLIC_REDIRECT_URL);
    params.append("state", state);

    cookies.set("state", state, {
        path: "/auth/spotify",
        httpOnly: true,
        sameSite: "lax",
        maxAge: 60 * 5
    });

    console.log("Initializing spotify auth...");
    redirect(302, "https://accounts.spotify.com/authorize?" + params.toString());
}

const createRandomString = (length: number) => {
    const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    const randVals = new Uint8Array(length);
    crypto.getRandomValues(randVals);

    let result = "";
    
    for (let i = length; i > 0; --i)
        result += chars[randVals[i] % chars.length];

    return result;
}
