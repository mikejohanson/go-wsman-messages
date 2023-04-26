/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_IEEE8021xSettings = "CIM_IEEE8021xSettings"

type Settings struct {
	base wsman.Base
}

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettings(wsmanMessageCreator *wsman.WSManMessageCreator) Settings {
	return Settings{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_IEEE8021xSettings)),
	}
}
func (b Settings) Get() string {
	return b.base.Get(nil)
}

func (b Settings) Enumerate() string {
	return b.base.Enumerate()
}
func (b Settings) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}