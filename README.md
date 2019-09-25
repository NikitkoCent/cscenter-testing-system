# cscenter-testing-system [![Actions Status](https://github.com/NikitkoCent/cscenter-testing-system/workflows/CI/badge.svg)](https://github.com/NikitkoCent/cscenter-testing-system/actions) [![GitHub release (latest by date)](https://img.shields.io/github/v/release/NikitkoCent/cscenter-testing-system)](https://github.com/NikitkoCent/cscenter-testing-system/releases/latest) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
Testing system for Computer Science Center C++ course.

## Tests JSON configuration file format
```json
{
  "executable": "path/to/exe",
  "tests":
  [
    {
      "name": "test1",
      "params": ["param1", "param2", "param3"],
      "reference": "path/to/correct-output-file",
      "tested_file": "path/to/tested-file",
      "exitCodes": [ 0, 1, 2, 3 ]
    },
    ...
  ]
}
```

### Fields description
* `name`: Optional. Name of the test that will be displayed. `Unnamed~i` by default where `i` is the test index;
* `executable`: Required. Path to tested executable;
* `params`: Optional. Parameters that will be passed to tested executable at invocation like `path/to/exe param1 param2 param3`. Empty array by default;
* `exitCodes`: Optional. Determines valid exit codes of the executable invocation. `[0]` by default;
* `reference`: Required. File contents of which will be treated as correct for `stdout` output of the executable invocation;
* `tested_file`: Optional. File contents of which will be tested (compared with `reference` file). If ommited, `stdout` stream of the execution will be tested.

### Examples
* Test that check `stdout` of `echo` command:
    ```json
    {
      "executable": "echo",
      "tests":
      [
        {
          "name": "test1",
          "params": ["Hello, world!"],
          "reference": "echo-with-hello-world.txt"
        },
        {
          "name": "test2",
          "reference": "echo-empty.txt"
        }
      ]
    }
    ```
* Test that check correctness of files copying via `cp` command:
    ```json
    {
      "executable": "cp",
      "tests":
      [
        {
          "name": "copy empty file",
          "params": ["existing-empty-file.txt", "empty-copy.txt"],
          "reference": "existing-empty-file.txt",
          "tested_file": "empty-copy.txt"
        },
        {
          "name": "copy this exe (windows-only)",
          "params": ["cstest.exe", "cstest-copy.exe"],
          "reference": "cstest.exe",
          "tested_file": "cstest-copy.exe"
        }
      ]
    }
    ```

## Usage
```shell script
# same: ./cstest -help
./cstest --help
Usage of ./cstest:
  -config string
        Path to tests configuration JSON file

# same: ./cstest --config path/to/config.json
# same: ./cstest -config=path/to/config.json
# same: ./cstest --config=path/to/config.json
./cstest -config path/to/config.json
```
