Deno.test({
    name: "leaky test",
    fn() {
        Deno.open("hello.txt");
    },
    sanitizeResources: false,
    sanitizeOps: false,
});