ユルを倒す為に作られたツールです

ユル討伐時は爆弾の処理などツギハギで実装していたものを
より高速にしたものになっています
落としの実装により、かなり高速に算出できるようになりました

    go get github.com/secondarykey/yuru
    go install github.com/secondarykey/yuru/cmd/yuru

でインストールします

$GOPATH/bin/yuru [config file]

で実行します
ファイルを省略するとその場にあるyuru.xml を設定ファイルとして動作します

github.com/secondarykey/yuru/cmd/yuru/yuru.xml

が設定ファイルの例です

```xml
<yuru max="true" startR="0" startC="0">
  <turn>50</turn>                                                             
  <beam>50</beam>
  <board r="5" c="6">
     2,5,5,3,2,1
     4,0,5,1,5,2
     2,5,5,3,2,1
     2,5,5,3,2,1
     2,5,5,3,2,1
  </board>
</yuru>
```

yuru max  = 最大コンボでない場合に再度計算するか？(未実装
yuru startR = 開始位置を指定します(1行目なら1)
yuru startC = 開始位置を指定します(1列目なら1)
turn tag  = 最大ターン数
beam tag  = ビーム幅（大きいほどいろいろな手を試します。もちろん遅くなります
board r   = 盤面の行数
board c   = 盤面の列数
board tag = 盤面

盤面数値でドロップを表して、[,]で区切ります
数値は同じ色を同じ数値にすれば何でもOKです。

考慮してない点

- 盤面が1面同色でも10コンボで算出してしまう(単純に3の倍数で求めています

やりたいこと

- 盤面の解析を行って盤面を読み込む
- 2way を調整を可能に
- 爆弾を必ず消す

コンボ重視、攻撃重視などを設定していけるようにしたいと思っています。
