# クリーンアーキテクチャ用自動ファイル生成（Java）

## Content
下記ファイルの自動生成を行います。
- Controller
- Interactor:prod（domain/applicationService）
- Interactor:dev（stubs）
- Usecase interface
- InputData class（for usecase）
- OutputData class（for usecase）
- Jsons（for stub）
※repository関連の生成は行わない

## Using this project
```bash
go run main.go

Enter default path (e.g., /Users/xxx/java/todo-api/src/main/java/com/example/todoapi):
→デフォルトプロジェクトパス
Enter domain name (e.g., task):
→ドメイン名
Enter method name (e.g., create):
→メソッド名
```
