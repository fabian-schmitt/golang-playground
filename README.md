# golang-playground
Learning TDD in GO - https://quii.gitbook.io/learn-go-with-tests/


# Writing tests

Writing a test is just like writing a function, with a few rules
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
For now it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.