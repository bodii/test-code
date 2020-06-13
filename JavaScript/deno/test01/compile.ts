
const [diagnostics, emitMap] = await Deno.compile(
    "/foo.ts", 
    {
        "/foo.ts": `import * as bar from "./bar.ts";\nconsole.log(bar);\n`,
        "/bar.ts": `export const bar = "bar"\n`,
    }
);

console.assert(diagnostics == null); // 确保没有返回诊断信息
console.log(emitMap);