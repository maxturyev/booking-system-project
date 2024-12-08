package databases

import "gorm.io/gorm"

// Database shares global database instance
var Database *gorm.DB
