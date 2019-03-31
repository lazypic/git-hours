# git-hours

git log 를 이용해서 작업시간을 계산하는 유틸리티.


## Install
명령어를 다운로드 받습니다. 터미널에서 명령어를 자동인식 할 수 있는 $PATH 환경변수에 잡혀있는 bin 폴더에 다운받은 git-hours 명령어를 넣어줍니다.
명령어 이름이 "git-hours"형태이기 때문에 git이 자동으로 서브명령어로 인식, 작동됩니다.

## How to use.
터미널에 아래처럼 간단하게 타이핑해주세요.

```
$ git hours
From 2019-03-01 to 2019-03-31 : 80h29m10s
```

## Detail

help
```
$ git hours -help
```

시작일, 마지막날짜를 설정하는 방법
기본값은 지난달의 시작일, 마지막 날짜로 설정되어있습니다.


```
$ git hours -start 2019-02-01 -end 2019-02-28
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
