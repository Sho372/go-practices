# echo

- JSONの空配列はnilではなく、len:0, cap:0の配列にマップされる
- クエリパラメーターは`query`タグ、リクエストボディはGo標準の`json`タグを使う
- テストで擬似的なリクエストを用意する時、[クエリパラメーター](https://echo.labstack.com/docs/testing#setting-query-params)と[リクエストボディ](https://echo.labstack.com/docs/testing#using-form-payload)は文字列で直接書くより、set使った方がいい。最後にURLエンコードを忘れないようにする。
