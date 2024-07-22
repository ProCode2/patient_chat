-- Add up migration script here
INSERT INTO users (id, name, role, phone) VALUES("a5450f4c-a163-42e4-a9b7-432b3ea807a3", "Philip Jones", "doctor", "1234567890");
INSERT INTO doctors (id, user_id, hospital, qualification) VALUES("f139395e-886d-4ee5-b996-33e3db66ebe0", "a5450f4c-a163-42e4-a9b7-432b3ea807a3", "Virgina City Hospital", "MBBS/Eye Specialist");
