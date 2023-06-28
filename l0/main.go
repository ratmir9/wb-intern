package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "wb"
	dbPassword = "test123"
	dbName     = "l0"
)

type Data struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int64  `json:"amount"`
		PaynentDt    int64  `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int64  `json:"delivery_cost"`
		GoodsTotal   int64  `json:"goods_total"`
		CustomFee    int64  `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int64  `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int64  `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int64  `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int32  `json:"total_price"`
		NmId        int64  `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int32  `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int32     `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

var dataCache = make(map[string]Data, 10)

func getDataById(c *gin.Context) {
	dataId := c.Param("id")
	result, ok := dataCache[dataId]
	if ok {
		resultJSON, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "data-detail.html", gin.H{
			"data": string(resultJSON),
		})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "message not found"})
}

func ConnectDb() (*sql.DB, error) {
	dbURI := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	fmt.Println(dbURI)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	sc, err := stan.Connect("test-cluster", "cluster-123")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sc.Close()

	sub, _ := sc.Subscribe("wb", func(m *stan.Msg) {
		var data Data
		err = json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Fatal(err)
		}
		curId := data.OrderUid
		_, err := db.Exec("INSERT INTO "+
			"orders (order_uid, track_number, entry, locale,"+
			"internal_signature, customer_id, delivery_service,"+
			"shardkey, sm_id, date_created, oof_shard)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			&data.OrderUid, &data.TrackNumber, &data.Entry, &data.Locale,
			&data.InternalSignature, &data.CustomerId, &data.DeliveryService,
			&data.Shardkey, &data.SmId, &data.DateCreated, &data.OofShard)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO "+
			"payments (order_uid, transaction, request_id, currency, provider, "+
			"amount, payment_dt, bank, goods_total, custom_fee, delivery_cost) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			&data.OrderUid, &data.Payment.Transaction, &data.Payment.RequestId,
			&data.Payment.Currency, &data.Payment.Provider, &data.Payment.Amount,
			&data.Payment.PaynentDt, &data.Payment.Bank, &data.Payment.GoodsTotal,
			&data.Payment.CustomFee, &data.Payment.DeliveryCost)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO "+
			"deliveries (order_uid, name, phone, zip, city, "+
			"address, region, email) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			&data.OrderUid, &data.Delivery.Name, &data.Delivery.Phone,
			&data.Delivery.Zip, &data.Delivery.City, &data.Delivery.Address,
			&data.Delivery.Region, &data.Delivery.Email)
		if err != nil {
			log.Fatal(err)
		}
		dataItems := data.Items
		fmt.Println(dataItems)
		for _, item := range dataItems {
			_, err = db.Exec("INSERT INTO "+
				"items (order_uid, chrt_id, track_number, price,"+
				"rid, name, sale, size, total_price, nm_id, brand, status) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
				&data.OrderUid, item.ChrtId, item.TrackNumber,
				item.Price, item.Rid, item.Name, item.Sale,
				item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
			if err != nil {
				log.Fatal(err)
			}
		}
		dataCache[curId] = data
	})
	defer sub.Unsubscribe()
	time.Sleep(time.Second * 25)

	server := gin.Default()
	server.Static("/assets", "./assets")
	server.LoadHTMLGlob("templates/*.html")
	server.GET("/data/:id", getDataById)
	server.Run(":8000")

}
