![logo](figures/git-hours.svg)

![travisCI](https://secure.travis-ci.org/lazypic/git-hours.png)

Git-hours is a command that calculates the working time using the git log data.


## Download & Run
You can download Git-hours from links down below.

- [Windows x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.6/git-hours_windows_x86-64.tgz)
- [macOS x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.6/git-hours_darwin_x86-64.tgz)
- [Linux x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.6/git-hours_linux_x86-64.tgz)

## Set Environment
Put the downloaded file into the bin folder, which is set in the $PATH environment variable.
Then it will be recognized as a subcommand of git.
Because if a command starts with ‘git-‘, git automatically recognizes it as a subcommand of git.

## Install using Golang
```
$ go get -u github.com/lazypic/git-hours
```

## How to use

1. Open the terminal
1. Move to your git repository
1. Type as shown below

```bash
$ git hours
From "2019-03-01 00:00:00 +0900" to "2019-03-31 23:59:59 +0900" : 13h20m9s
```
- The value of timezone offset is automatically set depending on your region.
- By default, the start date is set as the first day of last month and the end date is set as the last day of last month. 

## Detail Options

### help
```
$ git hours -help
```

### since, before
You can set the start date and the end date with this options.
If you don’t enter any value, start date and end date will be set as the first and last date of last month by default.

```bash
$ git hours -since 2019-02-01 -before today
```

If you want to set timezone value, put timezone value at the end of the command as shown below.

```bash
$ git hours -since "2019-03-29 13:55:33 +0800"
```

If you want to customize every value, enter the date, time, and timezone offset value as shown below.

```bash
$ git hours -since "2019-03-01 00:00:00 +0900" -before "2019-03-31 23:59:59 +0900"
```

### Authors
If you want to show specific user data then you can use the `-author` option as shown below.

```bash
$ git hours -author name
```

Also, You can set multiple users setting as shown below.

```bash
$ git hours -author name1,name2
```

### Duration
By default, working time was set to 1 hour between the git commit and the next git commit.
If you want to change the duration, use the `-duration` option as shown below.
In the example is setting 30 minutes.

```bash
$ git hours -duration 0.5h
```

### debug
If you add the -debug option to the command, it prints the time difference between the commit and the commit.
You can view the time, author, and commit details.

```
$ git hours -debug
```

Output example:
```
	 2019-03-31 23:28:40 +0900 kim hanwoong edit go fmt
2m26s >
	 2019-03-31 23:31:06 +0900 kim hanwoong 설명을 추가함.
6m34s >
	 2019-03-31 23:37:40 +0900 kim hanwoong edit comment
1m46s >
	 2019-03-31 23:39:26 +0900 hanwoong kim Update README.md
38s >
	 2019-03-31 23:40:04 +0900 hanwoong kim Update README.md
1m12s >
	 2019-03-31 23:41:16 +0900 hanwoong kim Update README.md
From 2019-02-01 to 2019-03-31 : 13h1m48s
```


## Why did I make it?
https://github.com/kimmobrunfeldt/git-hours

To install the git hours command through the above repository uri, you need to install node.js and dependencies.
I thought that it would be good if there was a tool that can be executed immediately after downloading.

reference model : https://github.com/kimmobrunfeldt/git-hours#how-it-works

The code I wrote before the first commit was set to a model that basically add 1 hour, including programming immersion time.

I like projects that are carried out by telecommuting. For me, it is more convenient than https://toggl.com service. It is also a trust model compared to the toggle service.
