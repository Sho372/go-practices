# Go Practices

| No  | Module                              | Descriptoin                                                | Note                                                              |
| --- | ----------------------------------- | ---------------------------------------------------------- | ----------------------------------------------------------------- |
| 1   | cording-rule                        | Goのコーディングルール。業務で実装する時に毎回悩む箇所をルール化しておく。ルール自体はNotionに記載している。 |                                                                   |
| 2   | echo-practice                       | echoを使ったAPIサーバーの練習。とりあえずエンドポイント追加して動作確認できるので、色々試せる。        | https://github.com/labstack/echo </br> https://echo.labstack.com/ |
| 3   | database-sql-practice               | 標準のsqlパッケージを使ったMySQLのDB操作の練習                               | https://pkg.go.dev/database/sql                                   |
| 4   | concurrency-patterns                | go routine. Learning Goに載っているパターンを模写。                      | https://www.oreilly.com/library/view/learning-go/9781492077206/   |
| 5   | goroutines                          | go routine. A Tour of GoのGo routineのセクションの模写。              | https://go.dev/tour/list                                          |
| 6   | visualizing-concurreucy             | go routine. Go confのセッションの内容の模写                            | https://youtu.be/KyuFeiG3Y60?si=MNCnbXNCfiUWInvq                  |
| 7   | validator-practice                  | validatorの練習                                               | https://github.com/go-playground/validator                        |
| 8   | http-client-error-handling-practice | シンプルなapi client                                            |                                                                   |
| 9   | unexported-exported                 | パッケージのexport/unexportの確認                                   |                                                                   |
| 10  | rabbitmq-tutorials                  | Rabbit MQの公式チュートリアルの模写                                     | https://www.rabbitmq.com/tutorials/tutorial-one-go.html           |

# Tips

## go.work
複数のモジュール（go.mod）が入っている。モジュールを追加するごとにgo.workに追加

## パッケージの作成
パッケージ作る時は、ディレクトリも必要

```
Goでは、パッケージを作成するためにディレクトリを作成する必要があります。Goの標準的なプラクティスでは、パッケージごとに個別のディレクトリを使用し、そのディレクトリ名とパッケージ名を一致させることが推奨されています。

例えば、`person`パッケージを作成する場合、`person`という名前のディレクトリを作成し、その中に`person.go`という名前のファイルを作成することになります。同様に、`book`パッケージも`book`という名前のディレクトリを作成し、その中に`book.go`という名前のファイルを作成します。

この構造を使うことで、コードを整理し、パッケージを明確に区別しやすくなります。また、パッケージを利用する他のコードからも、明示的なインポートパスでアクセスできるようになります。
```


