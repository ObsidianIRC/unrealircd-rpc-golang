package unrealircd

import (
	"errors"
)

// Server handles server-related operations
type Server struct {
	querier Querier
}

// GetAll returns a list of all servers
func (s *Server) GetAll() (interface{}, error) {
	result, err := s.querier.Query("server.list", nil, false)
	if err != nil {
		return nil, err
	}

	if res, ok := result.(map[string]interface{}); ok {
		if list, ok := res["list"]; ok {
			return list, nil
		}
	}

	return nil, errors.New("invalid JSON response from UnrealIRCd RPC")
}

// Get returns a server object
func (s *Server) Get(server *string) (interface{}, error) {
	params := map[string]interface{}{}
	if server != nil {
		params["server"] = *server
	}

	result, err := s.querier.Query("server.get", params, false)
	if err != nil {
		return nil, err
	}

	if res, ok := result.(map[string]interface{}); ok {
		if srv, ok := res["server"]; ok {
			return srv, nil
		}
	}

	return nil, nil // not found
}

// Rehash asks the server to perform a rehash. The response can be either a boolean or
// an object containing rehash details depending on the target server's RPC support.
func (s *Server) Rehash(server *string) (interface{}, error) {
	var params map[string]interface{}
	if server != nil {
		params = map[string]interface{}{"server": *server}
	}
	result, err := s.querier.Query("server.rehash", params, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Connect instructs the server to make a server link to another server.
func (s *Server) Connect(link string) (bool, error) {
	result, err := s.querier.Query("server.connect", map[string]interface{}{"link": link}, false)
	if err != nil {
		return false, err
	}
	if b, ok := result.(bool); ok {
		return b, nil
	}
	// Accept "ok" string as success for flexibility
	if str, ok := result.(string); ok && str == "ok" {
		return true, nil
	}
	return false, nil
}

// Disconnect instructs the server to terminate a server link.
func (s *Server) Disconnect(link string) (bool, error) {
	result, err := s.querier.Query("server.disconnect", map[string]interface{}{"link": link}, false)
	if err != nil {
		return false, err
	}
	if b, ok := result.(bool); ok {
		return b, nil
	}
	if str, ok := result.(string); ok && str == "ok" {
		return true, nil
	}
	return false, nil
}

// ModuleList retrieves the list of modules loaded on a server.
func (s *Server) ModuleList(server *string) (interface{}, error) {
	var params map[string]interface{}
	if server != nil {
		params = map[string]interface{}{"server": *server}
	}
	result, err := s.querier.Query("server.module_list", params, false)
	if err != nil {
		return nil, err
	}
	if res, ok := result.(map[string]interface{}); ok {
		if list, ok := res["list"]; ok {
			return list, nil
		}
		// If it isn't a keyed response, return the whole object
		return res, nil
	}
	return nil, errors.New("invalid JSON response from UnrealIRCd RPC")
}

// ConfigTest runs a configuration test on the server and returns the full test object.
func (s *Server) ConfigTest() (map[string]interface{}, error) {
	result, err := s.querier.Query("server.config_test", nil, false)
	if err != nil {
		return nil, err
	}
	if res, ok := result.(map[string]interface{}); ok {
		return res, nil
	}
	return nil, errors.New("invalid JSON response from UnrealIRCd RPC")
}