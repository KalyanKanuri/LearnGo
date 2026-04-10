package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learning/go/dbms/repo_pattern/models"
	"learning/go/dbms/repo_pattern/utils"

	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(ord *models.Order)
	GetCustomerOrders(name string) []models.Order
}

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

func GenOrderID() string {
	uuidStr := uuid.New()
	orderID := fmt.Sprintf("ORD-%x", uuidStr[:4])
	return orderID
}

func (or *OrderRepo) CreateOrder(ord *models.Order) string {
	fmt.Printf("[CreateOrder] Creating order: %+v by customer: %s\n", ord, ord.Customer.Name)
	txn, err := or.DB.Begin()
	repoutils.CheckErr(err)
	defer txn.Rollback()

	ordInsertSql := `
				INSERT INTO orders (order_id, description, price, customer_id)
				VALUES (?, ?, ?, ?)
			`
	ordStmnt, err := txn.Prepare(ordInsertSql)
	repoutils.CheckErr(err)
	defer ordStmnt.Close()

	ord.OrderID = GenOrderID()
	ordRes, err := ordStmnt.Exec(ord.OrderID, ord.Description, ord.Price, ord.Customer.CustomerID)
	repoutils.CheckErr(err)

	_, err = ordRes.LastInsertId()

	ordBS, err := json.MarshalIndent(ord, "", " ")
	fmt.Printf(
		"[CreateOrder] Order inserted | %+v\n", string(ordBS),
	)

	ordItemsSql := `
				INSERT INTO order_item (quantity, product_name, product_type, unit_price, order_id)
				VALUES (?, ?, ?, ?, ?)
			`
	ordItemStmnt, err := txn.Prepare(ordItemsSql)
	repoutils.CheckErr(err)
	defer ordItemStmnt.Close()

	var ordItemIDs []int64
	for _, item := range ord.Items {
		itemRes, err := ordItemStmnt.Exec(
			item.Quantity, item.ProductName, item.ProductType, item.UnitPrice, ord.OrderID,
		)
		repoutils.CheckErr(err)
		itemID, err := itemRes.LastInsertId()
		repoutils.CheckErr(err)

		fmt.Printf(
			"[CreateOrderItem] Row -> %+v\n", string(ordBS),
		)
		ordItemIDs = append(ordItemIDs, itemID)
	}

	fmt.Printf(
		"[CreateOrder] Order Created with total items=%d for customer=%s\n",
		len(ordItemIDs), ord.Customer.Name,
	)

	err = txn.Commit()
	repoutils.CheckErr(err)
	return ord.OrderID
}

func (or *OrderRepo) GetOrderDetails(ordID string) models.Order {
	fmt.Printf("[GetOrderDetails] Fetching Order Details for Order ID=%s\n", ordID)

	var order models.Order
	selectOrderSql := `
				    SELECT
						o.order_id,
						o.description,
						o.price,
						o.customer_id,
						oi.item_id,
						oi.quantity,
						oi.product_name,
						oi.product_type,
						oi.unit_price,
						o.created_at
					FROM orders o JOIN order_item oi
						ON o.order_id = oi.order_id
					WHERE o.order_id = ?
				`
	ordStmnt, err := or.DB.Prepare(selectOrderSql)
	repoutils.CheckErr(err)
	defer ordStmnt.Close()

	ordRows, err := ordStmnt.Query(ordID)
	repoutils.CheckErr(err)
	defer ordRows.Close()

	for ordRows.Next() {
		var item models.OrderItem
		err := ordRows.Scan(
			&order.OrderID,
			&order.Description,
			&order.Price,
			&order.Customer.CustomerID,
			&item.ItemID,
			&item.Quantity,
			&item.ProductName,
			&item.ProductType,
			&item.UnitPrice,
			&order.CreatedAt,
		)
		repoutils.CheckErr(err)
		order.Items = append(order.Items, item)
	}

	fmt.Printf(
		"[GetOrderDetails] Order found | ID=%s Description=%s Price=%v Customer=%d Items=%+v\n",
		order.OrderID,
		order.Description,
		order.Price,
		order.Customer.CustomerID,
		order.Items,
	)

	return order
}

func (or *OrderRepo) GetCustomerOrders(name string) []models.Order {
	var orders []models.Order
	orderMap := make(map[string]*models.Order)

	fmt.Printf("[GetCustomerOrders] Fetching orders for customer=%s\n", name)
	selectOrdersSql := `
					SELECT
						c.customer_id,
						c.customer_name,
						c.customer_address,
						o.order_id,
						o.description,
						o.price,
						oi.item_id,
						oi.quantity,
						oi.product_name,
						oi.product_type,
						oi.unit_price
					FROM orders o JOIN customer c
						ON o.customer_id = c.customer_id
					JOIN order_item oi
						ON oi.order_id = o.order_id
					WHERE c.customer_name = ?
				`
	stmnt, err := or.DB.Prepare(selectOrdersSql)
	repoutils.CheckErr(err)
	defer stmnt.Close()

	res, err := stmnt.Query(name)
	repoutils.CheckErr(err)
	defer res.Close()
	fmt.Println("[GetCustomerOrders] Orders query executed successfully")

	for res.Next() {
		// re-initialize order in loop to avoid order overwrite
		var (
			tmpOrder models.Order
			item     models.OrderItem
		)

		err := res.Scan(
			&tmpOrder.Customer.CustomerID,
			&tmpOrder.Customer.Name,
			&tmpOrder.Customer.Address,
			&tmpOrder.OrderID,
			&tmpOrder.Description,
			&tmpOrder.Price,
			&item.ItemID,
			&item.Quantity,
			&item.ProductName,
			&item.ProductType,
			&item.UnitPrice,
		)
		repoutils.CheckErr(err)

		order, exists := orderMap[tmpOrder.OrderID]
		if !exists {
			order = &tmpOrder
			orderMap[tmpOrder.OrderID] = order
		}
		order.Items = append(order.Items, item)

		ord, err := json.MarshalIndent(order, "", " ")
		repoutils.CheckErr(err)
		fmt.Printf(
			"[GetCustomerOrders] Row -> Order: %+v\n",
			string(ord),
		)
	}

	for _, ord := range orderMap {
		orders = append(orders, *ord)
	}
	fmt.Printf(
		"[GetCustomerOrders] Total orders fetched=%d for customer=%s\n",
		len(orders),
		name,
	)

	return orders
}
