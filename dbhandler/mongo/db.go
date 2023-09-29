package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"project/github/comics/dbhandler/config"
	"sync"
	"time"
)

// Mongo holds the mongo session
type Mongo struct {
	lock       sync.RWMutex
	enabled    bool
	connection string
	dbName     string
	client     *mongo.Client
}

func (m *Mongo) Close() error {
	if m.getClient() != nil {
		if err := m.getClient().Disconnect(context.TODO()); err != nil {
			log.Println(fmt.Sprintf("Unable to close mongo db (%s) connection", m.dbName))
		}
		m.setClient(nil)
	}
	return nil
}

func (m *Mongo) connect() error {
	timeOut := 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	opts := options.Client().ApplyURI(m.connection)

	maxConn := 100
	maxIdleTimeout := 60 * 5 * 1000
	minConn := 10

	opts = opts.SetMaxPoolSize((uint64)(maxConn))
	duration := time.Duration(maxIdleTimeout) * time.Millisecond
	opts = opts.SetMaxConnIdleTime(duration)
	opts = opts.SetMinPoolSize((uint64)(minConn))
	client, err := mongo.NewClient(opts)

	if err != nil {
		return err
	}

	if err := client.Connect(ctx); err != nil {
		_ = client.Disconnect(ctx)
		return err
	}

	//if err := client.Ping(ctx, nil); err != nil {
	//	_ = client.Disconnect(ctx)
	//	return err
	//}

	m.setClient(client)
	return nil
}

func (m *Mongo) setClient(c *mongo.Client) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.client = c
}

func (m *Mongo) getClient() *mongo.Client {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.client
}

func Init(c config.DBAppConfig) (*Mongo, error) {
	m := &Mongo{
		lock:       sync.RWMutex{},
		enabled:    false,
		connection: c.DbURL,
		dbName:     c.DBName,
	}

	err := m.connect()
	if err != nil {
		return nil, err
	}

	return m, err
}
