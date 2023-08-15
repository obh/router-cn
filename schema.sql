CREATE TABLE paymentintent (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_email VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(255) NOT NULL,
    customer_address TEXT NOT NULL,
    status VARCHAR(255) NOT NULL,
    amount float(10,2) NOT NULL,
    currency varchar(10) NOT NULL,
    added_on DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_on DATETIME DEFAULT CURRENT_TIMESTAMP
);

 CREATE TABLE `paymentattempt` (
  `id` int NOT NULL AUTO_INCREMENT,
  `payment_intent_id` varchar(255) NOT NULL,
  `payment_method` varchar(255) NOT NULL,
  `processor` varchar(255) DEFAULT NULL,
  `processor_ref` varchar(255) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `amount` float(10,2) NOT NULL,
  `currency` varchar(10) NOT NULL,
  `added_on` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_on` datetime DEFAULT CURRENT_TIMESTAMP,
  `metadata` text,
  `payment_hash` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

