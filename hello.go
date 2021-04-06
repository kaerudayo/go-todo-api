
package main

import (
  "fmt"
  "errors"
)

type (
  UserId        uint64
  UserSessionId uint64
)

const (
  uid  UserId        = 1
  usid UserSessionId = 1
)

func main() {
  // 変数定義
  var (
    name string = ""
    age int = 1
  )
  name,age     = "kaoru", 24

  // 変数定義v2。他言語同様代入不可
  const (
    namev2 string = ""
    agev2  int    = 1
  )

  // namev2,agev2 = "kaoru", 24

  // セミコロンは基本書かなくていいけど、人によっては1行でさっと書きたいときは使いそう
  name = "kaoruv2"; age = 1;

  fmt.Println(name)
  fmt.Println(age)

  // bool				真偽値(true or false)
  // int8/int16/int32/int64     nビット整数。scalaでいうところのshort,int,longみたいな感じ
  // uint8/uint16/uint32/uint64 unsigned
  // float32/float64            nビット浮動小数点数 floatとdouble的な
  // complex64/complex128       nビット虚数 // なんじゃこれ
  // byte                       1バイトデータ(uint8と同義)
  // rune                       1文字(int32と同義) scalaでいうところのchar
  // uint                       uint32 または uint64 uintNの省略かな
  // int                        int32 または int64
  // uintptr                    ポインタを表現するのに十分な非負整数 //この辺は要勉強
  // string                     文字列

  // fmt.Println(uid == usid) // この段階でerrorになるのか。compileすごい
  // typeCheck(usid)          // これも怒られる

  // 型変換
  // fmt.Println(int16(1111111111111111111)) // これも怒られる

  /**
   * Map
   */
  map1 := map[string]int{
    "kaoru" : 100,  // 改行する場合はカンマ必須
    "hoge"  : 200,
  }

  fmt.Println(map1["kaor"])  // これはerrorにならないみたい。アプリケーションでのバグを逃しそう
  v, isDefined := map1["kaor"] // vの部分は_使用しない値の場合は_にした方が可読性が上がりそう
  if isDefined {
    fmt.Println(v)
  } else {
    fmt.Println("ないよ")
  }

  if v, hoge := map1["kaoru"]; hoge { // こんな風にもかける。これは人によりそう
    fmt.Println(v)
  } else {
    fmt.Println("ないよ")
  }

  /**
   * ループ
   */
   for a,b := range map1 {
     fmt.Printf("%s=%d\n",a,b)
   }

   /**
    * 例外処理
    */
    if err := doError(); err != nil {
      fmt.Println("err", err)
    } else {
      fmt.Println("終わりだよ～") // ここにはこない
    }

    fmt.Println(multiRes(1,2))

    for i := 0; i < 10; i ++ {
      fmt.Println(i)
    }

    lang := "Go"

    if lang == "Go" {
      fmt.Println(lang)
    }

    // ポインタについて。この記事が分かりやすかった https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
    hoge := "太郎"
    fmt.Printf("hoge :%v\n", hoge)

    hogePoint := &hoge

    // namePointは、&hogeが格納されているだけなので、stringへのポインタである*string型の値が格納されている。
    fmt.Printf("namePoint :%v\n", hogePoint)

    // namePointが指している変数は、"*namePoint"という感じで、"*"をつけて表す。
    fmt.Printf("namePoint :%v\n", *hogePoint)

    *hogePoint = "二郎"

    // *namePointに値を代入することもできる。
    fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", *hogePoint)

    // 再代入したところで、namePointに格納されている*string型の値(アドレス)自体は、変わらない
    fmt.Printf("*namePointに二郎を代入後のnamePoint :%v\n", hogePoint)

    // stringへのポインタである*string型の値(hogeに格納されている値)を書き換えたので、hogeの値も変更される。
    fmt.Printf("*namePointに二郎を代入後のhoge :%v\n", hoge)
}

func doError() error { return errors.New("これはerrorだ") }

func multiRes(a,b int) (int, int) {
  return a,b
}

func typeCheck(uid UserId) { // voidは書かなくても良さそう
  fmt.Println(uid)
}

