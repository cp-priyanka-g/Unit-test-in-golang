# Unit-test-in-golang

Implementing unit testing in golang

# WHAT IS UNIT TEST?

- A unit test is a function that tests a specific piece of code from a program or package.
  The job of unit tests is to check the correctness of an application, and they are a crucial part of the Go programming language.

In our example in the previous section, the got variable inside the TestMultiply() function is assigned to the result of the Multiply(2, 3) function call. want is assigned to the expected result 6.

The latter part of the test checks if the values of want and got are equal. If not, the Errorf() method is invoked, failing the test.

- Step 1 — Creating a Sample Program to Unit Test
- Step 2 — Writing Unit Tests in Go
  -this test file must always end with \_test.go. By convention, Go testing files are always located in the same folder,
  or package, where the code they are testing resides.

  - The Go language provides a minimal yet complete package called testing that developers use alongside the go test command. The testing package provides some useful conventions, such as coverage tests and benchmarks

- Step 3 — Testing Your Go Code Using the go test command

  - The go test subcommand only looks for files with the \_test.go suffix. go test then scans those file(s) for special functions including func TestXxx.
    Test Commands

  1.  to run the tests for all the packages in the codebase

  - $ go test .

  2.  If you append the -v flag to go test, the test will print out the names of all the executed test functions and the time spent for their execution.

  - $ go test -v

- Step 4 — Writing Table-Driven Tests in Go

  - Using this approach, we get to test several combinations of inputs and their respective output.
  - Now run the test with the -v flag

- Step 5 — Writing Coverage Tests in Go

  - Go provides a built-in method for checking your code coverage. It is often important to know how much of your actual code the tests cover..We’ve managed to achieve 100 percent test coverage for our code, however, we’ve only tested a single function comprehensively.

    - go test -coverprofile=coverage.out

- Step 6 — Writing Benchmarks in Go
  - Benchmarking measures the performance of a function or program. This allows you to compare implementations and to understand the impact of the changes you make to your code
