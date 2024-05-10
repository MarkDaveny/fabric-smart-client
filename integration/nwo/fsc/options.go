/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fsc

import (
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc/node"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver/sql/postgres"
)

// WithPostgresPersistence is a configuration with SQL vault persistence
func WithPostgresPersistence(config *postgres.PostgresConfig) node.Option {
	return func(o *node.Options) error {
		if config != nil {
			o.Put("fsc.persistence.sql", config.DataSource())
		}
		return nil
	}
}
