package main

import (
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	"fmt"
)

func main() {
	//初始配置文件
	settings.Init()
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	defer mysql.Close() // 程序退出关闭数据库连接

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
	//s := &http.Server{
	//	Addr:           setting.Conf.Add + ":" + setting.Conf.Port,
	//	Handler:        r,
	//	ReadTimeout:    setting.Conf.ReadTimeout,
	//	WriteTimeout:   setting.Conf.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//if err := s.ListenAndServe(); err != nil {
	//	log.Fatalf("run app failed: %s", err)
	//}

}
