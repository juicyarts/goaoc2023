# Advent of Code 2023

Advent of Code 2023 with Golang. All challenges can be found [here](https://adventofcode.com/2023).
The actual results for the input tests live in `.env` files in each of the folders. You need to create those for the tests to run properly.

There is no `runnable` code in this repository, everything is handled by the tests.

## Run all Tests

This might take a while, some days are inneficiently solved since efficiency is not my main goal of this year but rather getting more familiar with Golang.

```bash
make test
```

## Run a specific day

```bash
cd dayX
go test
```
