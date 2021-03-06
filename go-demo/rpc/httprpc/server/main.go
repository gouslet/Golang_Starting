/*
 * File: \rpc\httprpc\server\main.go                                           *
 * Project: go-demo                                                            *
 * Created At: Thursday, 2022/06/9 , 00:02:57                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:34:54                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	hellorpc "go_start/go-demo/rpc"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.RegisterName("HelloService", new(hellorpc.HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

		var conn io.ReadWriteCloser = struct {
			io.Writer

			io.ReadCloser
		}{

			ReadCloser: r.Body,

			Writer: w,
		}

		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	})

	http.ListenAndServe(":1234", nil)

}
