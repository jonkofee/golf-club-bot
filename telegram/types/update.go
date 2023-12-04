package types

type Update struct {
	Id      int64 `json:"update_id"`
	Message *struct {
		Id            int64 `json:"message_id"`
		Chat          *Chat
		From          *User
		Text          *string
		NewChatMember *User `json:"new_chat_member"`
	}
}
