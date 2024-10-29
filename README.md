# GoとNext.jsを使ってTODOアプリを作ってみようの巻

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
