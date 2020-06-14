// 创建子进程
const p = Deno.run({
    cmd: ["echo", "hello"],
});

// 等待完成
await p.status();
