<script type="module">

    import {onMount} from "svelte";
    import {editor, Uri} from 'monaco-editor'

    let inst
    let editorContainer
    export let document

    // The uri is used for the schema file match.
    const modelUri = Uri.parse('https://cartographer.sh/example_sc_1.yaml');

    onMount(() => {
        inst = editor.create(editorContainer,
            {
                automaticLayout: true,
                model: editor.createModel(document, 'yaml', modelUri),
            },
        )
        inst.onDidChangeModelContent = (e) => {
            document = inst.getValue()
        }
    })


</script>

<div class="monaco-editor-container" bind:this={editorContainer} style="height: 400px"></div>

<style>
    .monaco-editor-container {
        outline: lightgray solid 1px;
    }
</style>