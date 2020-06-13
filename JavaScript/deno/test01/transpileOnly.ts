const result = await Deno.transpileOnly({
    '/foo.ts': `enum Foo { Foo, Bar, Baz }; \n`,
});

console.log(result['/foo.ts'].source);
console.log(result['/foo.ts'].map);