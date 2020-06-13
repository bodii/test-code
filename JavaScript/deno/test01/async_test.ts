import { delay } from "./deps.ts";

Deno.test("async hello world", async () => {
    const x = 1 + 2;

    await delay(100);

    if (x !== 3) {
        throw Error("x should be equal to 3");
    }
});