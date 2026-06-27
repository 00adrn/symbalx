import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PUBLIC_SPOTIFY_API_BASE_URL } from '$env/static/public'

export const GET: RequestHandler = async ({ cookies, fetch }) => {
    const authToken = cookies.get('spotify_access_token');
    if (!authToken)
        return error(401, "Not connected to Spotify");
    
    const endpoint = `${PUBLIC_SPOTIFY_API_BASE_URL}/me`;

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