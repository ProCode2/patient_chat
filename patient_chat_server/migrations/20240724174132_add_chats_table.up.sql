-- Add up migration script here
CREATE TABLE chats(
    ID TEXT PRIMARY KEY,
    patient_id TEXT NOT NULL,
    doc_id TEXT NOT NULL,
    thread_id TEXT NOT NULL,
    query TEXT NOT NULL,
    response TEXT NOT NULL,
    time TEXT default CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES users(id),
    FOREIGN KEY (doc_id) REFERENCES doctors(id)
);
