package database

import (
	"database/sql"
	"glassnode_challenge/model"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
 * DoGetHourlyGasFee Get data from bata base
 */
func DoGetHourlyGasFee() ([]model.Data, error) {

	sqlQuery :=
		`SELECT 
			date_trunc('hour', t.block_time) as block_time_hour, 
			sum(t.gas_used * t.gas_price) / 10 ^ 18 as eth
		FROM 
			public.transactions t
		WHERE 
			t.to != '0x0000000000000000000000000000000000000000'
			AND t.value > 0
			AND not exists (select 1 from public.contracts c WHERE t.to = c.address)
		GROUP BY block_time_hour
		ORDER BY block_time_hour DESC;`

	stmt, err := GetSQLConnection().Prepare(sqlQuery)

	defer closeStmt(stmt)

	datas := make([]model.Data, 0)

	if err != nil {
		log.Error("GetSQLConnection- DoGetHourlyGasFee ", err.Error())
		return datas, err
	}

	rows, err := stmt.Query()

	defer closeRows(rows)

	if err != nil {
		log.Error("Query- DoGetHourlyGasFee ", err.Error())
		return datas, err
	}

	for rows.Next() {

		var time time.Time
		var fees float64

		err = rows.Scan(&time, &fees)

		if err != nil {
			log.Error("Scan rows- DoGetHourlyGasFee ", err.Error())
			return datas, err
		}

		datas = append(datas, model.Data{
			Time: time.Unix(),
			Fees: fees,
		})
	}
	return datas, err
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
