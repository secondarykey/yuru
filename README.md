It's a tool for that game that I found someday.
Born to defeat "Yuru". 

    go get github.com/secondarykey/yuru
    go install github.com/secondarykey/yuru/cmd/yuru
    go install github.com/secondarykey/yuru/cmd/yuru-cui -> v0 yuru

The previous logic was tagged as v0.
(https://github.com/secondarykey/yuru/tree/v0.0.0)

We are planning to implement UI mode with "v1".

https://user-images.githubusercontent.com/445407/119734833-7c14df00-beb6-11eb-858d-e97cbd747666.mp4

# 大きく変わった点

- 既存コマンドをyuru-cui
- 設定ファイルをHOME(USERPROFILE)に作成
- ドロップの属性を10の倍数にした
- ボードデータをパッケージ化
- 計算はlogic,データはdtoにした。
- BoardInfoを削除

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

# 考慮してない点

- 盤面が1面同色でも10コンボで算出してしまう(単純に3の倍数で求めています

# やりたいこと

- 盤面の解析を行って盤面を読み込む
- 2way を調整を可能に
- 爆弾を必ず消す
- 陣機能の作成(Random reset)
- 保存機能の作成(Save Board)
- 再生機能の作成(Play Board)
