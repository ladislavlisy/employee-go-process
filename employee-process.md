### Start Employee Process (POST)

You can create a new employee process with this action. It takes information about employee and payroll period and will set up a new process. The process will reset all period data for employee and create new data from input.

+ Request (application/json)

```json
	{
		"period" : 201701,
		"employee" : 
		{
			"first-name" : "alfred",
			"last-name" : "bob",
			"personal-number" : "A01",
			"contract-number" : "1",
			"contract-code" : "1",
			"contract-start-date" : "2017-01-01",
			"contract-end-date" : "",
			"department-code" : "D01",
			"cost-center-code" : "CC01",
			"full-name" : "alfred bob",
			"contract-name" : "Employment",
			"department-name" : "Main",
			"cost-center-name" : "Development"
		}
	}
```

+ Response 201 (application/json)
	+ Headers
	
		Location: /employee-process/5a003b78-409e-b456-a6f0dcee05bd
	
	+ Body
		
```json
	{
		"id" : "5a003b78-409e-b456-a6f0dcee05bd",
		"started_at" : "2017-09-01T08:40:51.620Z",
		"period" : 201701,
		"employee" : 
		{
			"first-name" : "alfred",
			"last-name" : "bob",
			"personal-number" : "A01",
			"contract-number" : "1",
			"contract-code" : "1",
			"contract-start-date" : "2017-01-01",
			"contract-end-date" : "",
			"department-code" : "D01",
			"cost-center-code" : "CC01",
			"full-name" : "alfred bob",
			"contract-name" : "Employment",
			"department-name" : "Main",
			"cost-center-name" : "Development"
		}
	}
```

