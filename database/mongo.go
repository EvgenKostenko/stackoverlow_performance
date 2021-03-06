package database

import (
	"github.com/EvgenKostenko/stackoverlow_performance/config"
	"gopkg.in/mgo.v2"
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
)

type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	Col     *mgo.Collection
}

func init() {

	if mainSession == nil {

		var err error
		mainSession, err = mgo.Dial(config.Config.DB.Host)

		if err != nil {
			panic(err)
		}

		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(config.Config.DB.Host)

	}

}

func (m *MgoDb) Init() *mgo.Session {

	m.Session = mainSession.Copy()
	m.Db = m.Session.DB(config.Config.DB.Name)

	return m.Session
}

func (m *MgoDb) C(collection string) *mgo.Collection {
	m.Col = m.Session.DB(config.Config.DB.Name).C(collection)
	return m.Col
}

func (m *MgoDb) Close() bool {
	defer m.Session.Close()
	return true
}

func (m *MgoDb) DropoDb() {
	err := m.Session.DB(config.Config.DB.Name).DropDatabase()
	if err != nil {
		panic(err)
	}
}

func (m *MgoDb) RemoveAll(collection string) bool {
	m.Session.DB(config.Config.DB.Host).C(collection).RemoveAll(nil)

	m.Col = m.Session.DB(config.Config.DB.Name).C(collection)
	return true
}

func (m *MgoDb) Index(collection string, keys []string) bool {

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := m.Db.C(collection).EnsureIndex(index)
	if err != nil {
		panic(err)

		return false
	}

	return true
}

func (m *MgoDb) IsDup(err error) bool {

	if mgo.IsDup(err) {
		return true
	}

	return false
}
