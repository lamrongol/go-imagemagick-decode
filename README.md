# Golang Image decode by ImageMagick(external software)
Golang `image.Decode()` raises error for many images and this problem is not solved for a long term. This module solves it by external software, ImageMagick. 

## Quickstart
First, you need to install ImageMagick(https://imagemagick.org/). Then
```
go get github.com/lamrongol/go-imagemagick-decode
```
Use `magick.Decode()` instead of `image.Decode()`

## Examples
`image.Decode()` fail cases are reported on following:
- https://github.com/golang/go/issues/2362
- https://github.com/golang/go/issues/10447
- https://github.com/golang/go/issues/45902

This module can decode for all of above cases. You can confirm by `image_test.go`.
You can also confirm by https://hub.docker.com/r/lamrongol/golang-imagemagick

---
In Japanese

#  GoでのImageMagick(外部ソフトウェア)による画像Decode
Goの`image.Decode()`は多くの画像でエラーを起こし、この問題は長年に渡って解決されていません。このモジュールは外部ソフトウェアImageMagickを使用してDecodeします。

## クイックスタート
まずImageMagick(https://imagemagick.org/)をインストールする必要があります。そして
```
go get github.com/lamrongol/go-imagemagick-decode
```
その後`image.Decode()`の代わりに`magick.Decode()`を使用します。

## 実際の例
`image.Decode()`が失敗するケースは公式レポジトリの以下のIssueで報告されています。
- https://github.com/golang/go/issues/2362
- https://github.com/golang/go/issues/10447
- https://github.com/golang/go/issues/45902

このモジュールは上のすべてのケースをDecodeできます。`image_test.go`で確認できます。
https://hub.docker.com/r/lamrongol/golang-imagemagick でも確認できます。
