for await (const _ of Deno.signal(Deno.Signal.SIGINT)) {
    console.log("   !");
}

/** 如果想要停止监控信号，可以使用信号对象的dispose()方法 */
const sig = Deno.signal(Deno.Signal.SIGINT);
setTimeout(
    () => {
        sig.dispose();
    },
    5000
);

for await (const _ of sig) {
    console.log("interrupted");
}