import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PUBLIC_SPOTIFY_API_BASE_URL } from '$env/static/public'

export const GET: RequestHandler = async ({ cookies, url, fetch }) => {

    const type = url.searchParams.get('type');
    const uri = url.searchParams.get('uri');
    const authToken = cookies.get('spotify_access_token');


    console.log(`Received request for ${type} with uri ${uri}`);
    console.log(authToken);

    if (!type || !uri) return error(400, "Missing type or uri query parameter");

    if (!authToken) return error(400, "Missing Spotify access token");

    const endpoint = `${PUBLIC_SPOTIFY_API_BASE_URL}/${type}/${uri}`;

    const response = await fetch(endpoint, {
        method : "GET",
        headers : {
            "Authorization" : `Bearer ${authToken}`
        }
    });

    if (!response.ok) 
        return error(response.status, `Spotify API returned ${response.status} ${response.statusText}`);
    const data = await response.json();

    return json(data);
}