package pg

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-pg/pg/v9"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

type Tbl struct {
	tableName struct{}  `sql:"tbl_test"`
	Id        int       `sql:"id"`
	Info      string    `sql:"info"`
	CTime     time.Time `sql:"c_time"`
	Info1     string    `sql:"info1"`
	Info2     string    `sql:"info2"`
	Info3     string    `sql:"info3"`
	Info4     string    `sql:"info4"`
	Info5     string    `sql:"info5"`
	Info6     string    `sql:"info6"`
	Info7     string    `sql:"info7"`
}

type TblCurriculum struct {
	tableName    struct{} `sql:"tbl_curriculum_test"`
	Id           int      `sql:"id"`
	TblId        int      `sql:"tbl_id"`
	CurriculumId int      `sql:"curriculum_id"`
}

func Test10(t *testing.T) {
	db := pg.Connect(&pg.Options{
		Addr:     ":5433",
		User:     "postgres",
		Password: "201104",
		Database: "postgres",
	})
	db.AddQueryHook(dbLogger{})
	start := time.Now()
	// 7140
	//	_, err := db.Exec(`
	//	select count( distinct tbl_test.info)
	//     from tbl_test,
	//          tbl_curriculum_test
	//     where tbl_test.id = tbl_curriculum_test.tbl_id
	//       and curriculum_id in (2229)
	//`)
	//	if err != nil {
	//		t.Fatal(err)
	//	}

		// 3617
	lastId := 0
	tblCurriculumList := make([]int, 0)
	for {
		tblCurriculumList = tblCurriculumList[0:0]
		query := db.Model(&TblCurriculum{}).
			ColumnExpr("distinct tbl_id")

		if lastId != 0 {
			query.Where("tbl_id > ? ", lastId)
		}

		err := query.Order("tbl_id asc").Limit(5000).Select(&tblCurriculumList)
		if err != nil {
			t.Fatal(err)
		}

		if len(tblCurriculumList) == 0 {
			break
		}
		lastId = tblCurriculumList[len(tblCurriculumList)-1]

		_, err = db.Model(&Tbl{}).
			Where("id in (?)", pg.In(tblCurriculumList)).Count()
		if err != nil {
			t.Fatal(err)
		}
	}

	end := time.Now()
	fmt.Print(end.Sub(start).Milliseconds())

}
