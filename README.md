# Patient Chat
The project is a AI chat interface for patients. Patients can chat with the AI and get their queries resolved.
Patients will save their medical history and linked doctor in the app and they can be used to generate context for the AI.

# Try it
[Live Link](https://patient-chat.onrender.com/)
- Phone Number: `1122334455`
- Otp: `123456` (or any 6 charcter text)

# Technology used
- The API is built using Golang with `chi` rouuter 
- The Frontend is built using `React`, `shadcn`

# Features
- Sign up with a patient account
- Login in with a patient account
- Logout
- Chat with the simulated AI
- View chat history
- Update patient information
  - Update patient medical history
  - Update patient name and linked doctor

# Database Schema
```sql
CREATE TABLE users (
    id TEXT UNIQUE NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE doctors (
    id TEXT UNIQUE NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    qualification TEXT,
    hospital TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE patients (
    id TEXT UNIQUE NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    doc_id TEXT NOT NULL,
    medical_history TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE sessions (
    user_id TEXT NOT NULL,
    session_id TEXT NOT NULL PRIMARY KEY,
    expires_at TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

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
```
