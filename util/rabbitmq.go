package util

import "github.com/streadway/amqp"

func SendDefaultTask(conn *amqp.Connection,queueName string,bodyData interface{},fn func(interface{},[]interface{}) interface{},fnParams ...interface{}) error{
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	bodyBytes,err := GetBytes(bodyData);
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			//Body:         []byte(body),
			Body:         bodyBytes,
		})
	if fn != nil {
		fn(bodyData,fnParams)
	}
	FailOnError(err, "Failed to publish a message in "+queueName)
	log.Printf(" [x] Sent %s", bodyData)
	log.Printf(" [x] Sent 2 %s", bodyBytes)

	return nil
}

func ReceiveDefaultTask(conn *amqp.Connection,queueName string,fn func([]byte,[]interface{}) interface{},fnParams ...interface{}) error{
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			log.Printf("Received a message in bytes: %s", d.Body)
			if fn != nil {
				fn(d.Body,fnParams)
			}


			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages in "+queueName+". To exit press CTRL+C")
	<-forever
	return nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Println("%s: %s", msg, err)
	}
}

func SendDelayedTask(conn *amqp.Connection,delayTime int64,exchangeName string,queueName string,bodyData interface{},fn func(interface{},[]interface{}) interface{},fnParams ...interface{}) error{
	ch, err := conn.Channel()
	defer ch.Close()
	log.Printf(" Error %s", err)
	headersChannel := make(amqp.Table)
	headersChannel["x-delayed-type"]= "direct";
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"x-delayed-message",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		headersChannel,//nil,           // arguments
	)
	//q, err := ch.QueueDeclare(
	//	"task_queue", // name
	//	true,         // durable
	//	false,        // delete when unused
	//	false,        // exclusive
	//	false,        // no-wait
	//	nil,          // arguments
	//)

	headers := make(amqp.Table)
	headers["x-delay"] = delayTime

	bodyBytes,err := GetBytes(bodyData);
	err = ch.Publish(
		exchangeName,     // exchange
		queueName,//q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			Headers:headers,
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bodyBytes,
		})
	if fn != nil {
		fn(bodyData,fnParams)
	}
	log.Printf(" [x] Sent Delayed.Message %s", bodyData)
	return nil
}

func ReceiveDelayedTask(conn *amqp.Connection,exchangeName string,queueName string,fn func([]byte,[]interface{}) interface{},fnParams ...interface{}) error{
	ch, err := conn.Channel()
	defer ch.Close()
	log.Printf(" Error %s", err)
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"x-delayed-message",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	q, err := ch.QueueDeclare(
		"",//"task_queue", // name
		false,//true,         // durable
		false,        // delete when unused
		true,//false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	err = ch.QueueBind(
		q.Name,        // queue name
		queueName,             // routing key
		exchangeName, // exchange
		false,
		nil)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received Delayed.Message a message: %s", d.Body)
			if fn != nil {
				fn(d.Body,fnParams)
			}
			log.Printf("Done Delayed.Message")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for Delayed.Message messages. To exit press CTRL+C 22222")
	<-forever
	return nil
}