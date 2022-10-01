<script>
  import Breadcrumb from "../bars/Breadcrumb.svelte";
  import FaHistory from "svelte-icons/fa/FaHistory.svelte";
  import { fade } from "svelte/transition";
  import { scan } from "../../stores/scan";
  import { useNavigate } from "svelte-navigator";
  import { history } from "../../stores/history";
  import FaGitSquare from "svelte-icons/fa/FaGitSquare.svelte";
  import SvelteMarkdown from "svelte-markdown";
  import { sessionRepo } from "../../stores/path";
  import { onMount } from "svelte";
  import DiGitBranch from "svelte-icons/di/DiGitBranch.svelte";

  const REPLACE_URI = "##REPLACE##";
  const GIT_HUB_SPLITTER = "https://github.com/";
  const GIT_LAB_SPLITTER = "https://gitlab.com/";
  const GIT_HUB =
    "https://raw.githubusercontent.com/##REPLACE##/master/README.md";
  const GIT_LAB = "https://gitlab.com/sulinos/##REPLACE##/-/raw/master/README";

  let breadcrumb = ["Repo GIT"];
  let remoteSearch = "";
  let markdown = "";
  let markdownIsLoading = false;
  const navigate = useNavigate();
  let filterHistory = [];

  $: if ($history) {
    filterHistory = $history.filter((h) => h.type === "repo-git");
  }

  onMount(() => {
    remoteSearch = $sessionRepo;
    if (remoteSearch && remoteSearch !== "") {
      checkRepo();
    }
  });

  function handleScan(name, save = true) {
    if (name !== "") {
      sessionRepo.savePath(name);
      scan.getScanRemoteRepo(name);
      navigate("/scan-report");
      if (save) {
        history.saveScan(name, "repo-git");
      }
    }
  }

  function checkRepo() {
    markdownIsLoading = true;
    markdown = "";
    if (remoteSearch !== "") {
      if (remoteSearch.includes(GIT_HUB_SPLITTER)) {
        findReadMe(remoteSearch, GIT_HUB_SPLITTER, GIT_HUB);
      } else if (remoteSearch.includes(GIT_LAB_SPLITTER)) {
        findReadMe(remoteSearch, GIT_LAB_SPLITTER, GIT_LAB);
      }
    }
  }

  function findReadMe(uri, splitter, uri_provider) {
    const splits = uri.split(splitter);
    const repoName = splits[splits.length - 1];
    const readmeUri = uri_provider.replace(REPLACE_URI, repoName);
    fetch(readmeUri)
      .then((res) => res.text())
      .then((res) => {
        if (res && res !== "") {
          markdown = res;
        }
      })
      .finally(() => (markdownIsLoading = false));
  }
</script>

<div class="container">
  <Breadcrumb bind:paths={breadcrumb} />
  <div class="menu-container">
    <div class="menu-items">
      <div class="remote card light">
        <div class="title">
          <div class="icon purple">
            <FaGitSquare />
          </div>
          <h2>Remote Repo GIT - SCA (Software Compisotion Analysis)</h2>
        </div>
        <hr />
        <form on:submit|preventDefault={checkRepo}>
          <input
            type="text"
            bind:value={remoteSearch}
            placeholder="https://your/repo/uri"
          />
          <button disabled={remoteSearch === ""}>Check Repo</button>
          <button
            disabled={remoteSearch === ""}
            on:click={() => handleScan(remoteSearch)}>Scan Repo</button
          >
        </form>
        <div class="dark card markdown-container">
          {#if markdown !== "" && markdown !== "404: Not Found" && !markdownIsLoading}
            <SvelteMarkdown source={markdown} />
          {:else if markdownIsLoading}
            <div class="no-data">
              <div class="lds-ripple">
                <div />
                <div />
              </div>
            </div>
          {:else}
            <div class="no-data">
              <div class="icon">
                <DiGitBranch />
              </div>
              <h3>No preview for your URI</h3>
              <p class="red">
                Preview feature works only for public github project
              </p>
              <h4 style="font-weight: bold;" class="purple">
                Scan works for all accessible repo
              </h4>
            </div>
          {/if}
        </div>
      </div>
    </div>
    <div class="history card light">
      <div class="title">
        <div class="icon purple">
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
              class="purple selectable"
              on:click={() => handleScan(hist.name, false)}
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
    grid-template-rows: 1fr;
  }

  .remote {
    max-width: 1123px;
    height: calc(100vh - 175px);
  }

  .remote form input,
  button {
    width: 100%;
    margin-bottom: var(--margin);
  }

  .remote .markdown-container {
    height: calc(100% - 235px);
    overflow: auto;
  }

  .remote .markdown-container .no-data {
    height: calc(100% - 20px);
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
