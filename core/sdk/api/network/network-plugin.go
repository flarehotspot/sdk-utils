package network

type INetworkPlugin interface {
	OnInit() error
}
