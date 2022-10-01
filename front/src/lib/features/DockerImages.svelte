<script>
  import FaBox from "svelte-icons/fa/FaBox.svelte";
  import FaParachuteBox from "svelte-icons/fa/FaParachuteBox.svelte";
  import Breadcrumb from "../bars/Breadcrumb.svelte";
  import FaHistory from "svelte-icons/fa/FaHistory.svelte";
  import FaDropbox from "svelte-icons/fa/FaDropbox.svelte";
  import {
    localDockerImages,
    localDockerImagesIsLoading,
  } from "../../stores/docker-images";
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import { scan } from "../../stores/scan";
  import { useNavigate } from "svelte-navigator";
  import { history } from "../../stores/history";

  let breadcrumb = ["Docker Images"];
  let remoteSearch = "";
  const navigate = useNavigate();
  let filterHistory = [];

  $: if ($history) {
    filterHistory = $history.filter((h) => h.type === "docker-images");
  }

  onMount(() => {
    localDockerImages.get();
  });

  function handleScan(mode, name, save = true) {
    if (name !== "") {
      scan.getScanImage(mode, name);
      navigate("/scan-report");
      if (save) {
        history.saveScan(name, "docker-images");
      }
    }
  }
</script>

<div class="container">
  <Breadcrumb bind:paths={breadcrumb} />
  <div class="menu-container">
    <div class="menu-items">
      <div class="local card light">
        <div class="title">
          <div class="icon yellow">
            <FaBox />
          </div>
          <h2>Local Docker Images - SCA (Software Compisotion Analysis)</h2>
        </div>
        <hr />
        <div class="table header">
          <span>Id</span>
          <span>Name</span>
          <span>Tag</span>
          <span>Date</span>
          <span>Size</span>
        </div>
        {#if $localDockerImages.length}
          <div in:fade class="table body">
            {#each $localDockerImages as image}
              <span
                on:click={() =>
                  handleScan("local", `${image.Name}:${image.Tag}`)}
                class="selectable yellow"
                style="font-weight: bold;">{image.Id}</span
              >
              <span>{image.Name}</span>
              <span>{image.Tag}</span>
              <span>{image.Date}</span>
              <span>{image.Size}</span>
            {/each}
          </div>
        {:else if $localDockerImagesIsLoading}
          <div class="no-data">
            <div class="lds-ripple">
              <div />
              <div />
            </div>
          </div>
        {:else}
          <div class="no-data">
            <div class="icon">
              <FaDropbox />
            </div>
            <h3>No scan have been doned</h3>
          </div>
        {/if}
      </div>
      <div class="remote card light">
        <div class="title">
          <div class="icon yellow">
            <FaParachuteBox />
          </div>
          <h2>Remote Docker Images</h2>
        </div>
        <hr />
        <div class="remote-search">
          <form
            on:submit|preventDefault={() => handleScan("remote", remoteSearch)}
          >
            <input
              type="text"
              bind:value={remoteSearch}
              placeholder="docker image"
            />
            <button disabled={remoteSearch === ""}>Search Image</button>
          </form>
        </div>
      </div>
    </div>
    <div class="history card light">
      <div class="title">
        <div class="icon yellow">
          <FaHistory />
        </div>
        <h2>Reports History</h2>
      </div>
      <hr />
      <div class="table header">
        <span>Name</span>
        <span>Date</span>
      </div>
      {#if filterHistory.length > 0}
        <div in:fade class="table body">
          {#each filterHistory as hist}
            <span
              class="yellow selectable"
              on:click={() => handleScan("local", hist.name, false)}
            >
              {hist.name}
            </span>
            <span>{hist.date}</span>
          {/each}
        </div>
      {:else}
        <div class="no-data">
          <div class="icon">
            <FaHistory />
          </div>
          <h3>No scan have been doned</h3>
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .container {
    width: 100%;
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
    grid-template-rows: calc(100% - 220px) 220px;
  }

  .local {
    margin-bottom: var(--margin);
  }

  .local .table {
    grid-template-columns: repeat(5, 1fr);
  }

  .local .table.body {
    max-height: calc(100% - 120px);
  }

  .local .no-data {
    height: calc(100% - 100px);
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

  .remote-search {
    margin: var(--margin);
    margin-left: 0;
    margin-right: 0;
  }

  .remote-search input {
    width: 100%;
    margin-bottom: var(--margin);
  }

  .remote-search button {
    width: 100%;
  }
</style>
