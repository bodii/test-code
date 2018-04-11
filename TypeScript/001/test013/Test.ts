import { StringValidator } from './Validation';
import { ZipCodeValidator } from './ZipCodeValidator';
import { LettersOnlyValidator } from './LettersOnlyValidator';

let strings = ['Hello', '98052', '101'];

let validators: { [s: string]: StringValidator; } = {}

validators['Zip code'] = new ZipCodeValidator();
validators['Letters only'] = new LettersOnlyValidator();

strings.forEach(s => {
    for (let name in validators) {
        console.log(`"${ s }" - ${ validators[name].isAcceptable(s) ? "matches" : 'does not match' } ${ name }`);
    }
});