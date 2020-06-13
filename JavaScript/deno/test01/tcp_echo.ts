const hostname = "0.0.0.0";
const port = 8080;
const listener = Deno.listen({ hostname, port });
console.log(`Listening on ${hostname}:${port}`);
for await (const conn of listener) {
    if (Deno.stdin) {
        Deno.copy(Deno.stdin, conn);
    }
    if (conn) {
        console.log(conn)
        Deno.copy(conn, Deno.stdout);
    }
}