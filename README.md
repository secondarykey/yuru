It's a tool for that game that I found someday.
Born to defeat "Yuru". 

    go get github.com/secondarykey/yuru
    go install github.com/secondarykey/yuru/cmd/yuru

The previous logic was tagged as v0.
(https://github.com/secondarykey/yuru/tree/v0.0.0)

We are planning to implement UI mode with "v2", and we will create "v1" in preparation for it.


# v2に向けて変更していく点

- 既存コマンドをyuru-cuiにする
- 設定ファイルをHOME(USERPROFILE)に作成
- ドロップの属性を10,20,,,にする
- ボードデータをパッケージ化

# 設定ファイル

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
