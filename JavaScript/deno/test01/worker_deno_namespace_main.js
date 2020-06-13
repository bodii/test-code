const worker = new Worker('./worker.js', { type: "module", deno: true });
worker.postMessage({ filename: "./log.txt" });