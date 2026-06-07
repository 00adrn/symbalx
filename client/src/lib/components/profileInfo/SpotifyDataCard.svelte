<script lang="ts">
    import { goto } from '$app/navigation'
    import type { spotifyUserItem } from '$lib/types';
    import PillButton from '$lib/components/navigation/PillButton.svelte';

    const { spotifyProfileData }: { spotifyProfileData: spotifyUserItem | null } = $props();
</script>

<div class="w-full min-h-full bg-transparent flex flex-row items-center gap-8 rounded-md">
    <div class="flex flex-col gap-4">
        <p class="font-bold text-md text-taupe-500">{spotifyProfileData ?"Connected Spotify account:" : "Connect your Spotify account"}</p>
        
        <div class="flex flex-row gap-4 items-center">
            {#if spotifyProfileData}
                <img class="h-12 w-12 rounded-md" src={spotifyProfileData.getImage} alt="Spotify PFP"/>
                <p class="font-semibold text-xl text-taupe-200">{spotifyProfileData.display_name}</p>
            {:else}
                <PillButton text="Connect Spotify" onClick={() => goto("/auth/spotify/login")} />
            {/if}
        </div>
    </div>
</div>