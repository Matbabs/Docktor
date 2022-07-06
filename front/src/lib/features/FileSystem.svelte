<script>
    import Breadcrumb from "../bars/Breadcrumb.svelte"
    import FaHistory from 'svelte-icons/fa/FaHistory.svelte'
    import FaFileContract from 'svelte-icons/fa/FaFileContract.svelte'
	import { useNavigate } from "svelte-navigator"
    import {history} from "../../stores/history"
    import {fileSystem} from "../../stores/file-system"
    import { onMount } from "svelte";
    import { scan } from "../../stores/scan"
    import Doughnut from "svelte-chartjs/src/Doughnut.svelte"
    import {fade} from "svelte/transition"
    import FaChartPie from 'svelte-icons/fa/FaChartPie.svelte'
    import {sessionPath} from "../../stores/path"

    let breadcrumb = ["File System"]
    const navigate = useNavigate()
    let localPath = ""
    let filterHistory = []

    $: if($history) {
        filterHistory = $history.filter(h => h.type === "file-system")
    }

    let data = {
        labels: [],
        datasets: [{
            data: [],
            backgroundColor: ['#f69666', '#fdd765', '#f15f86', '#aa9cf2', '#78dbe7', '#a8db75'],
            borderWidth: 1
        }]
    };

    $: if($fileSystem) {
        const repartition = new Map()
        data.labels = []
        data.datasets[0].data = []
        if($fileSystem.stdout){
            $fileSystem.stdout.forEach(line => {
                if(line !== "") {
                    const elem = line.split(".")
                    const type = line.charAt(0) === "." ? ".hide" : (line.includes(".") ? `.${elem[elem.length - 1]}` : "/dir")
                    repartition.set(type, repartition.has(type) ? repartition.get(type) + 1 : 1)
                }
            });
        }
        repartition.forEach((v, k) => {
            data.labels.push(k)
            data.datasets[0].data.push(v)
        });
        data = {...data}
    }

    onMount(() => {
        localPath = $sessionPath
        fileSystem.check(localPath)
    })

    function addPath(path) {
        localPath += `${localPath.charAt(localPath.length - 1) != "/" ? "/" : ""}${path}`
        handleLocalPath()
    }

    function revertPath() {
        const paths = localPath.split("/")
        localPath = paths.slice(0, paths.length - 1).join("/")
        handleLocalPath()
    }

    function handleLocalPath() {
        sessionPath.savePath(localPath)
        fileSystem.check(localPath)
    }

    function handleScan(path, save = true){
        sessionPath.savePath(path)
        scan.getScanPath(path)
        navigate('/scan-report')
        if(save) {
            history.saveScan(path, "file-system")
        }
    }
</script>   

<div class="container">
    <Breadcrumb bind:paths={breadcrumb}/>
    <div class="menu-container">
        <div class="menu-items">
            <div class="local card light">
                <div class="title">
                    <div class="icon orange">
                        <FaFileContract/>
                    </div>
                    <h2>Local File System - SCA (Software Compisotion Analysis)</h2>
                </div>
                <hr>
                <div class="scan-tool">
                    <div>
                        <div class="form">
                            <form on:submit|preventDefault={handleLocalPath}>
                                <input type="text" bind:value={localPath} placeholder="/my/local/path">
                                <button>Check Path</button>
                                <button 
                                    disabled={!($fileSystem?.stdout.length > 1)}
                                    on:click={() => handleScan(localPath)}>
                                        Scan Path
                                </button>
                            </form>
                        </div>
                        {#if $fileSystem?.stdout.length > 1}
                            <div class="graph">
                                <Doughnut data={data} options={{}}/>
                                <h3>Directory type files repartition</h3>
                            </div>
                        {:else}
                            <div class="no-data">
                                <div class="icon">
                                    <FaChartPie/>
                                </div>
                                <h3>No check have been doned</h3>
                            </div>
                        {/if}
                    </div>
                    <div class="terminal dark card">
                        {#if $fileSystem}
                            <span class="orange head">$ ls {localPath}</span>
                            {#each $fileSystem.stdout as line}
                                {#if !line.includes(".") && line != "" && line != ".." && line != "."}
                                    <span in:fade class="clickable" on:click={() => addPath(line)}>/{line}</span>
                                {:else if line != ".." && line != "."}
                                    <span in:fade >{line}</span>
                                {/if}
                            {/each}
                            {#if localPath !== ""}
                                <span class="clickable foot" on:click={revertPath}>/..</span>
                            {/if}
                        {/if}
                    </div>
                </div>
            </div>
        </div>
        <div class="history card light">
            <div class="title">
                <div class="icon orange">
                    <FaHistory/>
                </div>
                <h2>Reports History</h2>
            </div>
            <hr>
            <div class="table header">
                <span>Name</span>
                <span>Date</span>
            </div>
            {#if filterHistory.length > 0}
                <div in:fade class="table body">
                    {#each filterHistory as hist}
                        <span class="orange selectable"
                            on:click={() => handleScan(hist.name, false)} >
                            {hist.name}
                        </span>
                        <span>{hist.date}</span>
                    {/each}
                </div>
            {:else}
                <div class="no-data">
                    <div class="icon">
                        <FaHistory/>
                    </div>
                    <h3>No scan have been doned</h3>
                </div>
            {/if}
        </div>
    </div>
</div>

<style>
    .container {
        width: 100%
    }

    .menu-container {
        padding: var(--padding);
        display: grid;
        grid-template-columns: auto 500px;
        gap: var(--margin);
    }

    .menu-items {
        height: calc(100vh - 175px);
        display: grid;
        grid-template-rows: 1fr 1fr;
    }

    .local {
        height: calc(100vh - 175px);
    }

    .scan-tool {
        width: 100%;
        height: calc(100% - 55px);
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: var(--margin);
    }

    .scan-tool input, button{
        width: 100%;
    }

    .scan-tool .form {
        display: flex;
        flex-direction: column;
        margin-bottom: 0;
    }

    .scan-tool .form input, button {
        margin-bottom: var(--margin);
    }

    .scan-tool .graph{
        width: 70%;
        margin: auto;
        text-align: center;
    }

    .scan-tool .graph h3 {
        margin-top: var(--margin);
    }

    .scan-tool .terminal {
        display: flex;
        flex-direction: column;
        overflow-y: auto;
    }

    .scan-tool .terminal .head {
        font-weight: bold;
        margin-bottom: var(--margin);
    }

    .scan-tool .terminal .foot {
        margin-top: var(--margin);
    }

    .scan-tool .clickable{
        color: var(--cyan);
    }

    .scan-tool .clickable:hover {
        cursor: pointer;
        text-decoration: underline;
    }

    .history {
        height: calc(100vh - 175px);
    }

    .history .no-data {
        height: calc(100vh - 370px);
    }

    .history .table {
        grid-template-columns: repeat(2, 1fr);
    }

    .history .table.body {
        max-height: calc(100% - 115px);
    }
</style>