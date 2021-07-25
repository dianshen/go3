package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)


type DnsRecord struct {
	DomainName string `db:"domain_name"`
	DigTime string	`db:"dig_time"`
	RecordList string	`db:"record_list"`
}

func GetDomain(record *DnsRecord) error{
	db, err := sqlx.Open("mysql", "test:xlk98Ala2@tcp(10.12.77.56:3306)/test")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//err = db.QueryRow("select domain_name from dns_record where dial_point_uuid = $1", dial_point_uuid).Scan(&karma)
	err = db.QueryRow("select domain_name, record_list, dig_time from dns_record " +
		"where id=100").Scan(
		&record.DomainName,  &record.RecordList, &record.DigTime)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func main() {
	var dnsRecord = DnsRecord{}
	err := GetDomain(&dnsRecord)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("domain_name is %s, dig_time is %s, record_list is %s",
		dnsRecord.DomainName, dnsRecord.DigTime, dnsRecord.RecordList)
}
