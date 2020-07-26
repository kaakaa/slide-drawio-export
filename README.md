# VS Code Draw.io Integrationで作成した図を画像ファイルとして取り込む

## Agenda
* [背景] ドキュメンテーション時の作図ツールにVS Code Draw.io Integratioは有効そう
* [疑問] スタンドアローンで起動しているのか？
* [疑問] Draw.ioで作図したファイルを画像ファイルかするには？
* [課題] VS Code Draw.io Integrationではエクスポート機能が提供されていい
* [解決の方法] Draw.ioファイルを画像ファイルに一括変換するツールの開発
  * [kaakaa/dio\-exporter: dio\-exporter is a CLI tool for exporting \.drawio or \.dio file as image file\.](https://github.com/kaakaa/dio-exporter)
* [さらなる解法] `*.drawio.png`, `*.drawio.svg` ファイル
* [調査] `*.drawio.png` ファイルの中身

この資料は [Marpit](https://marpit.marp.app/)を使用して作成されています。

* [VisualStudio Code](https://code.visualstudio.com/)
* [Marp for VS Code](https://marketplace.visualstudio.com/items?itemName=marp-team.marp-vscode)

## 参考

* VS Code Draw.io Integration
  * [jgraph/drawio: Source to app\.diagrams\.net](https://github.com/jgraph/drawio)
  * [hediet/vscode\-drawio: This unofficial extension integrates Draw\.io into VS Code\.](https://github.com/hediet/vscode-drawio)
  * [VSCodeでDraw\.ioが使えるようになったらしい！ \- Qiita](https://qiita.com/riku-shiru/items/5ab7c5aecdfea323ec4e)
  * [Visual Studio Code \- \*\.drawio\.svg や \*\.drawio\.png の衝撃 \- anfangd's blog](https://blog.anfangd.me/entry/2020/07/08/220628)
* Marp
  * [【VS Code \+ Marp】Markdownから爆速・自由自在なデザインで、プレゼンスライドを作る \- Qiita](https://qiita.com/tomo_makes/items/aafae4021986553ae1d8)
* misc
  * [PNGを読む \- Qiita](https://qiita.com/kouheiszk/items/17485ccb902e8190923b)
