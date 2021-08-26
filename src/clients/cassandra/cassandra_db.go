package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to Cassandra cluster
	tmp := gocql.NewCluster("cassandra")
	tmp.Keyspace = "oauth"
	tmp.Consistency = gocql.Quorum

	cluster = tmp
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
