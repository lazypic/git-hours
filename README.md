# git-hours

git log 를 이용해서 작업시간을 계산하는 명령어 입니다.


## Download & Run
- [Windows 64bit](https://github.com/lazypic/git-hours/releases/download/v0.0.3/git-hours_windows.tgz)
- [macOS 64bit](https://github.com/lazypic/git-hours/releases/download/v0.0.3/git-hours_darwin.tgz)
- [Linux 64bit](https://github.com/lazypic/git-hours/releases/download/v0.0.3/git-hours_linux.tgz)

각 OS에 맞는 명령어를 다운로드 받습니다.

터미널에서 명령어를 자동인식 할 수 있는 $PATH 환경변수에 잡혀있는 bin 폴더에 다운받은 git-hours 명령어를 넣어줍니다.
명령어 이름이 "git-hours"형태이기 때문에 git이 자동으로 서브명령어로 인식합니다.

## Install use go
```
$ go get -u github.com/lazypic/git-hours
```

## How to use
git 리포지터리로 이동해서 아래처럼 터미널에서 타이핑해주세요.

```
$ git hours
From 2019-03-01 to 2019-03-31 : 80h29m10s
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
$ git hours -since 2019-02-01 -before 2019-02-28
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
