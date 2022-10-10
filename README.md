# GiCom
Check the commit history of your repository or URL's repository.
```txt
Usage: gitcomicome [options] slug path
  
Options:
  -version            		now version
  -filter=<head comment>	head string of git comment
  -mode=<option>      		now "comment", "committer". (with filter option)
```

# Install
`brew tap lll-lll-lll-lll/gitcomicome`<br>
`brew install lll-lll-lll-lll/gitcomicome/gitcomicome`

# howto

```bash
gitcomicome -url https://github.com/lll-lll-lll-lll/webvtt-reader -filter modify -mode comment
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

## 動機
cli作ったことなかったので、作ってみたくなった<br>
指定リポジトリのコミット履歴から指定の文字列にフィルタリングする<br>
理想はコミットの書き方にルールを設ける<br>

## Reference
[Go言語でテストしやすいコマンドラインツールをつくる](https://deeeet.com/writing/2014/12/18/golang-cli-test/)<br>
[homebrewに上げる方法](https://qiita.com/kcwebapply/items/4777dfc9151ebb3e8a19)


## リリースまでの流れ
- コードの変更
- タグ作成. git tag -a v0.1.2 -m 'TestRelease'
- tag push. git push origin v0.1.2 