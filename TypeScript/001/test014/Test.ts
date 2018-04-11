/// <reference path="Validation.ts" />
/// <reference path="LettersOnlyValidator.ts" />
/// <reference path="ZipCodeValidator.ts" />

let strings = ['Hello', '98052', '101'];

let validations: { [s: string]: Validation.StringValidator } = {};
validations['ZIP code'] = new Validation.ZipCodeValidator();
validations['Letters only'] = new Validation.LettersOnlyValidator();

for (let s of strings) {
    for (let name in validations) {
        console.log(`"${ s }" - ${ validations[name].isAcceptable(s) ? "matches": "does not match" } ${ name }`);
    }
}

