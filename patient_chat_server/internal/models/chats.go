package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/patient_chat/patient_chat_server/internal/ai"
	types "github.com/patient_chat/patient_chat_server/internal/data"
)

func GetChatsForUser(uid string) ([]types.Chat, error) {
	cs, err := db.GetChats(uid)

	if err == sql.ErrNoRows {
		return make([]types.Chat, 0), nil
	}
	if err != nil {
		log.Println("Can not get chats for user: ", err)
		return nil, errors.New("Something went wrong while getting chats")
	}

	return cs, nil
}

func AddUserChat(pid, did, tid, query string) (bool, error) {
	res := ai.GetResponse(query)
	if tid == "" {
		tid = GenID()
	}
	err := db.AddChat(GenID(), pid, did, tid, query, res)

	if err != nil {
		log.Println("can not add chat: ", err)
		return false, errors.New("Something went wrong while adding chats")
	}
	return true, nil
}

func GetChatsByThreadID(tid string) ([]types.Chat, error) {
	cs, err := db.GetChatsByThreadID(tid)

	if err != nil {
		log.Println("Can not get chats by thread id: ", err)
		return nil, errors.New("Something went wrong while getting chats in a thread")
	}

	return cs, nil
}
