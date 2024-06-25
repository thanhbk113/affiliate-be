package redisapp

import "fmt"

const (
	PrefixStaff                = "staffs"
	PrefixStaffDepartmentAdmin = "staff_list_department_admin"
)

// GetKeyStaff ...
func GetKeyStaff(staff, device string) string {
	return fmt.Sprintf("%s_%s_%s", PrefixStaff, staff, device)
}

// GetKeyStaffListDepartmentAdmin ...
func GetKeyStaffListDepartmentAdmin(staff string) string {
	return fmt.Sprintf("%s_%s", PrefixStaffDepartmentAdmin, staff)
}
