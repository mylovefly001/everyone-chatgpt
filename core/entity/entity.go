package entity

type Config struct {
	Base struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"base"`
	Resty struct {
		TimeOut          int `json:"time_out"`
		RetryCount       int `json:"retry_count"`
		RetryWaitTime    int `json:"retry_wait_time"`
		RetryMaxWaitTime int `json:"retry_max_wait_time"`
	} `json:"resty"`
	Mysql struct {
		Master struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			User    string `json:"user"`
			Pass    string `json:"pass"`
			Db      string `json:"db"`
			Charset string `json:"charset"`
		} `json:"master"`
		Slave struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			User    string `json:"user"`
			Pass    string `json:"pass"`
			Db      string `json:"db"`
			Charset string `json:"charset"`
		} `json:"slave"`
	} `json:"mysql"`
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Db   int    `json:"db"`
		Auth string `json:"auth"`
	} `json:"redis"`
	Rabbit struct {
		Default struct {
			Host      string `json:"host"`
			Port      int    `json:"port"`
			User      string `json:"user"`
			Pass      string `json:"pass"`
			Vhost     string `json:"vhost"`
			Heartbeat int    `json:"heartbeat"`
		} `json:"default"`
	} `json:"rabbit"`
	Email struct {
		Host     string `json:"host"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Port     int    `json:"port"`
	} `json:"email"`
}
