package stype

// Permission 管理员1，用户2，普通3，禁止4
type Permission int

const (
	PermissionAdmin Permission = 1 + iota
	PermissionUser
	PermissionNormal
	PermissionBanned
)

//func (p Permission) MarshalJSON() ([]byte, error) {
//	return json.Marshal(p.String())
//}

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
