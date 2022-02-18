package util

const UserRoleAdmin = "admin"
const UserRoleScraper = "scraper"

const EventFieldString = "string"
const EventFieldNumber = "number"
const EventFieldBoolean = "boolean"

func GetAllRoles() []string {
	return []string{
		UserRoleAdmin,
		UserRoleScraper,
	}
}

func GetAllEventFieldTypes() []string {
	return []string{
		EventFieldString,
		EventFieldNumber,
		EventFieldBoolean,
	}
}
