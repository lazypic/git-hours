![logo](figures/git-hours.svg)

![travisCI](https://secure.travis-ci.org/lazypic/git-hours.png)

git-hours is a command that calculates the working time using the git log data.


## Download & Run
- [Windows x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.5/git-hours_windows_x86-64.tgz)
- [macOS x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.5/git-hours_darwin_x86-64.tgz)
- [Linux x86-64](https://github.com/lazypic/git-hours/releases/download/v0.0.5/git-hours_linux_x86-64.tgz)

Download the command for your OS.

put the downloaded command into the bin folder, which is held in the $PATH environment variable.
Because the command name is of the form "git-hours", git automatically recognizes it as a subcommand.

## Install use go
```
$ go get -u github.com/lazypic/git-hours
```

## How to use
Go to the git repository and type in the terminal as shown below.
- The time zone offset value is automatically supported depending on the region
- By default, the start and end dates of the last month are specified.

```
$ git hours
From "2019-03-01 00:00:00 +0900" to "2019-03-31 23:59:59 +0900" : 13h20m9s
```

## Detail Options

### help
```
$ git hours -help
```

### since, before
By default, if no value is entered, it is set to the start date and last day of the last month.
If desired, the user can set the start date and the last date.

```
$ git hours -since 2019-02-01 -before today
```

If you want to time zone value, you can type:
```
$ git hours -since "2019-03-29 13:55:33 +0800"
```


If you need an accurate quote, enter the date, time, and time zone offset value as shown below.
```
$ git hours -since "2019-03-01 00:00:00 +0900" -before "2019-03-31 23:59:59 +0900"
```

### author
It can be calculated by specifying only the desired user.
```
$ git hours -author name
```

### mutliple author
You can calculate multiple users.

```
$ git hours -author name1,name2
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

For me, it is more convenient than https://toggl.com service.
