package sql

import (
	"database/sql"
	"errors"
	"github.com/developerdong/review/conf"
	"github.com/developerdong/review/fgt"
	_ "modernc.org/sqlite"
	"net/url"
	"time"
)

const DriverName = "sqlite"

const tableUrl = `CREATE TABLE IF NOT EXISTS url
(
	id	INTEGER	NOT NULL	PRIMARY KEY	AUTOINCREMENT,
	url	TEXT	NOT NULL	UNIQUE
);
`

const tableRecord = `CREATE TABLE IF NOT EXISTS record
(
    url_id	INTEGER	NOT NULL,
    time	INTEGER	NOT NULL
);
`

type Sqlite struct {
	db *sql.DB
}

// connect opens the connection to the database set by the environment variables,
// and creates the needed tables if not exist.
func (s *Sqlite) connect() error {
	if driverName := conf.GetEnv(conf.DriverName); driverName != DriverName {
		return errors.New("the database should be sqlite")
	} else if db, err := sql.Open(DriverName, conf.GetEnv(conf.DataSourceName)); err != nil {
		return err
	} else if err := db.Ping(); err != nil {
		return err
	} else if _, err := db.Exec(tableUrl); err != nil {
		return err
	} else if _, err := db.Exec(tableRecord); err != nil {
		return err
	} else {
		s.db = db
		return nil
	}
}

func (s *Sqlite) Add(u *url.URL) error {
	if err := s.connect(); err != nil {
		return err
	} else {
		var urlId uint32
		if _, err := s.db.Exec(
			"INSERT OR IGNORE INTO url (url) VALUES (?);",
			u.String(),
		); err != nil {
			return err
		} else if err := s.db.QueryRow(
			"SELECT id FROM url WHERE url=?;",
			u.String(),
		).Scan(&urlId); err != nil {
			return err
		} else if _, err := s.db.Exec(
			"INSERT INTO record (url_id, time) VALUES (?, ?);",
			urlId,
			time.Now().Unix(),
		); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *Sqlite) Review() (*url.URL, error) {
	if urlRows, err := s.db.Query("SELECT id, url FROM url;"); err != nil {
		return nil, err
	} else {
		now := time.Now()
		var minRetrievabilityUrl string
		var minRetrievability float64
		for urlRows.Next() {
			var id int64
			var u string
			if err := urlRows.Scan(&id, &u); err != nil {
				return nil, err
			} else if recordRows, err := s.db.Query("SELECT time FROM record WHERE url_id=? ORDER BY time;", id); err != nil {
				return nil, err
			} else {
				points := make([]time.Time, 0)
				for recordRows.Next() {
					var seconds int64
					if err := recordRows.Scan(&seconds); err != nil {
						return nil, err
					} else {
						points = append(points, time.Unix(seconds, 0))
					}
				}
				if recordRows.Err() != nil {
					return nil, recordRows.Err()
				} else if retrievability := fgt.GetRetrievability(points, now); minRetrievabilityUrl == "" || retrievability < minRetrievability {
					minRetrievabilityUrl = u
					minRetrievability = retrievability
				}
			}
		}
		if urlRows.Err() != nil {
			return nil, urlRows.Err()
		} else {
			return url.Parse(minRetrievabilityUrl)
		}
	}
}
