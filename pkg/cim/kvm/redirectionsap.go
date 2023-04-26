/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

const CIM_KVMRedirectionSAP = "CIM_KVMRedirectionSAP"

type RedirectionSAP struct {
	base wsman.Base
}

// NewKVMRedirectionSAP returns a new instance of the KVMRedirectionSAP struct.
func NewKVMRedirectionSAP(wsmanMessageCreator *wsman.WSManMessageCreator) RedirectionSAP {
	return RedirectionSAP{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_KVMRedirectionSAP)),
	}
}
func (k RedirectionSAP) RequestStateChange(requestedState int) string {
	return k.base.RequestStateChange(actions.RequestStateChange(string(CIM_KVMRedirectionSAP)), requestedState)
}
func (b RedirectionSAP) Get() string {
	return b.base.Get(nil)
}

func (b RedirectionSAP) Enumerate() string {
	return b.base.Enumerate()
}
func (b RedirectionSAP) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}