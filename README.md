# git-hours

git log 를 이용해서 작업시간을 계산하는 유틸리티.

## Roadmap
리포지터리에서 아래처럼 타이핑하면 commit 시간이 출력된다.
시간포맷은 git에서 기본적으로 지원하는 ISO8601 방식을 채택했다.

```
$ git --no-pager log --reverse --date=iso --pretty=format:"%ad %an %s" --after="2018-01-01 00:00:00 +0900" --before="2018-12-31 23:59:59 +0900"
```

계산 모델은 지금까지는 아래 모델이 가장 마음에 든다.(그나마 지지를 받는 모델)

https://github.com/kimmobrunfeldt/git-hours#how-it-works

위 리포지터리에서 git hour 를 설치하는것은 꽤나 불편하다.
개인적인 견적을 위해서 간단하게 작성해볼 예정이다.

- 앞에 한시간을 더할 때
    - 커밋사이가 2시간 이상차이가 나는 뒷쪽 커밋
- 각 커밋간격이 2시간 미만이면 작업중으로 체크한다.
    - 기존시간에 작업시간을 더한다.
