import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PRIVATE_BACKEND_URL } from '$env/static/private'
import { trackItem } from '$lib/types';

export const GET: RequestHandler = async ({ cookies, fetch }) => {
    const sessionToken = cookies.get("smblx-session");
    if (!sessionToken)
        return error(400, "No session token found");

    const response = await fetch(`${PRIVATE_BACKEND_URL}/user/top-tracks`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${sessionToken}`
        },
    });

    if (!response.ok) {
        console.error(`Error fetching recent tracks: ${response.status} - ${response.statusText}`);
        return error(response.status, "Error fetching recent tracks.");
    }

    const data = await response.json();
    
    let tracks: trackItem[] = [];
    data.tracks.map(async (track: any, index: number) => {
        const cleanUri = track.uri.split(':').pop();
        console.log(`Track #${index}: ${cleanUri} | Count: ${track.count}`)

    });

    return json(tracks)
}
