package hbase2go

import (
	"fmt"
	"net"
	"strconv"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type HBConnect struct {
	Host   string
	Port   int
	Trans  *thrift.TSocket
	Client *THBaseServiceClient
}

func NewHbObj(host string, port int) HBConnect {
	return HBConnect{
		Host: host,
		Port: port,
	}
}

func (h *HBConnect) Close() {
	h.Trans.Close()
}
func (h *HBConnect) Connect() error {
	var err error
	h.Trans, err = thrift.NewTSocket(net.JoinHostPort(h.Host, strconv.Itoa(h.Port)))
	if err != nil {
		return fmt.Errorf("error resolving address:%s", err)

	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	h.Client = NewTHBaseServiceClientFactory(h.Trans, protocolFactory)
	return h.Trans.Open()
}

func (h *HBConnect) GetRow(table string, tget *TGet) (*TResult_, error) {
	return h.Client.Get([]byte(table), tget)
}

func GenTGet(rowKey, TableCF, Qual string) (tget *TGet) {
	tget = &TGet{
		Row:     []byte(rowKey),
		Columns: []*TColumn{},
	}
	if TableCF != "" {
		Tcol := &TColumn{
			Family: []byte(TableCF),
		}
		if Qual != "" {
			Tcol.Qualifier = []byte(Qual)
		}
		tget.Columns = append(tget.Columns, Tcol)
	}
	return
}
