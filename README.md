# introduction
It is a Golang static code analyzer, which helps you to improve your code. Now only contains one linter, more linters will be developed in the future(if I have time).

# linters
| linter             | description                                                                             | document        |
|--------------------|-----------------------------------------------------------------------------------------|-----------------|
| loopgoroutinecheck | the loop variable captured by func literal in go statement might have unexpected values | comming soon... |

# how to use it
If you want to use it in local, you can download the repository and build it, and then execute the linter.

```shell
// step 1: download the repository
git clone https://github.com/HandsomeDong/xianggolint

// step 2: run build.sh, if the os of your computer is windows, you can use git bash to run it or build it by yourself instead of build.sh
cd xianggolint
bash build.sh

// step 3: cd to the root of repository you want to scan, then run xxx ./... (./... mean to scan the whole golang project)
cd xxx
/xxx/xxx/xianggolint/output/bin/loopgoroutinecheck ./...

// step 4: waiting for the output, such as...
/Program/test/test.go:8:16: loopgoroutinecheck: loop variable `index` captured by func literal in go statement might have unexpected values
/Program/test/test.go:9:16: loopgoroutinecheck: loop variable `val` captured by func literal in go statement might have unexpected values
/Program/test/test.go:15:16: loopgoroutinecheck: loop variable `i` captured by func literal in go statement might have unexpected values
```
