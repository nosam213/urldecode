# urldecode
Simple utility for decoding url-encoded URLs.

### Flags
None currently implemented.


## Usage

### Windows
`urldecode.exe "https%3A%2F%2Fexample.com%2Fpage%3Fparameter%3D1%26parameter%3D2"` 

### Unix
`urldecode "https%3A%2F%2Fexample.com%2Fpage%3Fparameter%3D1%26parameter%3D2"`


## Compiling
NOTE: Requires network connection for the first compile.

### Windows
```
C:\Users\username\urldecode> dir
go.mod  go.sum  main.go  README.md

C:\Users\username\urldecode> go build -o urldecode.exe -ldflags="-s -w"
go.mod  go.sum  urldecode.exe  main.go  README.md
```

### Unix
```
/home/user/urldecode:~$ ls
go.mod  go.sum  main.go  README.md

/home/user/urldecode:~$ go build -o urldecode -ldflags="-s -w"
go.mod  go.sum  urldecode  main.go  README.md
```