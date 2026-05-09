import type { PageLoad } from './$types'
import type { trackItem } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
    const testFetchResponse = await fetch("/api/fetch?type=tracks&uri=3n3Ppam7vgaVa1iaRUc9Lp");
    const testFetchData = await testFetchResponse.json();

    return {
        track: testFetchData,
    }
}

const testTrackFetch = async () => {
    const response = await fetch("/api/fetch?type=tracks&uri=3n3Ppam7vgaVa1iaRUc9Lp");
    const data = await response.json();
    console.log(data);
}