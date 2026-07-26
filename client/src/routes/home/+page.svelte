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


<div class="w-full md:w-full min-h-screen bg-gray-200 px-4 py-4 md:px-64 flex flex-col gap-2">
    <div class="w-full flex flex-col md:flex-row items-center justify-between gap-4 md:gap-2">

        <div class="flex flex-row gap-2 items-center">
            <Avatar.Root class="w-24 h-24">
                <Avatar.Image src="" alt="profile picture" />
                <Avatar.Fallback>{profileData ? profileData.username[0] : "username"}</Avatar.Fallback>
            </Avatar.Root>

            <div class="flex flex-col gap-1 items-start justify-center">
                <p class="text-lg sm:text-xl text-purple-500">Welcome back,</p>
                <p class="text-2xl sm:text-3xl font-bold text-purple-900">{profileData ? profileData.username : "username"}</p>
            </div>
        </div>

        <div class="w-full md:w-auto">
            <CurrentTrackCard track={data.currentTrackData} />
        </div>
    </div>

    <div class="w-full flex flex-row gap-2 pt-8">
        <p class="text-xl text-purple-900 font-bold">Recent Tracks:</p>
    </div>

    <div class="w-full flex flex-row gap-2">
    
        <RecentTracks tracks={data.recentTracksData}/>
    
    </div>

    <div class="w-full flex flex-row gap-2 pt-8">
        <p class="text-xl text-purple-900 font-bold">Top Tracks, all time: </p>
    </div>

    <div class="w-full flex flex-row gap-2">
    
        <BarChart
            class=""
            data={topTrackChartData}
            orientation="horizontal"
            x="value"
            y="key"
            axis="y"
            rule={false}
            padding={{ left: 2, bottom: 0, top: 0, right: 2 }}
            height={250}
            labels={true}
            props={{
                labels: {
                    textAnchor: 'end',
                    fill: 'purple',
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
</div>