const hostname = "0.0.0.0";
const port = 8080;
const lister = Deno.listen({ hostname, port });
console.log(`Listening on ${hostname}:${port}`);
for await (const conn of lister) {
    Deno.copy(conn, conn);
}