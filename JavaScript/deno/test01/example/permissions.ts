// 查找一个权限
const status = await Deno.permissions.query({ name: "write" });
if (status.state !== 'granted') {
    throw new Error("need write permission");
}

const log = await Deno.open("request.log", { write: true, append: true });

// 放弃一些权限
await Deno.permissions.revoke({ name: "read" });
await Deno.permissions.revoke({ name: "write" });

// 使用日志文件
const encoder = new TextEncoder();
await log.write(encoder.encode("hello\n"));

// 这将会失败
await Deno.remove("request.log");