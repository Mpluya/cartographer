<script>
    import mermaid from 'mermaid';
    import {afterUpdate, onMount} from 'svelte';

    export let doc = ""
    let graph = null;

    mermaid.initialize({
        startOnLoad: false,
        theme: 'forest',
    });

    const rerender = () => {
        mermaid.mermaidAPI.render('graph-div', doc, (svgCode) => {
            graph.innerHTML = svgCode;
        });
    }

    onMount(() => {
        rerender()
    });

    afterUpdate(() => {
        rerender()
    })
</script>

<pre bind:this={graph}>
</pre>
