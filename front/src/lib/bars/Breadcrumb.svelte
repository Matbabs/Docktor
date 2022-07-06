<script>
    export let paths = []
</script>

<div class="breadcrumb">
    {#each paths as path, i}
        <h4>{path}</h4>
        {#if i !== paths.length - 1}
           <h4> > </h4>
        {/if}
    {/each}
</div>

<style>
    .breadcrumb {
        background: var(--light-dark);
        padding: var(--padding);
        margin: 0;
        display: flex;
    }

    .breadcrumb h4 {
        margin-right: var(--margin);
    }
</style>