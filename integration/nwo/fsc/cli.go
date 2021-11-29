/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fsc

import (
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc/commands"
	"github.com/hyperledger-labs/fabric-smart-client/pkg/api"
)

type fscCLIViewClient struct {
	timeout time.Duration
	p       *Platform
	CMD     commands.View
}

func (f *fscCLIViewClient) CallView(fid string, in []byte) (interface{}, error) {
	f.CMD.Input = string(in)
	f.CMD.Function = fid

	sess, err := f.p.FSCCLI(f.CMD)
	if err != nil {
		return nil, err
	}
	gomega.Eventually(sess, f.timeout).Should(gexec.Exit(0))

	return string(sess.Out.Contents()), nil
}

func (f *fscCLIViewClient) IsTxFinal(txid string, opts ...api.ServiceOption) error {
	panic("not implemented yet")
}
