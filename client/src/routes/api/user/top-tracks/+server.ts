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
    
    const tracks: trackItem[] = [];

    for (const track of data.tracks) {
        const cleanUri = track.uri.split(':').pop();
        console.log(`${cleanUri} | Count: ${track.count}`);

        const trackResponse = await fetch(`/api/spotify/fetch?type=tracks&uri=${cleanUri}`);
        if (!trackResponse.ok) {
            console.error(`Error fetching track details: ${trackResponse.status} ${trackResponse.statusText}`);
            continue;
        }

        const tdata = await trackResponse.json();
        const newTrack = new trackItem(tdata, track.count);

        tracks.push(newTrack);
    }

    return json(tracks);
}
