import type { PageLoad } from './$types'
import { trackItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const currentTrackResponse = await fetch("/api/spotify/user/current-track");
    const recentTracksResponse = await fetch("/api/user/recent-tracks");


    return {
        currentTrackData: currentTrackResponse.ok ? new trackItem(await currentTrackResponse.json()) : null,
        recentTracksData: recentTracksResponse.ok ? (await recentTracksResponse.json()).map((item: any) => new trackItem(item)) : null
    }
}