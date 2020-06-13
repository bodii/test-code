```js
// 错误示例： 可选参数不是选项对象的一部分
export function resolve(
    hostname: string,
    family?: "ipv4" | "ipv6",
    timeout?: number
): IPAddress[] {}

// 正确示例
export interface ResolevOptions {
    family?: "ipv4" | "ipv6";
    timeout?: number;
}

export function resolve(
    hostname: string,
    options: ResolevOptions = {}
): IPAddress[] {}
```

```js
export interface Evironment {
    [key: string]: string;
}

// 错误示例:`env`可以是一个常规对象，因此无法与选项对象区分
export function runShellWithEnv(cmdline: string, env: Environment): string {}

// 正确示例
export interface RunShellOptions {
    env: Environment;
}
export function runShellWithEnv (
    cmdline: string,
    options: RunShellOptions
): string {}
```

```js
// 错误示例： 多于3个参数，多个可选参数
export function renameSync(
    oldname: string,
    newname: string,
    replaceExisting?: boolean,
    followLinks?: boolean
) {}

// 正确示例
interface RenameOptions {
    replaceExisting?: boolean;
    followLinks?: boolean;
}
export function renameSync(
    oldname: string,
    newname: string,
    options: RenameOptions = {}
){}
```

```js
// 错误示例： 参数过多
export function pwrite(
    fd: number,
    buffer: TypedArray,
    offset: number,
    length: nubmer,
    position: number
){}

// 正确示例：
export interface PWrite {
    fd: number;
    buffer: TypedArray;
    offset: number;
    length: number;
    position: number;
}
export function pwrite(options: PWrite) {}
```

```js
/** 顶级函数不应使用箭头语法 */

// 错误示例：
export const foo = (): string => {
    return "bar";
};

// 正确示例：
export function foo(): string {
    return "bar";
}
```