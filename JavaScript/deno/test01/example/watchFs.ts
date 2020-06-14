// 轮询文件系统事件
const watcher = Deno.watchFs("/");
for await (const event of watcher) {
    console.log(">>>> event", event);
}