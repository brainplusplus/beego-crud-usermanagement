package responses

type Response struct {
	Success    	bool
	Code  		string
	Data   		interface{}
	Message		string
}

type ResponseList struct {
	Success    		bool
	Code  			string
	Rows   			interface{}
	RecordsTotal	int
	RecordsFiltered	int
	Message			string
}
