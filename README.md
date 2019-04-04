![logo](figures/git-hours.svg)

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

## Detail

### help
```
$ git hours -help
```

### since, before date
기본적으로 아무값도 넣지 않으면 지난달의 시작일, 마지막일로 설정됩니다.
원한다면 사용자가 아래처럼 시작일, 마지막일을 설정할 수 있습니다.

```
$ git hours -since 2019-02-01 -before today
```

타임존값을 입력해야하는 상황이 필요할때는 아래처럼 입력할 수 있습니다.
```
$ git hours -since "2019-03-29 13:55:33 +0800"
```

기본적으로 before값에 날짜만 넣으면 그 날짜의 시작시간이 기준이됩니다.
정확한 견적이 필요할 때는 아래처럼 해당 날짜의 시간, 타임존까지 모두 입력해주세요.
```
$ git hours -since "2019-03-01 00:00:00 +0900" -before "2019-03-31 23:59:59 +0900"
```

### set author
```
$ git hours -author name
```

### set mutliple author
```
$ git hours -author name1,name2
```

### debug mode
```
$ git hours -debug
```

커밋과 커밋간 시간차를 표기합니다, 시간, 작성자, 커밋내용을 볼 수 있습니다.
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

위 리포지터리를 통해서 git hours 명령어를 설치하기 위해서는 node.js 및 의존성을 생각하며 설치해야합니다.
다운로드 하면 바로 사용할 수 있는 툴이 있으면 좋을 것 같아서 제작했습니다.

git 리포지터리에서 아래명령어를 타이핑하면 commit 시간이 출력됩니다.

```
$ git --no-pager log --reverse --date=iso-local --pretty=format:"%ad %an %s" --after="2018-01-01 00:00:00 +0900" --before="2018-12-31 23:59:59 +0900"
```

시간포맷은 git에서 기본적으로 지원하는 ISO8601 방식을 사용했습니다. 작업시간을 계산하는 방식은 아래 형태를 참고했습니다.

https://github.com/kimmobrunfeldt/git-hours#how-it-works

첫번째 커밋전에 작성하는 코드는 프로그래밍 몰입시간을 포함하여 1시간을 기본값으로 셋팅했습니다.
