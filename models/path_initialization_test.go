package models

import (
	"testing"
)

func TestDataBase_PathInitReady(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()

	proxyAddress:="0x306dca0180732AAF92be106F108baa0D254a62dc"
	anchorAddress:="0x9Ed4FcDF5DA6525113e7010c9894D7e918a63e8e"
	resourceAnchorPath:="judicial.geth.evidence"

	flag := dataBase.PathInitReady(proxyAddress,anchorAddress,resourceAnchorPath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id=", flag)
}
