import type { PageLoad } from './$types'
import { spotifyUserItem } from '$lib/types';
import { getUserContext } from '$lib/context';

export const load: PageLoad = async ({ fetch }) => {
    const spotifyProfileResponse = await fetch("/api/spotify/user");
    
    return {
        spotifyProfileData: spotifyProfileResponse.ok ? new spotifyUserItem(await spotifyProfileResponse.json()) : null,

    }
}