package cmd

import (
	"time"
	"context"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/cobra"
	"github.com/asdfsx/zhihu-golang-web/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "web server's main process",
	Long:  "read the config file, read logs according to the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := readConfig()
		if err != nil{
			return err
		}
		err = serveFunc(cmd, args)
		if err == nil{
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
}

func logrotate(ctx context.Context){
	//判断是否应该rotate日志
	ticker := time.Tick(5 * time.Minute)
	for {
		select {
		case <- ctx.Done():
			return
		case <-ticker:
			now := time.Now()
			if now.Day() > rotateTime.Day() {
				if err := logger.Rotate(); err != nil {
					jww.ERROR.Printf("Failed to rotate the logger %v\n", err)
				}
				rotateTime = now
			}
		}
	}
}

func serveFunc(cmd *cobra.Command, args []string) error {
	jww.INFO.Printf("server name: %v\n", config.Server.Name)

	ctx := ContextWithSignal(context.Background())

	go logrotate(ctx)

	svr, err := server.NewServer(&config)
	if err != nil{
		return err
	}
	defer svr.Close()

	svr.StartWithContext(ctx)

	svr.Wait()

	jww.INFO.Println("Shutting down")
	return nil
}
