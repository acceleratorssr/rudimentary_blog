package status_type

import "encoding/json"

type Permission int

const (
	PermissionAdmin Permission = iota
	PermissionUser
	PermissionNormal
	PermissionBanned
)

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p Permission) String() string {
	switch p {
	case PermissionAdmin:
		return "Admin"
	case PermissionUser:
		return "User"
	case PermissionNormal:
		return "visitor"
	case PermissionBanned:
		return "Banned"
	default:
		return "Unknown"
	}
}
