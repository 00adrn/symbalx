<script lang="ts">
    import type { PageProps } from './$types';
    import { dataItem, trackItem } from '$lib/types'
    import { getUserContext } from '$lib/context';
    import CurrentTrackCard from '$lib/components/activity/CurrentTrackCard.svelte';
    import RecentTracks from '$lib/components/activity/RecentTracks.svelte';
    import { BarChart } from 'layerchart';
    import * as Avatar from "$lib/components/ui/avatar/index.js"

    const { data }: PageProps = $props();
    const profileData = getUserContext();

    const topTrackChartData: dataItem[] = data.topTracksData?.map((track: trackItem) => {
        return new dataItem(`${track.name} - ${track.getArtistString}`, track.timesListened);
    });


</script>


<div class="w-4/5 min-h-screen bg-taupe-700 p-2 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-row items-center justify-between gap-2">

        <div class="w-full flex flex-row gap-2">
            <Avatar.Root class="w-24 h-24">
                <Avatar.Image src="" alt="profile picture" />
                <Avatar.Fallback>{profileData ? profileData.username[0] : "username"}</Avatar.Fallback>
            </Avatar.Root>

            <div class="flex flex-col gap-1 items-center justify-center">
                <p class="text-xl text-taupe-500">Welcome back,</p>
                <p class="text-3xl font-bold text-taupe-200">{profileData ? profileData.username : "username"}</p>
            </div>
        </div>

        <CurrentTrackCard track={data.currentTrackData} />
    </div>

    <div class="w-full flex flex-row gap-2 pt-8">
        <p class="text-xl text-taupe-200 font-bold">All time top tracks:</p>
    </div>

    <div class="w-full flex flex-row gap-2">
    
        <BarChart
            class="bg-taupe-800 rounded-md"
            data={topTrackChartData}
            orientation="horizontal"
            x="value"
            y="key"
            axis="y"
            rule={false}
            padding={{ left: 12, bottom: 0, top: 0, right: 8 }}
            height={250}
            labels={true}
            props={{
                labels: {
                    textAnchor: 'end',
                    fill: 'white',
                },
                yAxis: {
                    tickLabelProps: {
                        textAnchor: "start",
                        dx: 10,
                        class: "text-sm text-taupe-200",
                    }
                },
                bars: {
                    stroke: 'none',
                    height: 35,
                }
            }}
        />
        
    </div>

    <div class="w-full flex flex-row gap-2 pt-8">
        <p class="text-xl text-taupe-200 font-bold">Recently Played Tracks:</p>
    </div>

    <div class="w-full flex flex-row gap-2">
    
        <div class="w-full">
            <RecentTracks tracks={data.recentTracksData}/>
        </div>
    
    </div>
</div>