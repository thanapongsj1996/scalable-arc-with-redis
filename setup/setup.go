package setup

import (
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) error {
	// Create Table if not exists
	db.Exec("CREATE TABLE IF NOT EXISTS members (id VARCHAR(50) NOT NULL, username VARCHAR(500) NOT NULL, is_active INT NOT NULL, CONSTRAINT PK_members PRIMARY KEY (id));")

	return nil

	//// Check if data exists
	//var count int64
	//db.Table("members").Count(&count)
	//if count > 0 {
	//	return nil
	//}
	//
	//// Seed member to table members
	//numberOfMembers := 100
	//members := make([]*model.Member, numberOfMembers)
	//for i := 0; i < numberOfMembers; i++ {
	//	id := i
	//	// Half of member is active, another half is inactive
	//	isActive := 0
	//	if i%2 == 0 {
	//		isActive = 1
	//	}
	//
	//	member := &model.Member{
	//		ID:       fmt.Sprintf("id_%d", id),
	//		Username: fmt.Sprintf("user_%d", id),
	//		IsActive: isActive,
	//	}
	//	members[i] = member
	//}
	//
	//db.CreateInBatches(members, 1000)
	//return nil
}
