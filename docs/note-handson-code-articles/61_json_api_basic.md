# 61 JSON API Basic: 構造体とJSONの間を行き来する

このハンズオンでは `encoding/json` を使って、構造体と JSON の変換を確認します。Go で API を書くなら避けて通れない基本です。

最初の `json.Marshal` は、構造体を JSON 文字列へ変える流れです。`omitempty` を使った例もあるので、空値の扱いがどう変わるかも見えます。

次の `json.Unmarshal` では、JSON を構造体へ戻しています。ここで、JSON と Go の型がどう対応するかがかなり直感的につかめます。

最後の `MarshalIndent` は、見やすい JSON を作るためのものです。API の中身を観察するときにも便利です。

この回は、HTTP と組み合わさる直前の JSON 基礎としてとても重要です。
