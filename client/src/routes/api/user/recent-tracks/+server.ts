import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from  '@sveltejs/kit'
import { PRIVATE_BACKEND_URL } from '$env/static/private'
import { trackItem } from '$lib/types';

export const GET: RequestHandler = async ({ cookies, fetch }) => {
    const sessionToken = cookies.get("smblx-session");
    if (!sessionToken)
        return error(400, "No session token found");

    const response = await fetch(`${PRIVATE_BACKEND_URL}/user/recent-tracks`, {
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
    const trackUris: string[] = data.tracks.map((track: any) => track.uri);

    let tracks: trackItem[] = [];

    for (const uri of trackUris) {
        const cleanUri = uri.split(':').pop();

        const trackResponse = await fetch(`/api/spotify/fetch?type=tracks&uri=${cleanUri}`, {});

        const tdata = await trackResponse.json();
        console.log(`Fetched track data for URI ${uri}:`, tdata);
        tracks.push(new trackItem(tdata));
    }

    return json(tracks)
}
