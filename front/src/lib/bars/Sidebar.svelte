<script>
  import { link } from "svelte-navigator";
  import { onMount } from "svelte";
  import { features } from "../../stores/features";
  import FaDocker from "svelte-icons/fa/FaDocker.svelte";
  import FaFileMedicalAlt from "svelte-icons/fa/FaFileMedicalAlt.svelte";
  import FaLaptopMedical from "svelte-icons/fa/FaLaptopMedical.svelte";
  import FaGitAlt from "svelte-icons/fa/FaGitAlt.svelte";
  import FaTools from "svelte-icons/fa/FaTools.svelte";

  onMount(() => {
    features.set([
      { name: "Dashboard", icon: FaLaptopMedical, path: "/", color: "green" },
      {
        name: "Docker Images",
        icon: FaDocker,
        path: "docker-images",
        color: "yellow",
      },
      {
        name: "File System",
        icon: FaFileMedicalAlt,
        path: "file-system",
        color: "orange",
      },
      {
        name: "Config Files",
        icon: FaTools,
        path: "config-files",
        color: "cyan",
      },
      { name: "Repo GIT", icon: FaGitAlt, path: "repo-git", color: "purple" },
    ]);
  });
</script>

<div class="sidebar dark">
  {#each $features as feature}
    <a href={feature.path} use:link>
      <div class="menu-item title">
        <div class={`icon ${feature.color}`}>
          <svelte:component this={feature.icon} />
        </div>
        <h4>{feature.name}</h4>
      </div>
    </a>
  {/each}
</div>

<style>
  a {
    text-decoration: none;
    color: var(--light);
  }

  .sidebar {
    width: 270px;
    height: 100vh;
    background: var(--dark-light);
  }

  .menu-item {
    padding: var(--padding);
  }

  .menu-item .icon {
    width: 30px;
    height: 30px;
    transition: all 250ms;
  }

  .menu-item:hover .icon {
    transform: scale(1.3);
  }
</style>
