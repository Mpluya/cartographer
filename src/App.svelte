<script>
    import Viz from "./lib/Viz.svelte";
    import SupplyChainEditor from "./lib/SupplyChainEditor.svelte";
</script>

<main>
    <h1>Cartographer Live Editor</h1>

    <SupplyChainEditor></SupplyChainEditor>

    <Viz/>

</main>

<style global>
    :root {
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    }

    h1 {
        text-align: center;
        color: #3e00ff;
        text-transform: uppercase;
        font-size: 2rem;
        font-weight: 200;
        line-height: 1.1;
        margin: 2rem auto;
        max-width: 14rem;
    }

    @media (min-width: 480px) {
        h1 {
            max-width: none;
        }
    }
</style>
