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
+------+------------+-------------+---------+
| CODE |    CLASS   | DESCRIPTION |   RFC   |
+------+------------+-------------+---------+
|  204 | Successful | No Content  | RFC2616 |
+------+------------+-------------+---------+
```

### Class
```bash
$ hssp class --help
This command displays the list of http status codes corresponding
to the given class, which may be specified as a number (1-5),
a class category string (1xx, 2xx, 3xx, 4xx, 5xx),
or the class name, i.e. informational, successful, redirect, clienterror, or servererror

Usage:
  hssp class [flags]

Flags:
  -h, --help   help for class
```
#### List of status codes for the Successful class
```bash
$ hssp class 2
+------+-------------+-------------------------------+---------+
| CODE |    CLASS    |          DESCRIPTION          |   RFC   |
+------+-------------+-------------------------------+---------+
|  200 | Successful | OK                            | RFC2616 |
|  201 | Successful | Created                       | RFC2616 |
|  202 | Successful | Accepted                      | RFC2616 |
|  203 | Successful | Non-Authoritative Information | RFC2616 |
|  204 | Successful | No Content                    | RFC2616 |
|  205 | Successful | Reset Content                 | RFC2616 |
|  206 | Successful | Partial Content               | RFC2616 |
|  207 | Successful | Multi-Status                  | RFC4918 |
|  208 | Successful | Already Reported              | RFC5842 |
|  226 | Successful | IM Used                       | RFC3229 |
+------+-------------+-------------------------------+---------+
```

## Contribution
No other words than "Welcome guys" :)
