import { assertEquals } from "./deps.ts";

Deno.test("hello world", () => {
    const x = 1 + 2;
    assertEquals(x, 3);
});