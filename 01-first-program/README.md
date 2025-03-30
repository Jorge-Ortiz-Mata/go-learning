# Go Course

This is the official repository for the Go Course project

## Notes

Here are notes for this section

### Function

- `main`. Every project must have a main function to execute.

### Packages

There are different packages we can create on Go projects.

- `main`. Every project should have a main package named **main** that will be the entrypoint of the app when you execute it. **Only one file must be have a main function on each project**.

- `utility`. This package can be used on many Go projects such as the **fmt** utility package.

Initialize Go config:
```$
go mod init example/01-first-program
```

Build Go project:
```$
go build
```

Execute Go project:
```$
./01-first-program
```