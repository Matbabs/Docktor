<script>
  import Breadcrumb from "../bars/Breadcrumb.svelte";
  import FaBong from "svelte-icons/fa/FaBong.svelte";
  import FaSkullCrossbones from "svelte-icons/fa/FaSkullCrossbones.svelte";
  import FaMicroscope from "svelte-icons/fa/FaMicroscope.svelte";
  import { scan, scanIsLoading } from "../../stores/scan";
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import FaThumbsUp from "svelte-icons/fa/FaThumbsUp.svelte";

  let breadcrumb = ["Scan Report"];
  let totalVulnStats;
  let totalMisconfig;
  let vuln;
  let stats = [
    { name: "UNKNOWN", value: 0, color: "purple" },
    { name: "LOW", value: 0, color: "green" },
    { name: "MEDIUM", value: 0, color: "yellow" },
    { name: "HIGH", value: 0, color: "orange" },
    { name: "CRITICAL", value: 0, color: "red" },
  ];

  $: if ($scan) {
    totalMisconfig = 0;
    totalVulnStats = 0;
    stats.forEach((s) => (s.value = 0));
    if ($scan.Results) {
      const vulnStats = (v, isVuln = true) => {
        if (!vuln) {
          vuln = v;
        }
        stats.find((s) => s.name === v.Severity).value++;
        if (isVuln) {
          totalVulnStats++;
        } else {
          totalMisconfig++;
        }
      };
      $scan.Results.forEach((result) => {
        if (result.Vulnerabilities) {
          result.Vulnerabilities.forEach((v) => vulnStats(v));
        }
        if (result.Misconfigurations) {
          result.Misconfigurations.forEach((v) => vulnStats(v, false));
        }
      });
    }
    stats = [...stats];
  }

  onMount(() => {
    vuln = undefined;
  });

  function handleChangeVuln(_vuln) {
    vuln = _vuln;
  }
</script>

<div class="container">
  <Breadcrumb bind:paths={breadcrumb} />
  <div class="menu-container">
    <div class="menu-items">
      <div class="metadata card light">
        <div class="title">
          <div class="icon red">
            <FaBong />
          </div>
          <h2>Artifact Metadata</h2>
        </div>
        <hr />
        {#if $scan}
          <div class="data" in:fade>
            <span class="data-value">{$scan.ArtifactName}</span>
            <span class="data-value">{$scan.ArtifactType}</span>
          </div>
        {/if}
      </div>
      <div class="results card light">
        <div class="title">
          <div class="icon red">
            <FaMicroscope />
          </div>
          {#if $scanIsLoading}
            <h2>Scan Results - <span class="red">Scanning ...</span></h2>
          {:else if totalMisconfig === 0}
            <h2>
              Scan Results - <span
                class={`${totalVulnStats === 0 ? "green" : "red"}`}
                >{totalVulnStats} Vulnerabilities</span
              >
            </h2>
          {:else}
            <h2>
              Scan Results - <span
                class={`${totalMisconfig === 0 ? "green" : "red"}`}
                >{totalMisconfig} Misconfigurations</span
              >
            </h2>
          {/if}
        </div>
        <hr />
        {#if totalMisconfig === 0}
          <div class="table header">
            <span>Id</span>
            <span>Lib / Pkg</span>
            <span>Version</span>
            <span>Severity</span>
          </div>
        {:else}
          <div class="table header">
            <span>Id</span>
            <span>Title</span>
            <span>Status</span>
            <span>Severity</span>
          </div>
        {/if}
        {#if $scan && (totalVulnStats > 0 || totalMisconfig > 0) && !$scanIsLoading && $scan.Results}
          <div in:fade class="table body">
            {#each $scan.Results as result}
              {#if result.Vulnerabilities}
                {#each result.Vulnerabilities as vuln}
                  <span
                    title={vuln.VulnerabilityID}
                    on:click={() => handleChangeVuln(vuln)}
                    class="selectable red"
                  >
                    {vuln.VulnerabilityID}
                  </span>
                  <span class="ellipsis" title={vuln.PkgName}
                    >{vuln.PkgName}</span
                  >
                  <span class="ellipsis" title={vuln.InstalledVersion}
                    >{vuln.InstalledVersion}</span
                  >
                  <span
                    style="font-weight: bold;"
                    class={`${
                      stats.find((s) => s.name === vuln.Severity).color
                    }`}
                    title={vuln.Severity}
                  >
                    {vuln.Severity}
                  </span>
                {/each}
              {/if}
              {#if result.Misconfigurations}
                {#each result.Misconfigurations as misconfig}
                  <span
                    title={misconfig.ID}
                    on:click={() => handleChangeVuln(misconfig)}
                    class="selectable red"
                  >
                    {misconfig.ID}
                  </span>
                  <span class="ellipsis" title={misconfig.Title}
                    >{misconfig.Title}</span
                  >
                  <span class="ellipsis" title={misconfig.Status}
                    >{misconfig.Status}</span
                  >
                  <span
                    style="font-weight: bold;"
                    class={`${
                      stats.find((s) => s.name === misconfig.Severity).color
                    }`}
                    title={misconfig.Severity}
                  >
                    {misconfig.Severity}
                  </span>
                {/each}
              {/if}
            {/each}
          </div>
        {:else if $scanIsLoading}
          <div class="no-data">
            <div class="lds-ripple">
              <div />
              <div />
            </div>
          </div>
        {:else if totalVulnStats === 0 && totalMisconfig === 0}
          <div class="no-data">
            <div class="icon green">
              <FaThumbsUp />
            </div>
            <h3 class="green">
              No {totalMisconfig === 0 ? "vulnerability" : "misconfiguration"} have
              been founded
            </h3>
          </div>
        {:else}
          <div class="no-data">
            <div class="icon">
              <FaMicroscope />
            </div>
            <h3>
              No {totalMisconfig === 0 ? "vulnerability" : "misconfiguration"} have
              been founded
            </h3>
          </div>
        {/if}
      </div>
    </div>
    <div class="report">
      <div class="stats">
        {#each stats as stat}
          <div class={`stat card ${stat.color}-bg-light`}>
            <div
              class={`status ${stat.color}-bg`}
              style={`height: ${
                (stat.value * 100) /
                (totalMisconfig === 0 ? totalVulnStats : totalMisconfig)
              }%;`}
            />
            <div class="info">
              <span>{stat.name}</span>
              <span style="font-weight: bold;">{stat.value}</span>
            </div>
          </div>
        {/each}
      </div>
      <div class="vuln card light">
        <div class="title">
          <div class="icon red">
            <FaSkullCrossbones />
          </div>
          {#if $scanIsLoading}
            <h2>No Review during Scanning</h2>
          {:else}
            <h2>
              {totalMisconfig === 0 ? "Vulnerability" : "Misconfiguration"} Review
            </h2>
          {/if}
        </div>
        <hr />
        {#if vuln}
          <div in:fade class="details">
            <div
              class={`alert ${
                stats.find((s) => s.name === vuln.Severity).color
              }-bg-light`}
            >
              <h3>{totalMisconfig === 0 ? vuln.VulnerabilityID : vuln.ID}</h3>
              <span
                >{totalMisconfig === 0
                  ? vuln.PublishedDate
                  : "Misconfiguration"}</span
              >
            </div>
            <p>{vuln.Title}</p>
            <p>{vuln.Description}</p>
            {#if vuln.Resolution}
              <span class="alert green-bg-light">
                <span style="font-weight: bold;">Resolution: </span>
                <br />
                {vuln.Resolution}
              </span>
            {/if}
            {#each vuln.References as ref}
              <a href={ref}>{ref}</a>
            {/each}
          </div>
        {:else}
          <div class="no-data">
            <div class="icon">
              <FaSkullCrossbones />
            </div>
            {#if $scanIsLoading}
              <h2>No item to review during Scanning</h2>
            {:else}
              <h3>
                No {totalMisconfig === 0 ? "vulnerability" : "misconfiguration"}
                to review
              </h3>
            {/if}
          </div>
        {/if}
      </div>
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
    grid-template-columns: auto 950px;
    gap: var(--margin);
  }

  .menu-items {
    height: calc(100vh - 175px);
    display: grid;
    grid-template-rows: min-content calc(100% - 182px);
  }

  .metadata {
    margin-bottom: var(--margin);
  }

  .metadata .data {
    margin-top: var(--margin);
    display: grid;
    gap: var(--margin);
  }

  .results {
    height: auto;
  }

  .results .table {
    grid-template-columns: repeat(4, 1fr);
  }

  .results .table.body {
    max-height: calc(100% - 115px);
  }

  .results .no-data {
    height: calc(100% - 150px);
  }

  .report {
    height: calc(100vh - 175px);
    display: grid;
    grid-template-rows: 70px auto;
    gap: var(--margin);
  }

  .report .stats {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    gap: var(--margin);
  }

  .report .stats .stat {
    position: relative;
    display: flex;
    align-items: center;
    overflow: hidden;
  }

  .report .stats .stat .status {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
  }

  .report .stats .stat .info {
    position: relative;
    width: 100%;
    display: flex;
    justify-content: space-between;
  }

  .vuln {
    min-height: calc(100% - 50px);
  }

  .vuln .details {
    margin: var(--margin);
    margin-left: 0;
    margin-right: 0;
    height: calc(100% - 65px);
    display: flex;
    flex-direction: column;
    overflow-y: auto;
  }

  .vuln .details * {
    margin-bottom: var(--margin);
  }

  .vuln .details .alert {
    border-radius: var(--radius);
    padding: var(--padding);
    height: min-content;
  }

  .vuln .details p {
    border-top: solid 1px var(--shadow);
    margin: 0;
    padding: 0;
    margin-bottom: var(--margin);
    padding-top: var(--padding);
  }

  .vuln .no-data {
    height: calc(100% - 130px);
  }
</style>
