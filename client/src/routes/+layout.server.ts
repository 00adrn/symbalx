import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ cookies, fetch }) => {
    const token = cookies.get('smblx-session');
    console.log(token ? "got token" : "no token");

    const response = await fetch("/api/user/profile-data", {
        method: "GET"
    });

    let data = null;

    if (response.ok) {
        data = await response.json();
        console.log(`Got data: ${JSON.stringify(data)}`)
    }

    return {
        profileData: data,
    }
}