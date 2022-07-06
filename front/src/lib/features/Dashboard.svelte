<script>
    import { link } from "svelte-navigator"
    import Breadcrumb from "../bars/Breadcrumb.svelte";
    import { features } from "../../stores/features"
    import FaLaptopCode from 'svelte-icons/fa/FaLaptopCode.svelte'
    import FaSpaceShuttle from 'svelte-icons/fa/FaSpaceShuttle.svelte'
    import FaGithub from 'svelte-icons/fa/FaGithub.svelte'
    import FaBookMedical from 'svelte-icons/fa/FaBookMedical.svelte'

    let breadcrumb = ["Dashboard"]

    const subjects = [
        {name: "About", icon: FaLaptopCode, text: "Docktor is a Web App that deploys an easy-to-use kit of analysis and scanning tools.\n\nToday, developers use a variety of resources. It is more and more difficult to ensure the security of our artifacts.\nEspecially Docker environments which are an obvious source of vulnerability.\n\nThe objective is to have a simple, fast, lightweight and everywhere approach to ensure the security of our productions."},
        {name: "What's New ?", icon: FaGithub, text: "All the news and documentation on the official Github repo:\n\n<a href='https://github.com/Matbabs/Docktor'>https://github.com/Matbabs/Docktor</a>"},
        {name: "Getting Started", icon: FaSpaceShuttle, text: "To start, choose a scan environment: Docker Images, File System, Repo GIT, Config Files."},
    ]
</script>

<div class="container">
    <Breadcrumb bind:paths={breadcrumb}/>
    <div class="menu-container">
        <div class="menu">
            {#each $features as feature, i}
                {#if i !== 0}
                    <a class="menu-item card light" href={feature.path} use:link>
                        <div class={`icon ${feature.color}`}>
                            <svelte:component this={feature.icon}/>
                        </div>
                        <h3>{feature.name}</h3> 
                    </a>   
                {/if}
            {/each}
        </div>
        <div class="home card light">
            <div class="title">
                <div class="icon">
                    <FaBookMedical/>
                </div>
                <h1>Welcome to Docktor</h1>
            </div>
            <hr>
            {#each subjects as subject, i}
                <div class="subject">
                    <div class="icon red">
                        <svelte:component this={subject.icon}/>
                    </div>
                    <div class="explanations">
                        <h2>{subject.name}</h2> 
                        <p>
                            {@html subject.text?.replace(/\n/g, '<br>')}
                        </p>
                    </div>
                </div>
                {#if i !== subjects.length - 1}
                    <hr>
                {/if}
            {/each}
        </div>
    </div>
</div>

<style>

    a {
        text-decoration: none;
        color: var(--color);
    }

    .container {
        width: 100%
    }

    .menu-container {
        display: grid;
        grid-template-columns: 1fr 1fr;
    }

    .menu {
        padding: var(--padding);
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: var(--margin);
    }

    .menu-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .menu-item .icon {
        width: 70px;
        height: 70px;
        margin-bottom: var(--margin);
        transition: all 250ms;
    }

    .menu-item:hover {
        cursor: pointer;
    }

    .menu-item:hover .icon{
        transform: scale(1.3);
    }

    .home {
        height: calc(100vh - 175px);
        margin: var(--margin);
        margin-left: 0;
    }

    .subject {
        margin: var(--margin);
        margin-left: 0;
        display: grid;
        grid-template-columns: 60px auto;
    }

    .subject .icon {
        width: 60px;
        height: 60px;
    }

    .subject .explanations {
        margin-left: var(--margin);
    }

    .subject .explanations p {
        margin: var(--margin);
        margin-left: 0;
        margin-right: 0;
    }
</style>