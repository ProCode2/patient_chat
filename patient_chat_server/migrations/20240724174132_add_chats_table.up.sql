-- Add up migration script here
CREATE TABLE chats(
ID TEXT PRIMARY KEY,
patient_id TEXT,
doc_id TEXT,
query TEXT,
response TEXT,
time TEXT default CURRENT_TIMESTAMP,
FOREIGN KEY (patient_id) REFERENCES users(id),
FOREIGN KEY (doc_id) REFERENCES doctors(id)
);
