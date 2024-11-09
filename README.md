# GoとNext.jsを使ってTODOアプリを作ってみよう

## Docker構成
- Go(Gin)
- Next.js
- Mysql

## API設計
- 一覧取得 done
- 新規登録 done
- 更新 done
- 削除 done
- 完了 done

## フロント
<!-- - レイアウト決める -->
<!-- - 一覧取得・表示 -->
<!-- - 更新処理 -->
- 追加ボタン実装
- 新規登録処理
- 完了処理
- 削除処理

## 課題
- Goを修正しても再度go run main.goをして立ち上げ直さないと変更が反映されない
↓
airパッケージをインストールして解決(https://github.com/air-verse/air)

- deleted_atが日本時間にならない
↓
mysqlドライバが悪さをしていた
db, err := sql.Open("mysql", "user:password@tcp(db:3306)/testdb?loc=Asia%2FTokyo")
データベースに接続するときにloc=Asia%2Tokyoとしたらうまくいった

- バリデーション処理
↓
create, update時にcontentが入力されているかチェック
↓
gin  go-playground/validator/v10 を使ってバリデーション
