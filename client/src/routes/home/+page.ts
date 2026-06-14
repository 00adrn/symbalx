import type { PageLoad } from './$types'
import { trackItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const currentTrackResponse = await fetch("/api/user/current-track");


    return {
        currentTrackData: currentTrackResponse.ok ? new trackItem(await currentTrackResponse.json()) : null
    ,
    }
}