const filenams = Deno.args;
for (const filename of filenams) {
    const file = await Deno.open(filename);
    await Deno.copy(file, Deno.stdout);
    file.close();
}