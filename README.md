# cscenter-testing-system
Testing system for Computer Science Center C++ course.

## Tests JSON file format
```json
{
  "executable": "path/to/exe",
  "tests":
  [
    {
      "name": "test1",
      "params": "param1 param2 param3",
      "reference": "path/to/correct-output-file",
      "exitCodes": [ 0, 1, 2, 3 ]
    },
    ...
  ]
}
```

### Fields description
* `executable`: Required. Path to tested executable
* `name`: Optional. Name of the test that will be displayed. Value of the field `params` by default.
* `reference`: Required. File contents of which will be treated as correct for `stdout` output of the executable invocation.
* `params`: Required. Parameters that will be passed to tested executable at invocation like `path/to/exe param1 param2 param3`.
* `exitCodes`: Optional. Determines valid exit codes of the executable invocation. `[0]` by default.
