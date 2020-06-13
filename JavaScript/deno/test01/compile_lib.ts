const [ errors, emitted ] = await Deno.compile({
    "main.ts",
    {
        "main.ts": `document.getElementById("foo");\n`,
    },
    {
        lib: ['dom', 'esnext'],
    }
});

console.assert(errors == null);
console.log(emitted);