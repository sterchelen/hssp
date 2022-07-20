# hssp
hssp for Http StatuS where the two capitals S replace the two ts of `http`.

## Why?
This CLI is here to help you find/remember the meaning of an http status code.

Historically speaking, this tool was written after struggling with my memory to find the meaning of a code.
Some tools already exist but installing Node.js is too much for me...

## Installation
### From source
To install hssp, first you need to install the dependencies:

* Arch Linux  
```bash
pacman -S go make
```

Then, run:
```bash
make build
make install
```

## Quick start
### Code
```bash
$ hssp code --help
This command displays the given http code description 
with its corresponding class and its RFC.

Usage:
  hssp code [flags]

Flags:
  -h, --help   help for code
```
#### Meaning of 204
```bash
$ hssp code 204
+------+-------------+-------------+---------+
| CODE |    CLASS    | DESCRIPTION |   RFC   |
+------+-------------+-------------+---------+
|  204 | Successfull | No Content  | RFC7231 |
+------+-------------+-------------+---------+
```

### Class
```bash
$ hssp class --help
This command displays the list of http status codes corresponding
to the given class number (1,2,3,4,5).

Usage:
  hssp class [flags]

Flags:
  -h, --help   help for class
```
#### List of status codes for the Successfull class
```bash
$ hssp class 2
+------+-------------+-------------------------------+---------+
| CODE |    CLASS    |          DESCRIPTION          |   RFC   |
+------+-------------+-------------------------------+---------+
|  200 | Successfull | OK                            | RFC7231 |
|  201 | Successfull | Created                       | RFC7231 |
|  202 | Successfull | Accepted                      | RFC7231 |
|  203 | Successfull | Non-Authoritative Information | RFC7231 |
|  204 | Successfull | No Content                    | RFC7231 |
|  205 | Successfull | Reset Content                 | RFC7231 |
|  206 | Successfull | Partial Content               | RFC7231 |
|  207 | Successfull | Multi-Status                  | RFC4918 |
|  208 | Successfull | Already Reported              | RFC5842 |
|  226 | Successfull | IM Used                       | RFC3229 |
+------+-------------+-------------------------------+---------+
```

## Contribution
No other words than "Welcome guys" :)
