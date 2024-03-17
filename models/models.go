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
