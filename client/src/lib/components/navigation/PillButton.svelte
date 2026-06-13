<script lang="ts">
    const { text, onClick }:
        { text: string, onClick?: () => void } = $props();

</script>

<button class="w-full font-bold text-taupe-200 text-xl py-2 px-4 rounded-full bg-taupe-600 hover:bg-taupe-500 transition-colors duration-200"
    onclick={onClick ||  (() => {})}>
    {text}
</button>