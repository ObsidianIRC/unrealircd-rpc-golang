package unrealircd

// Log handles log operations
type Log struct {
	querier Querier
}

// Subscribe subscribes to log events. Any previous subscriptions are overwritten
func (l *Log) Subscribe(sources []string) (interface{}, error) {
	return l.querier.Query("log.subscribe", map[string]interface{}{
		"sources": sources,
	}, false)
}

// Unsubscribe unsubscribes from all log events
func (l *Log) Unsubscribe() (interface{}, error) {
	return l.querier.Query("log.unsubscribe", nil, false)
}

// GetAll gets past log events
func (l *Log) GetAll(sources []string) (interface{}, error) {
	params := map[string]interface{}{}
	if sources != nil {
		params["sources"] = sources
	}

	result, err := l.querier.Query("log.list", params, false)
	if err != nil {
		return nil, err
	}

	if res, ok := result.(map[string]interface{}); ok {
		if list, ok := res["list"]; ok {
			return list, nil
		}
	}

	return nil, nil
}

// Send sends a log message to the IRC server.
// Requires UnrealIRCd 6.1.8 or later.
// Parameters:
//   - msg: The human-readable log message
//   - level: Log level (info, advice, warn, debug, error, fatal)
//   - subsystem: A subsystem identifier (e.g., "webpanel")
//   - eventID: A unique event identifier (e.g., "WEBPANEL_LOGIN")
func (l *Log) Send(msg, level, subsystem, eventID string) (bool, error) {
	_, err := l.querier.Query("log.send", map[string]interface{}{
		"msg":       msg,
		"level":     level,
		"subsystem": subsystem,
		"event_id":  eventID,
	}, false)
	if err != nil {
		return false, err
	}
	return true, nil
}
