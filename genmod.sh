#!/bin/bash

# 引数が与えられていない場合はエラーメッセージを表示して終了
if [ "$#" -ne 1 ]; then
    echo "使用法: $0 <ディレクトリ名>"
    exit 1
fi

# 引数で指定されたディレクトリ名を変数に代入
DIR_NAME=$1

# ディレクトリを作成
mkdir -p "$DIR_NAME"

# 作成したディレクトリに移動
cd "$DIR_NAME"

# 現在のディレクトリパスを取得し、'github.com'までのパスを抽出
CURRENT_PATH=$(pwd)
GITHUB_PATH=$(echo $CURRENT_PATH | sed -n 's/.*\(github\.com.*\)/\1/p')

# go mod init コマンドを実行
go mod init $GITHUB_PATH/$DIR_NAME

go work use .
