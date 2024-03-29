

devops 
=== 

_devops_ is a command line tool that makes a copy of the 
[build/cicd](https://github.com/juiicesb/devops/tree/master/build/cicd) tool for configuration and 
deployment of your project. The goal is to help developers get a local copy of _cicd_ without a bunch of copy/paste. 


<!-- toc -->

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
    * [Commands](#commands)
    * [Examples](#examples)
- [Join us on Gopher Slack](#join-us-on-gopher-slack)

<!-- tocstop -->



## Overview

The command line tool creates a copy of the 
[build/cicd](https://github.com/juiicesb/devops/tree/master/build/cicd) tool in your desired project. This 
will allow you to define all your configuration locally in your project repo. 
 
For additional details regarding this tool, refer to 
[github.com/juiicesb/devops](https://github.com/juiicesb/devops)



## Installation

Make sure you have a working Go environment.  Go version 1.2+ is supported.  [See
the install instructions for Go](http://golang.org/doc/install.html).


To install _devops_, simply run:
```
$ go get github.com/juiicesb/devops/cmd/devops
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```



### Compiling locally 

In order to build _devops_, you will need [packr2](https://github.com/gobuffalo/packr/blob/master/v2) to include the 
Golang files and readme from the example [build/cicd](https://github.com/juiicesb/devops/tree/master/build/cicd). 
When code is changed in this project that needs to be included with _build/cicd_, _parkr2_ needs to be ran. 

    Packr is a simple solution for bundling static assets inside of Go binaries. Most importantly it does it in a way 
    that is friendly to developers while they are developing.

To install _packr2_ 
```bash
$ go get -u github.com/gobuffalo/packr/v2/packr2
```

This repo has a post-commit hook at [githooks/post-commit](https://github.com/juiicesb/devops/blob/master/githooks/post-commit) 
to detect changes to `build/cicd` tool and executes _parkr2_. Changed files are included with `git commit --amend`

To install the post-commit hook
```bash
cp githook/post-commit .git/hooks
```


## Usage 

```bash
$ cicd [global options] command [command options] [arguments...]
```

### Global Options 


* Show help 

    `--help, -h`  

* Print the version 

   `--version, -v`  

### Commands

* `inject-build cicd` - Copies the build tool to a target project. This copy the files for [build/cicd](https://github.com/juiicesb/devops/tree/master/build/cicd) 
to your specific project path. 

    The target project must have a `go.mod`. The value for `module` is used to replace all import path for the sub 
    packages copied. You can use `go mod init github.com/my/repo` to create the `go.mod` file. 

    ```bash
    $ devops [global options] inject-build cicd -project NNNNN [command options]
    ``` 
    
    Options: 
    ```bash
    --project value  the root directory of the project, required
    --force          force the files to be copied
    ``` 
    
* `help` - Shows a list of commands
       
    ```bash
    $ devops help
    ```
        
    Or for one command:    
    ```bash
    $ devops inject-build help
    ```


### Examples

Make a copy of the cicd tool for saas-starter-kit 
```bash
$ devops inject-build cicd -project $GOPATH/src/github.com/juiicesb/pilnook
```
