package userDatabase

import "github.com/quinntas/go-rest-template/internal/api/utils"

const (
	TABLE_NAME   utils.Key = "Users"
	TABLE_SCHEMA utils.Key = `
		CREATE TABLE IF NOT EXISTS Users (
			id INT AUTO_INCREMENT NOT NULL,
			pid VARCHAR(191) NOT NULL,
			email VARCHAR(191) NOT NULL,
			password VARCHAR(191) NOT NULL,
			roleId INT NOT NULL,
			createdAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY (pid),
			UNIQUE KEY (email)
		);`
)
