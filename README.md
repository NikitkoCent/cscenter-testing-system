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
      "exitCodes": [ 0, 1, 2, 3 ]
    },
    ...
  ]
}
```

### Fields description
* `executable`: Required. Path to tested executable
* `name`: Optional. Name of the test that will be displayed. `Unnamed~i` by default where `i` is test index.
* `reference`: Required. File contents of which will be treated as correct for `stdout` output of the executable invocation.
* `params`: Optional. Parameters that will be passed to tested executable at invocation like `path/to/exe param1 param2 param3`. Empty array by default.
* `exitCodes`: Optional. Determines valid exit codes of the executable invocation. `[0]` by default.

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
