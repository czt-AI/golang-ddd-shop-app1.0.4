package model

// UserStatistics represents statistics about users.
type UserStatistics struct {
	UserCount   int
	ActiveUsers int
}

// ProductStatistics represents statistics about products.
type ProductStatistics struct {
	ProductCount int
	SalesVolume  float64
}

// OrderStatistics represents statistics about orders.
type OrderStatistics struct {
	OrderCount int
	TotalSales float64
}

// PaymentStatistics represents statistics about payments.
type PaymentStatistics struct {
	PaymentCount int
	TotalAmount  float64
}

// LogisticsStatistics represents statistics about logistics.
type LogisticsStatistics struct {
	ShippedOrders int
	DeliveredOrders int
}

// NewUserStatistics creates a new UserStatistics instance.
func NewUserStatistics(userCount, activeUsers int) *UserStatistics {
	return &UserStatistics{
		UserCount:   userCount,
		ActiveUsers: activeUsers,
	}
}

// NewProductStatistics creates a new ProductStatistics instance.
func NewProductStatistics(productCount, salesVolume int) *ProductStatistics {
	return &ProductStatistics{
		ProductCount: productCount,
		SalesVolume:  salesVolume,
	}

}

// NewOrderStatistics creates a new OrderStatistics instance.
func NewOrderStatistics(orderCount, totalSales int) *OrderStatistics {
	return &OrderStatistics{
		OrderCount: orderCount,
		TotalSales: totalSales,
	}
}

// NewPaymentStatistics creates a new PaymentStatistics instance.
func NewPaymentStatistics(paymentCount, totalAmount int) *PaymentStatistics {
	return &PaymentStatistics{
		PaymentCount: paymentCount,
		TotalAmount:  totalAmount,
	}
}

// NewLogisticsStatistics creates a new LogisticsStatistics instance.
func NewLogisticsStatistics(shippedOrders, deliveredOrders int) *LogisticsStatistics {
	return &LogisticsStatistics{
		ShippedOrders: shippedOrders,
		DeliveredOrders: deliveredOrders,
	}
}