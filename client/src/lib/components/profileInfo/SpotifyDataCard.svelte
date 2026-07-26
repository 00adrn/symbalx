<script lang="ts">
    import { goto } from '$app/navigation'
    import type { spotifyUserItem } from '$lib/types';
    import { Button } from '$lib/components/ui/button/index.js'
    import * as Item from '$lib/components/ui/item/index.js'


    const { spotifyProfileData }: { spotifyProfileData: spotifyUserItem | null } = $props();
</script>

<div class="w-full min-h-full bg-transparent flex flex-row items-center gap-8 rounded-md">
    <div class="flex flex-col gap-4">
        <p class="font-bold text-md text-taupe-500">{spotifyProfileData ?"Connected Spotify account:" : "Connect your Spotify account"}</p>
        
        <div class="flex flex-row gap-4">
            {#if spotifyProfileData}
                <Item.Root>
                    <Item.Media variant="image" class="h-16 w-16">
                        <img alt="Current Track" src={spotifyProfileData.images[0].url} />
                    </Item.Media>
                    <Item.Content>
                        <Item.Title class="text-lg text-purple-950">Connected Account:</Item.Title>
                        <Item.Description>   
                            {spotifyProfileData.display_name}
                        </Item.Description>
                    </Item.Content>
                </Item.Root>
            {:else}
                <Item.Root>
                    <Item.Content>
                        <Item.Title class="text-lg text-purple-950">No Spotify account connected</Item.Title>
                        <Item.Description>   
                            Connect now to start tracking!
                        </Item.Description>
                    </Item.Content>
                    <Item.Actions>
                        <Button variant="outline" class="bg-purple-950 ml-10 hover:bg-purple-800 transition-all text-purple-200 rounded-full">Connect</Button>
                    </Item.Actions>
                </Item.Root>
            {/if}
        </div>
    </div>
</div>