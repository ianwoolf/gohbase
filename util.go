package hbase2go

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
