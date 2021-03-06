// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

type zkServer struct {
	Host string
	Port string
}

type zkCluster struct {
	Nodes []zkServer
}
