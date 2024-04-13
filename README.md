clone from this https://github.com/codebangkok/golang/tree/main/project/gokafka
add bruno collection in producer service
server = zookeeper kafka server
 need setup mysql server

## exmaple using kafka from console (need to install kafka source code)

>kafka-console-producer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent
{"ID":"abc123","AccountHolder":"Bond","AccountType":1,"Balance":1000}

>kafka-console-producer --bootstrap-server=localhost:9092 --topic=DepositFundEvent
{"ID":"abc123","Amount":1000}

>kafka-console-producer --bootstrap-server=localhost:9092 --topic=WithdrawFundEvent
{"ID":"abc123","Amount":600}

>kafka-console-producer --bootstrap-server=localhost:9092 --topic=CloseAccountEvent
{"ID":"abc123"}
