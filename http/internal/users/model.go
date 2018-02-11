package model

import (
	"time"
)

// User contains information about user.
type User struct {
	UserID       int        `json:"user_id" validate:"upwd"`
	UserType     int        `json:"type"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"lat_name"`
	Password     string     `json:"password"`
	Email        string     `json:"email"`
	Company      string     `json:"company"`
	Image        string     `json:"image"`
	DataCreated  *time.Time `json:"data_created"`
	DataModified *time.Time `json:"data_modified"`
}

/*
	create table users (id int primary key auto_increment,userType int not null,
						firstName varchar(100) not null,lastName varchar(100) not null,
						password varchar(36) not null,email varchar(60) not null,
						company varchar (256), image varchar(256) not null,
						dataCreated time ,dataModified timestamp
						 );
*/
