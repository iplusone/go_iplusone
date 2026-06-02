# 54 Worker Pool Basic: goroutine を無限に増やさず、数を決めて回す

このハンズオンでは Worker Pool を扱います。Go は goroutine が軽いですが、だからといって無制限に増やしてよいわけではありません。

ここでは `jobs` channel に仕事を流し、固定数の worker goroutine がそれを取り出して処理します。`close(jobs)` で `range` が終わるのも重要なポイントです。

Worker Pool の意味は、「仕事の数」と「同時に動く数」を分けることです。ジョブが10件あっても、worker が3本なら並列度は3です。

この回は、並行処理を「増やす」から「制御する」へ進む入口としてとても大事です。
