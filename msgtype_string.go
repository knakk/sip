// Code generated by "stringer -type msgType"; DO NOT EDIT

package sip

import "fmt"

const _msgType_name = "MsgUnknownMsgReqPatronStatusMsgReqCheckoutMsgReqCheckinMsgReqBlockPatronMsgReqStatusMsgReqResendMsgReqLoginMsgReqPatronInformationMsgReqEndPatronSessionMsgReqFeePaidMsgReqItemInformationMsgReqItemStatusUpdateMsgReqPatronEnableMsgReqHoldMsgReqRenewMsgReqRenewAllMsgRespPatronStatusMsgRespCheckoutMsgRespCheckinMsgRespStatusMsgRespLoginMsgRespPatronInformationMsgRespEndPatronSessionMsgRespFeePaidMsgRespItemInformationMsgRespItemStatusUpdateMsgRespPatronEnableMsgRespHoldMsgRespRenewMsgRespRenewAll"

var _msgType_index = [...]uint16{0, 10, 28, 42, 55, 72, 84, 96, 107, 130, 152, 165, 186, 208, 226, 236, 247, 261, 280, 295, 309, 322, 334, 358, 381, 395, 417, 440, 459, 470, 482, 497}

func (i msgType) String() string {
	if i < 0 || i >= msgType(len(_msgType_index)-1) {
		return fmt.Sprintf("msgType(%d)", i)
	}
	return _msgType_name[_msgType_index[i]:_msgType_index[i+1]]
}
