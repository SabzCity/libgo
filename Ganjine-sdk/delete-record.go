/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// DeleteRecord delete specific record by given ID in all cluster!
// We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
// It won't delete record history or indexes associate to it!
func DeleteRecord(c *ganjine.Cluster, req *gs.DeleteRecordReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.DeleteRecord(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return err
	}

	st.Service = &gs.DeleteRecordService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return err
	}
	return st.Err
}
