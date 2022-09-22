import { mkdir, readFile, writeFile, readdir } from "node:fs/promises";
import * as path from "node:path";
import { fileURLToPath } from "node:url";

const thisFilePath = fileURLToPath(import.meta.url);
const baseDir = path.normalize(path.join(path.dirname(thisFilePath), ".."));
const buildDir = path.join(baseDir, "build");
const contractDir = path.join(buildDir, "contracts");
const destDir = path.join(buildDir, "abi");

/**
 * Reduce artifact abi json for abigen
 */
async function main() {
  mkdir(destDir, { recursive: true }).then(() => 
    readdir(contractDir, { withFileTypes: true }).then(entries => 
      entries.filter(entry => !entry.isDirectory()).forEach(file =>
        readFile(path.join(contractDir, file.name))
          .then(JSON.parse)
          .then(json => json.abi)
          .then(JSON.stringify)
          .then(content => writeFile(path.join(destDir, file.name), content))
      )
    )
  );
};

main();
