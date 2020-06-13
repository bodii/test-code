Deno.test({
    name: "do macOS feature",
    ignore: Deno.build.os != "darwin",
    fn() {
        doMacOSFeature();
    },
});