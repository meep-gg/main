package scylla

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	. "meep.gg/log"
)

var Session *gocql.Session

func InitScylla(addr string) {
	cluster := gocql.NewCluster(addr)
	cluster.ProtoVersion = 4
	cluster.Keyspace = "meep"
	cluster.Consistency = gocql.Quorum
	cluster.ReconnectInterval = time.Second * 30

	// disable this for production
	cluster.DisableInitialHostLookup = true

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error connecting to ScyllaDB: %v", Error(err))
	}

	createKeyspaceQuery := "CREATE KEYSPACE IF NOT EXISTS meep WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 3};"
	err = Session.Query(createKeyspaceQuery).Exec()
	if err != nil {
		log.Fatalf("Error creating keyspace: %v", Error(err))
	}
}

func CreateIndex(region, endpoint, table, name string) {
	// Create index on accountId
	query := fmt.Sprintf(`CREATE INDEX IF NOT EXISTS 
		%s_idx ON %s_%s.%s (%s)`, name, region, endpoint, table, name)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("Error creating index: %v", Error(err))
	}
}

func CreateKeyspace(region, name string) {
	query := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s_%s WITH "+
		"replication = {'class': 'SimpleStrategy', 'replication_factor' : 3};", region, name)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("Error creating keyspace: %v", Error(err))
	}
}
