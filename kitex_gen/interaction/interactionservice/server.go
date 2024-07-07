// Code generated by Kitex v0.6.2. DO NOT EDIT.
package interactionservice

import (
	server "github.com/cloudwego/kitex/server"
	interaction "github.com/lukanzx/DouVo/kitex_gen/interaction"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler interaction.InteractionService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
