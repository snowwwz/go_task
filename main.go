package main

import (
	"github.com/yukinooz/go-task/service/infrastructure"
	"github.com/yukinooz/go-task/service/interfaces"
)

// DAY１
// go mysql 環境構築
// DB table作成
// 一覧にrowデータが表時されるところまで
// Day2
// git管理
//
// BASIC
// 今日アプリケーション
// todo --lang spanish flagは--のこと
// task list 完了でないタスク全取得 出力；表
// task list -a 完了含めて取得　出力:表
// task add name priority deadline(int) 追加 入力待つことできる?
// task done id 完了
// task delete id 削除

// FUTURE
// migration
// https://github.com/gizak/termui グラフ 消化率/ リストなど

func main() {
	//conf logging

	handler := infrastructure.NewSQLHandler()
	interfaces.NewTaskRepository(handler)

}
