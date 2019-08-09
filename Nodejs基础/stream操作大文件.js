const fs = require('fs')

const reader = fs.createReadStream('bigfile.txt')
const writer = fs.createWriteStream('bigfile-2.txt')
reader.pipe(writer)