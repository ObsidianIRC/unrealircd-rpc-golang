package unrealircd

// Message handles message-related operations
type Message struct {
	querier Querier
}

// Privmsg sends a PRIVMSG to a target (nick, channel, or multiple targets)
// target may be a string (single target) or []string (multiple targets)
func (m *Message) Privmsg(target interface{}, message string) (interface{}, error) {
	result, err := m.querier.Query("message.privmsg", map[string]interface{}{
		"target":  target,
		"message": message,
	}, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Notice sends a NOTICE to a target (nick, channel, or multiple targets)
// target may be a string (single target) or []string (multiple targets)
func (m *Message) Notice(target interface{}, message string) (interface{}, error) {
	result, err := m.querier.Query("message.notice", map[string]interface{}{
		"target":  target,
		"message": message,
	}, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Numeric sends a custom numeric message to a user
func (m *Message) Numeric(nick string, numeric int, message string) (interface{}, error) {
	result, err := m.querier.Query("message.numeric", map[string]interface{}{
		"nick":    nick,
		"numeric": numeric,
		"message": message,
	}, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StandardReply sends a standard reply to a user (IRCv3 standard replies)
func (m *Message) StandardReply(nick, replyType, code, description string, context *string) (interface{}, error) {
	params := map[string]interface{}{
		"nick":        nick,
		"type":        replyType,
		"code":        code,
		"description": description,
	}
	if context != nil {
		params["context"] = *context
	}
	result, err := m.querier.Query("message.standardreply", params, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Wallops sends a WALLOPS (server-wide admin message)
func (m *Message) Wallops(message string) (interface{}, error) {
	result, err := m.querier.Query("message.send_wallops", map[string]interface{}{"message": message}, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Globops sends a GLOBOPS (global operator message)
func (m *Message) Globops(message string) (interface{}, error) {
	result, err := m.querier.Query("message.send_globops", map[string]interface{}{"message": message}, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}
