package grifts

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/suite/fix"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	_ = grift.Desc("seed", "Seeds a database")
	_ = grift.Add("seed", func(c *grift.Context) error {

		box := packr.New("Fixtures", "../fixtures")
		if err := fix.Init(box); err != nil {
			log.Fatal(err)
		}

		sc, err := fix.Find("basic sample data")
		if err != nil {
			log.Fatal(err)
		}

		db, err := pop.Connect("development")
		if err != nil {
			log.Fatal(err)
		}

		// This code is ripped from the suite package's Model.LoadFixture call
		// Except for the lastID thing; we have to reset the sequence to match the data we inserted, since some of our
		// fixture data uses specific IDs rather than the primary key sequence. The code does this for any table that
		// has a standard primary key id column.
		for _, table := range sc.Tables {
			lastID := 1
			for _, row := range table.Row {
				lastID++
				q := "insert into " + table.Name
				keys := []string{}
				skeys := []string{}
				for k := range row {
					keys = append(keys, k)
					skeys = append(skeys, ":"+k)
				}

				q = q + fmt.Sprintf(" (%s) values (%s)", strings.Join(keys, ","), strings.Join(skeys, ","))
				_, err = db.Store.NamedExec(q, row)
				if err != nil {
					log.Fatal(err)
				}
			}
			_, err := db.Store.Exec(fmt.Sprintf(`SELECT setval('%s_id_seq', %d, FALSE)`, table.Name, lastID))
			if err != nil {
				if !strings.Contains(err.Error(), "does not exist") {
					log.Fatal(err)
				}
			}
		}

		return nil
	})
})
