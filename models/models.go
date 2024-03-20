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

type Category struct {
	CategId   int    `json:"categId"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type Product struct {
	ProdId          int     `json:"prodId"`
	ProdTitle       string  `json:"prodTitle"`
	ProdDescription string  `json:"prodescription"`
	ProdCreatedAt   string  `json:"prodCreatedAt"`
	ProdUpdated     string  `json:"prodUpdated"`
	ProdPrice       float64 `json:"prodPrice,omitempty"`
	ProdStock       int     `json:"prodStock"`
	ProdCategId     int     `json:"prodCategId"`
	ProdPath        string  `json:"prodPath"`
	ProdSearch      string  `json:"search,omitempty"`
	ProdCategPath   string  `json:"categPath,omitempty"`
}

type ProductRes struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"`
}
