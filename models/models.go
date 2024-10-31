package models

type SecretRDSJson struct {
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Password            string `json:"password"`
	Port                int    `json:"port"`
	Username            string `json:"username"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

type Address struct {
	AddId         int    `json:"addId"`
	AddTitle      string `json:"addTitle"`
	AddName       string `json:"addName"`
	AddAddress    string `json:"addAddress"`
	AddCity       string `json:"addCity"`
	AddState      string `json:"addState"`
	AddPostalCode string `json:"addPostalCode"`
	AddPhone      string `json:"addPhone"`
}

type Category struct {
	CategId        int    `json:"categId"`
	CategName      string `json:"categName"`
	CategPath      string `json:"categPath"`
	CategTotalSold int    `json:"categTotalSold"`
}

type Currencies struct {
	COP            float64 `json:"cop"`
	EUR            float64 `json:"eur"`
	TimeLastUpdate string  `json:"timeLastUpdate"`
	USD            float64 `json:"usd"`
}

type Currency struct {
	BaseCurrency   string  `json:"base_currency"`
	CurrencyRate   float64 `json:"rate"`
	LastUpdated    string  `json:"last_updated"`
	TargetCurrency string  `json:"target_currency"`
}

type ListUsers struct {
	TotalItems int    `json:"totalItems"`
	Data       []User `json:"data"`
}

type Orders struct {
	Order_Id       int     `json:"orderId"`
	Order_UserUUID string  `json:"orderUserUUID"`
	Order_AddID    int     `json:"orderAddID"`
	Order_Date     string  `json:"orderDate"`
	Order_Total    float64 `json:"orderTotal"`
	OrderDetails   []OrdersDetails
}

type OrdersDetails struct {
	OD_Currency             string  `json:"odCurrency"`
	OD_Currency_Last_Symbol string  `json:"odCurrencyLastSymbol"`
	OD_Currency_Symbol      string  `json:"odCurrencySymbol"`
	OD_Id                   int     `json:"odId"`
	OD_OrderId              int     `json:"odOrderId"`
	OD_ProdId               int     `json:"odProdId"`
	OD_Quantity             int     `json:"odQuantity"`
	OD_Price                float64 `json:"odPrice"`
}

type Product struct {
	ProdId          int     `json:"prodId"`
	ProdCategId     int     `json:"prodCategId"`
	ProdCategPath   string  `json:"categPath,omitempty"`
	ProdCreatedAt   string  `json:"prodCreatedAt"`
	ProdDescription string  `json:"prodDescription"`
	ProdDiscount    float64 `json:"prodDiscount,omitempty"`
	ProdPath        string  `json:"prodPath"`
	ProdPrice       float64 `json:"prodPrice,omitempty"`
	ProdSearch      string  `json:"search,omitempty"`
	ProdStock       int     `json:"prodStock"`
	ProdTitle       string  `json:"prodTitle"`
	ProdUpdated     string  `json:"prodUpdated"`
}

type ProductRes struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"`
}

type User struct {
	UserUUID      string `json:"userUUID"`
	UserEmail     string `json:"userEmail"`
	UserFirstName string `json:"userFirstName"`
	UserLastName  string `json:"userLastName"`
	UserStatus    int    `json:"userStatus"`
	UserDateAdd   string `json:"userDateAdd"`
	UserDateUpg   string `json:"userDateUpg"`
}

type UserRole struct {
	UserUUID string `json:"userUUID"`
	UserRole string `json:"userRole"`
}
