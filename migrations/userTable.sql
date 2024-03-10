CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE,
	username TEXT UNIQUE,
	password TEXT,
	session_token TEXT DEFAULT NULL,
	expiresAt DATETIME DEFAULT NULL,
	role TEXT DEFAULT 'user',
	created_at DATE DEFAULT (datetime('now','localtime'))
);

-- admin passowd: 123qweASD!@#
INSERT INTO user (email,username,password) 
SELECT 'admin@gmail.com','administrator','$2a$10$NtyegXYqaRIpCCA9fA.uguuypOqOQM34O7sqRdrlBMrmRFu5cdIKq'
WHERE NOT EXISTS (SELECT 1 FROM user WHERE id = 1);
UPDATE user SET role = 'admin' WHERE id= 1;
