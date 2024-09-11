'use strict';

function matchingStrings(stringList, queries) {
    const result = [];
    for (let i = 0; i < queries.length; i++) {
        let count = 0;
        for (let j = 0; j < stringList.length; j++) {
            if (queries[i] === stringList[j]) {
                count++;
            }
        }
        result.push(count);
    }
    return result;

}

function main() {
    const stringList = ['ab', 'ab', 'abc'];
    const queries = ['ab', 'abc', 'bc'];
    const result = matchingStrings(stringList, queries);
    console.log(result);
}
main();