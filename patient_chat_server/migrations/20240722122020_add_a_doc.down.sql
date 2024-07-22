-- Add down migration script here
DELETE FROM doctors WHERE user_id = "a5450f4c-a163-42e4-a9b7-432b3ea807a3";
DELETE FROM doctors WHERE id = "a5450f4c-a163-42e4-a9b7-432b3ea807a3";
