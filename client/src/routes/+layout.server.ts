import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ cookies, fetch }) => {
    const response = await fetch("/api/user/profile-data", {
        method: "GET"
    });

    let data = null;

    if (response.ok) {
        data = await response.json();
        cookies.set("spotify_access_token", data.access_token, {
            path: "/",
            httpOnly: true,
            sameSite: "lax",
            maxAge: 60 * 60,
        });
    }

    return {
        profileData: data,
    }
}