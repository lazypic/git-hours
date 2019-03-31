# git-hours

git log 를 이용해서 작업시간을 계산하는 유틸리티.


## Install

## How to use.
아래처럼 타이핑하면 된다.

```
$ git hours
2019-03-01~2019-03-31 : 80h29m10s
```

## Detail

help
```
$ git hours -help
```

set start, end

```
$ git hours -start 2018-01-01 -end 2019-12-31
```

set author
```
$ git hours -author woong
```

set mutliple author
```
$ git hours -author woong,kybin
```

timezone offset
기본적으로 현재 지역의 타임존 시간을 인식함.
```
$ git hours -zone +0900
```

debug mode
```
$ git hours -debug
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
