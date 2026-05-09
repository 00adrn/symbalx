<script lang="ts">
    import type { navButton } from '$lib/types'

    const { text, redirect }: navButton = $props();
</script>

<div class="h-12 w-12 static top-0 rounded-full bg-taupe-600 hover:bg-taupe-500 transition-all">
    <a aria-label={text} href={redirect} class="h-full w-full 
        text-taupe-200 font-semibold text-lg flex items-center justify-center pt-1 pb-1 pl-2 pr-2">
    </a>
</div>