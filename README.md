# git-hours

git log 를 이용해서 작업시간을 계산하는 명령어 입니다.


## Download & Install
- Windows
- [macOS](https://github.com/lazypic/git-hours/releases/download/v0.0.1/git-hours)
- Linux

각 OS에 맞는 명령어를 다운로드 받습니다.(준비중)

터미널에서 명령어를 자동인식 할 수 있는 $PATH 환경변수에 잡혀있는 bin 폴더에 다운받은 git-hours 명령어를 넣어줍니다.
명령어 이름이 "git-hours"형태이기 때문에 git이 자동으로 서브명령어로 인식합니다.

## How to use.
터미널에 아래처럼 간단하게 타이핑해주세요.

```
$ git hours
From 2019-03-01 to 2019-03-31 : 80h29m10s
```

## Detail

### help
```
$ git hours -help
```

### start, end date
기본적으로 아무값도 넣지 않으면 지난달의 시작일, 마지막일로 설정됩니다.
원한다면 사용자가 아래처럼 시작일, 마지막일을 설정할 수 있습니다.

```
$ git hours -start 2019-02-01 -end 2019-02-28
```

### set author
```
$ git hours -author woong
```

### set mutliple author
```
$ git hours -author woong,kybin
```

### timezone offset value
사용자가 이용하는 타임존 시간을 인식하고 기본값으로 사용합니다.
```
$ git hours -zone +0900
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


## Plan
git 리포지터리에서 아래명령어를 타이핑하면 commit 시간이 출력된다.

```
$ git --no-pager log --reverse --date=iso --pretty=format:"%ad %an %s" --after="2018-01-01 00:00:00 +0900" --before="2018-12-31 23:59:59 +0900"
```

시간포맷은 git에서 기본적으로 지원하는 ISO8601 방식을 사용했다.


작업시간을 계산하는 방식은 아래 형태를 참고했다.

https://github.com/kimmobrunfeldt/git-hours#how-it-works

팀내 많은 개발자가 위 리포지터리의 git hour 명령어를 설치하는것은 많은 의존성이 물려있서 이 명령어를 만들었다.
