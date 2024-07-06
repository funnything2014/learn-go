package coding

/**
1. 代码规范
	1. 命名规范
		包名：
	尽量和目录保持一致
	尽量采取有意义的包名，简短
	不要和标准库名冲突
	采用全部小写
		文件名
	user_name.go采用蛇形命名法
		变量名
	蛇形：python、php
	驼峰：java、c、go
	userName
	UserName
	un string
	有一些专有名词，APIVersion、URLVersion
	bool类型，has, is, can, allow开头
		结构体命名
	驼峰,User
		接口命名
	基本上和结构体差不多
	接口以er结尾
	type Writer interface
	type IRead interface
		常量命名
	全部大写，多个单词，蛇形命名法 APP_VERSION
2. 注释规范
	两种注释
	// 单行注释
	/** 大段注释
	变量后面加注释
	包注释
	接口注释
    函数注释
    代码逻辑的注释
2. import规范
	go 自带的包
	第三方的包
	自己内部的包
	os
	io

	gorm.io/gorm

	src.imooc.com/bobby/A
	src.imooc.com/bobby/B
*/
