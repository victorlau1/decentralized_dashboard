const fs = require("fs")
import fs from "fs";


//TODO: Add in a waiter for the DOM to render the next tab
// Does not work as intented
// https://ethernodes.org/nodes
const Scraper =  () => {
    
    const now = Date.now();
    const fileName = `ethernodes_${now}`
    var fileToWrite = ''
    var button = document.getElementsByClassName("paginate_button page-item next")[0]
    var counter = 2;
    var next = document.getElementsByClassName("paginate_button page-item next disabled").length
    while (counter > 0) {
        fileToWrite += ParseTable()
        ClickNext()
    }
    return totalFile
}

const ParseTable = () => {
    var totalFile = ''
    const rows = document.getElementsByTagName("table")[0].rows;
            for (var rowNum in rows) {
                if (rows[rowNum].innerText !== undefined && rows[rowNum].innerText != 'Id\tHost\tISP\tCountry\tClient\tVersion\tOS\tLast Seen\tIn Sync') {
                    totalFile += rows[rowNum].innerText + '\n'
                }
            }
    return totalFile;
}

const ClickNext = () => {
    button = document.getElementsByClassName("paginate_button page-item next")[0].click()
}

const writeToFile = (fileName) => {
    fs.writeFile(`~/workspace/scraper/${fileName}`, data, (err) => {
        if (err) throw err;
    })


export default  {
    NextButton,
    Scraper
}