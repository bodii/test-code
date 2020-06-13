import { serve } from "http/server.ts";

const body = new TextEncoder().encode("Hello world\n");
for await (const req of serve(":8000")) {
    req.respond({ body });
}