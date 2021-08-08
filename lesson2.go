package main

import (
	"common"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Domain interface {
	GetDomainInfo()
}

type DnsRecord struct {
	DomainName string `db:"domain_name"`
	DigTime string	`db:"dig_time"`
	RecordList string	`db:"record_list"`
}

func (record *DnsRecord) GetDomainInfo() error{
	db, err := sqlx.Open("mysql", common.GetMySQLDataSourceName())
	if err != nil {
		fmt.Println(err)
	}
	row := db.QueryRow("select domain_name, record_list, dig_time from dns_record where 1=0")
	err = row.Scan(&record.DomainName,  &record.RecordList, &record.DigTime)
	return err
}

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds)
	dnsRecord := &DnsRecord{}
	err := dnsRecord.GetDomainInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("domain_name is %s, dig_time is %s, record_list is %s",
		dnsRecord.DomainName, dnsRecord.DigTime, dnsRecord.RecordList)
}
