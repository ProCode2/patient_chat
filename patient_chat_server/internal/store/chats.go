package store

import (
	types "github.com/patient_chat/patient_chat_server/internal/data"
)

// gets chat by userID
func (d *DbStore) GetChats(uid string) ([]types.Chat, error) {
	s := `SELECT ID, patient_id, doc_id, thread_id, query, response, time FROM (SELECT *, ROW_NUMBER() OVER (PARTITION BY thread_id ORDER BY date(time) DESC) AS rn FROM chats) WHERE rn = 1;`
	var cs []types.Chat
	err := d.db.Select(&cs, s, uid)
	if err != nil {
		return []types.Chat{}, err
	}

	return cs, nil
}

func (d *DbStore) AddChat(id, pid, did, tid, q, r string) error {
	s := `INSERT INTO chats (ID, patient_id, doc_id, thread_id, query, response) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := d.db.Exec(s, id, pid, did, tid, q, r)

	if err != nil {
		return err
	}

	return nil
}

func (d *DbStore) GetChatsByThreadID(tid string) ([]types.Chat, error) {
	s := `SELECT * FROM chats WHERE thread_id = ?`
	var cs []types.Chat

	err := d.db.Select(&cs, s, tid)
	if err != nil {
		return []types.Chat{}, err
	}

	return cs, nil
}
