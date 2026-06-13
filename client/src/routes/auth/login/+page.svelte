<script lang="ts">
    import LineTextInput from '$lib/components/input/LineTextInput.svelte';
    import PillButton from '$lib/components/navigation/PillButton.svelte';

    const emailInput: string = $state('');
    const passwordInput: string = $state('');

    const onSubmit = () => {

    }
</script>

<div class="w-3/5 min-h-screen bg-taupe-700 p-4 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-row items-center justify-center gap-2">
        <form class="min-w-1/3 flex flex-col gap-2 bg-taupe-800 p-2 rounded-md items-center justify-center text-taupe-200 text-xl
                    border-md border-taupe-200">
            <p>Create an Account</p>
            <LineTextInput textInput={emailInput} placeholder="Email" />
            <LineTextInput textInput={passwordInput} placeholder="Password" />
            <div class="w-full px-2">
                <PillButton text="Submit" onClick={onSubmit} />
            </div>
        </form>
    </div>
</div>