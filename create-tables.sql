CREATE TABLE message (
	id VARCHAR PRIMARY KEY,
	payload VARCHAR NULL,
	senderId VARCHAR NULL,
	receiverId VARCHAR null,
	moment VARCHAR NULL,
    consumed BOOLEAN NOT NULL,
    consume_moment VARCHAR NULL
);

CREATE TABLE chat_user (
	id VARCHAR PRIMARY KEY,
	username VARCHAR NULL,
	email VARCHAR NULL,
	public_key VARCHAR null
);