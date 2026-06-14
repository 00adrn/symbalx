import type { PageLoad } from './$types'
import { spotifyUserItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const spotifyProfileResponse = await fetch("/api/user");
    
    return {
        spotifyProfileData: spotifyProfileResponse.ok ? new spotifyUserItem(await spotifyProfileResponse.json()) : null
    
    }
}