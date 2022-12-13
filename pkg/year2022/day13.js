let lines = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`;


function compareArray(left, right) {
    for (let i = 0; i < left.length; i++) {
        //console.log("\t- Compare", left, " vs ", right);
        if (right[i] === undefined) {
            //console.log("\t\t- Right side ran out of items, so inputs are not in the right order");
            return false;
        }
        if (Array.isArray(left[i])) {
            if (Array.isArray(right[i])) {
                //console.log("\t\t- Compare ", left, " vs ", right);
                let next = compareArray(left[i],right[i]);
                if (next == 'pass') {
                    //console.log("next in loop")
                    continue;
                }
                return next;
            }
            if (Number.isInteger(right[i])) {
                //console.log("\t\t- Mixed types; convert right to [", right[i], "] and retry comparison");
                let next =  compareArray(left[i], [right[i]]);
                if (next == 'pass') {
                    //console.log("next in loop")
                    continue;
                }
                return next;
            }
        }

        let next = compareLeftIsNumber(left[i], right[i]);
        if (next == 'pass') {
            //console.log("next in loop")
            continue;
        }
        return next;
    }
    if (left.length == right.length) {
        return 'pass';
    }
    //console.log("\t\t- Left side ran out of items, so inputs are in the right order")
    return true;
}

function compareLeftIsNumber(left, right) {
    if (left == right) {
        //console.log("\t- Compare", left, " vs ", right);
        //console.log("passing");
        return 'pass';
    }
    if (Number.isInteger(right)) {
        //console.log("\t- Compare", left, " vs ", right);
        return left < right;
    }
    // right is array
    //console.log("\t\t- Mixed types; convert left to [", left, "] and retry comparison");
    return compareArray([left],right);
}

function compare(left, right) {
    if (Array.isArray(left)) {
        if (Array.isArray(right)) {
            //console.log("\t\t- Compare ", left, " vs ", right);
            return compareArray(left, right);
        }
        if (Number.isInteger(right)) {
            //console.log("\t\t- Mixed types; convert right to [", right, "] and retry comparison");
            return compareArray(left, [right]);
        }
    }
    return compareLeftIsNumber(left, right)
}

function part1() {
    let pairs = lines.split("\n\n");
    let trueIndexes = [];
    for (let i = 0; i < pairs.length; i++) {
        let packetStrs = pairs[i].split("\n");
        let left = JSON.parse(packetStrs[0]);
        let right = JSON.parse(packetStrs[1]);
        //console.log(" == Pair ", i+1, " ==")
        //console.log("- Compare ", left, " vs ", right);
        let res = compare(left, right);
        if (res == true) {
            trueIndexes.push(i+1);
        }
        //console.log(res);
    }

    let indexSum = 0;
    trueIndexes.forEach(i => indexSum += i);
    return indexSum;
}


function part2() {
    let strings = lines.split("\n");
    let packets = [];
    for (i = 0; i < strings.length; i++) {
        if (strings[i].length == 0) {
            continue;
        }
        packets.push(JSON.parse(strings[i]));
    }
    const p1 =  [[2]];
    const p2 = [[6]];
    packets.push(p1);
    packets.push(p2);

    packets.sort(function(a, b) {
        if (compare(a, b) == true) {
            return -1;
        }
        if (compare(a, b) == false) {
            return 1;
        }
        return 0;
    })

    p1val = 0;
    for (i = 0; i < packets.length; i++) {
        if (packets[i] === p1) {
            p1val = i+1;
        }
        if (packets[i] == p2) {
            p2val = i+1;
        }
    }

    return p1val * p2val;
}

part1();
part2();