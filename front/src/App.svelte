<script>
	import Navbar from "./lib/bars/Navbar.svelte";
	import Sidebar from "./lib/bars/Sidebar.svelte";
	import { Router, Route } from "svelte-navigator";
	import DockerImages from "./lib/features/DockerImages.svelte";
	import FileSystem from "./lib/features/FileSystem.svelte";
	import RepoGit from "./lib/features/RepoGit.svelte";
	import Dashboard from "./lib/features/Dashboard.svelte";
	import Report from "./lib/features/Report.svelte";
	import ConfigFiles from "./lib/features/ConfigFiles.svelte";

	const routes = [
		{ path: "/", component: Dashboard },
		{ path: "docker-images", component: DockerImages },
		{ path: "file-system", component: FileSystem },
		{ path: "repo-git", component: RepoGit },
		{ path: "config-files", component: ConfigFiles },
		{ path: "scan-report", component: Report },
	];
</script>

<Navbar />
<main>
	<Router>
		<Sidebar />
		{#each routes as route}
			<Route path={route.path}>
				<svelte:component this={route.component} />
			</Route>
		{/each}
	</Router>
</main>

<style>
	:global(*, ::before, ::after) {
		box-sizing: border-box;
	}

	:root {
		--background: #f1f1f1;
		--shadow: #3a394433;
		--color: #4d4d4d;
		--dark: #272527;
		--dark-light: #333033;
		--light: #fcfcfc;
		--light-dark: #e1e1e1;
		--red: #f15f86;
		--red-light: #f15f8635;
		--green: #a8db75;
		--green-light: #a8db7535;
		--yellow: #fdd765;
		--yellow-light: #fdd76535;
		--orange: #f69666;
		--orange-light: #f6966635;
		--purple: #aa9cf2;
		--purple-light: #aa9cf235;
		--cyan: #78dbe7;
		--cyan-light: #78dbe735;
		--padding: 20px;
		--margin: 20px;
		--radius: 5px;
	}

	:global(body, html) {
		margin: 0;
		padding: 0;
		width: 100%;
		height: 100%;
		background: var(--background);
		overflow: hidden;
		color: var(--color);
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
			Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
	}

	:global(h1, h2, h3, h4) {
		margin: 0;
	}

	:global(hr) {
		color: var(--shadow);
		border: 0px;
		border-top: solid 1px;
	}

	:global(.dark) {
		background: var(--dark);
		color: var(--light);
	}

	:global(.light) {
		background: var(--light);
	}

	:global(.card) {
		padding: var(--padding);
		border-radius: var(--radius);
		box-shadow: 0px 2px 6px var(--shadow);
	}

	:global(.title) {
		display: flex;
		align-items: center;
	}

	:global(.title .icon) {
		width: 35px;
		height: 35px;
		margin-right: var(--margin);
	}

	:global(.table) {
		display: grid;
	}

	:global(.table.header) {
		background: var(--light-dark);
		font-weight: bold;
	}

	:global(.table.body) {
		overflow-y: auto;
	}

	:global(.table span) {
		padding: var(--padding);
		border-bottom: solid 1px var(--shadow);
	}

	:global(.table span.ellipsis) {
		white-space: nowrap;
		text-overflow: ellipsis;
		width: 150px;
		display: block;
		overflow: hidden;
	}

	:global(.no-data) {
		padding: var(--padding);
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		color: var(--light-dark);
	}

	:global(.no-data .icon) {
		height: 180px;
		width: 180px;
		margin: var(--margin);
	}

	:global(input, button, select, textarea) {
		font-family: inherit;
		font-size: inherit;
		box-sizing: border-box;
		border: 1px solid var(--shadow);
		padding: 10px;
		border-radius: var(--radius);
	}

	:global(button:hover) {
		cursor: pointer;
	}

	:global(.selectable) {
		font-weight: bold;
	}

	:global(.selectable:hover) {
		cursor: pointer;
		text-decoration: underline;
	}

	:global(a) {
		text-decoration: none;
		font-weight: 800;
		color: var(--purple);
	}

	:global(a:hover) {
		cursor: pointer;
		text-decoration: underline;
	}

	:global(.red) {
		color: var(--red);
	}

	:global(.green) {
		color: var(--green);
	}

	:global(.yellow) {
		color: var(--yellow);
	}

	:global(.orange) {
		color: var(--orange);
	}

	:global(.purple) {
		color: var(--purple);
	}

	:global(.cyan) {
		color: var(--cyan);
	}

	:global(.red-bg) {
		background: var(--red);
	}

	:global(.green-bg) {
		background: var(--green);
	}

	:global(.yellow-bg) {
		background: var(--yellow);
	}

	:global(.orange-bg) {
		background: var(--orange);
	}

	:global(.purple-bg) {
		background: var(--purple);
	}

	:global(.cyan-bg) {
		background: var(--cyan);
	}

	:global(.red-bg-light) {
		background: var(--red-light);
	}

	:global(.green-bg-light) {
		background: var(--green-light);
	}

	:global(.yellow-bg-light) {
		background: var(--yellow-light);
	}

	:global(.orange-bg-light) {
		background: var(--orange-light);
	}

	:global(.purple-bg-light) {
		background: var(--purple-light);
	}

	:global(.cyan-bg-light) {
		background: var(--cyan-light);
	}

	:global(.lds-ripple) {
		display: inline-block;
		position: relative;
		width: 180px;
		height: 180px;
	}

	:global(.lds-ripple div) {
		position: absolute;
		border: 4px solid var(--red);
		opacity: 1;
		border-radius: 50%;
		animation: lds-ripple 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
	}

	:global(.lds-ripple div:nth-child(2)) {
		animation-delay: -0.5s;
	}

	@keyframes lds-ripple {
		0% {
			top: 90px;
			left: 90px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		4.9% {
			top: 90px;
			left: 90px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		5% {
			top: 90px;
			left: 90px;
			width: 0;
			height: 0;
			opacity: 1;
		}
		100% {
			top: 0px;
			left: 0px;
			width: 180px;
			height: 180px;
			opacity: 0;
		}
	}

	main {
		display: flex;
	}
</style>
