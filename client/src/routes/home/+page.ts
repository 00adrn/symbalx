import type { PageLoad } from './$types'
import { trackItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const currentTrackResponse = await fetch("/api/user/current-track");
    let currentTrackData: trackItem | null = null;

    if(currentTrackResponse.ok) {
        currentTrackData = new trackItem(await currentTrackResponse.json());
    }

    return {
        currentTrackData: currentTrackData,
    }
}