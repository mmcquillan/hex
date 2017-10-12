package hexplugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Handshake Config
func GetHandshakeConfig() plugin.HandshakeConfig {
	hc := plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "HexBotPlugin",
		MagicCookieValue: "HexBotPlugin",
	}
	return hc
}

// Arguments
type Arguments struct {
	Debug   bool
	Command string
	Config  map[string]string
}

// response args
type Response struct {
	Output  string
	Success bool
}

// Action - interface for the plugin
type Action interface {
	Perform(args Arguments) (resp Response)
}

type ActionRPC struct{ client *rpc.Client }

func (g *ActionRPC) Perform(args Arguments) (resp Response) {
	err := g.client.Call("Plugin.Action", args, &resp)
	if err != nil {
		// do something with return
		panic(err)
	}

	return resp
}

type ActionRPCServer struct {
	Impl Action
}

func (s *ActionRPCServer) Action(args Arguments, resp *Response) error {
	*resp = s.Impl.Perform(args)
	return nil
}

type HexPlugin struct {
	Impl Action
}

func (p *HexPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ActionRPCServer{Impl: p.Impl}, nil
}

func (HexPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ActionRPC{client: c}, nil
}
