# GiCom
Check the commit history of your repository or URL's repository.
```txt
Usage: gicom [options] slug path
  
Options:
  -version            		now version
  -filter=<head comment>	head string of git comment
  -mode=<option>      		now "comment", "committer". (with filter option)
```

# howto

```bash
./gicom -url https://github.com/lll-lll-lll-lll/webvtt-reader -filter modify -mode comment
```

```
"modify readme\n"
"modify example code\n"
"modify file content is empty\n"
"modify when last line is not terminal point.\n"
"modify test print json\n"
"modify change readfile method name and file extension is only .vtt\n"
"modify change DeleteVTTElementOfEmptyText name\n"
"modify change filename\n"
"modify how to use this pkg\n"
"modify change vtt file name .\n"
...
```

## Reference
[Go言語でテストしやすいコマンドラインツールをつくる](https://deeeet.com/writing/2014/12/18/golang-cli-test/)<br>
