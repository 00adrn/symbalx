<script lang="ts">
    let { textInput = $bindable(), placeholder }: { textInput: string, placeholder: string } = $props();

</script>

<input class="w-full h-10 text-black bg-taupe-100 rounded-full text-xl px-2 py-1 "
    type="text"
    bind:value={textInput}
    {placeholder}
    />