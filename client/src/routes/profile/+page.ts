import type { PageLoad } from './$types'
import { spotifyUserItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const spotifyProfileResponse = await fetch("/api/user");
    let spotifyProfileData: spotifyUserItem | null = null; 
    
    if (spotifyProfileResponse.ok) {
        spotifyProfileData = new spotifyUserItem(await spotifyProfileResponse.json());
    }

    return {
        spotifyProfileData: spotifyProfileData,
    }
}